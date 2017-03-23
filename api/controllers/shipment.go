package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wyiemay/convoy-api/api/models"
)

// CreateShipment is a function to create shipment and push offers to drivers
func CreateShipment(w http.ResponseWriter, req *http.Request) {
	// create driver
	decoder := json.NewDecoder(req.Body)
	var d models.DriverPostInput
	err := decoder.Decode(&d)
	if err != nil {
		panic(err)
	}

	// open and connect to DB
	db, err := sql.Open("mysql", "convoy:convoy@tcp(database:3306)/convoy?parseTime=true")
	checkErr(err)

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	// create shipment
	stmt, err := db.Prepare("INSERT INTO shipments(weight) values(?)")
	checkErr(err)
	res, err := stmt.Exec(d.Capacity)
	checkErr(err)
	shipmentID, err := res.LastInsertId()
	checkErr(err)

	// create offers
	rows, err := db.Query("SELECT id FROM drivers WHERE capacity >=? ORDER BY received_offers ASC", d.Capacity)
	defer rows.Close()
	checkErr(err)
	createOffer, err := db.Prepare("INSERT INTO offers(shipment_id, driver_id) values(?,?)")
	checkErr(err)

	offerResp := []models.OfferResp{}

	var driverID int
	for rows.Next() {
		err = rows.Scan(&driverID)
		res, err = tx.Stmt(createOffer).Exec(shipmentID, driverID)
		checkErr(err)
		offerID, err := res.LastInsertId()
		checkErr(err)
		offerResp = append(offerResp, models.OfferResp{
			OfferID:  int(offerID),
			DriverID: driverID,
		})
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	sr := models.ShipmentPostOutput{
		ID:     int(shipmentID),
		Offers: offerResp,
	}
	json.NewEncoder(w).Encode(&sr)
}

// ViewShipment returns the accepted offer or all outstanding offers for this shipment
func ViewShipment(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	shipmentID := vars["id"]
	// open and connect to DB
	db, err := sql.Open("mysql", "convoy:convoy@tcp(database:3306)/convoy?parseTime=true")
	checkErr(err)

	// begin transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	// step 1, check shipment is still aviable
	var isShipmentAccepted bool
	err = db.QueryRow("SELECT accepted FROM shipments WHERE id =? ", shipmentID).Scan(&isShipmentAccepted)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	offerResp := []models.OfferResp{}

	var (
		offerID  int
		driverID int
	)
	// shipment is accepted
	if isShipmentAccepted {
		err = db.QueryRow("SELECT id, driver_id FROM offers WHERE shipment_id =? AND status = 'ACCEPT'", shipmentID).Scan(&offerID, &driverID)
		if err != nil {
			log.Fatal(err)
		}
		offerResp = append(offerResp, models.OfferResp{
			OfferID:  offerID,
			DriverID: driverID,
		})
		sg := models.ShipmentGetOutput{}

		json.NewEncoder(w).Encode(&sg)
		return
	}
	// shipment is not accepted yet
	rows, err := db.Query("SELECT id, driver_id FROM offers WHERE shipment_id =? ", shipmentID)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&offerID, &driverID)
		checkErr(err)
	}
	offerResp = append(offerResp, models.OfferResp{
		OfferID:  int(offerID),
		DriverID: driverID,
	})

	sg := models.ShipmentGetOutput{
		Accepted: false,
		Offers:   offerResp,
	}

	json.NewEncoder(w).Encode(&sg)
}
