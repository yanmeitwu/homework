package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wyiemay/convoy-api/api/models"
)

// CreateDriver is a function to handle create|view driver
func CreateDriver(w http.ResponseWriter, req *http.Request) {
	// open and connect to DB
	db, err := sql.Open("mysql", "convoy:convoy@tcp(database:3306)/convoy?parseTime=true")
	checkErr(err)
	// create driver
	decoder := json.NewDecoder(req.Body)
	var d models.DriverPostInput
	err = decoder.Decode(&d)
	if err != nil {
		panic(err)
	}
	// insert
	stmt, err := db.Prepare("INSERT INTO drivers(capacity) values(?)")
	checkErr(err)
	res, err := stmt.Exec(d.Capacity)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)

	dr := models.DriverPostOutput{
		ID: int(id),
	}

	json.NewEncoder(w).Encode(&dr)
}

// ViewDriverOffers is a function to get a list of shipment avaiable for driver
func ViewDriverOffers(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	driverID := vars["id"]
	// open and connect to DB
	db, err := sql.Open("mysql", "convoy:convoy@tcp(database:3306)/convoy?parseTime=true")
	checkErr(err)
	// query all the offers avabable to this driver
	rows, err := db.Query("SELECT offers.id, shipment_id from offers LEFT JOIN shipments ON offers.shipment_id = shipments.id WHERE offers.driver_id = ? AND shipments.accepted = 0 AND offers.status != 'PASS'", driverID)
	defer rows.Close()
	checkErr(err)
	var offerID int
	var shipmentID int

	dr := []models.DriverGetOutput{}

	for rows.Next() {
		err = rows.Scan(&offerID, &shipmentID)
		dr = append(dr, models.DriverGetOutput{
			OfferID:    offerID,
			ShipmentID: shipmentID,
		})
	}
	json.NewEncoder(w).Encode(&dr)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
