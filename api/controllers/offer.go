package controllers

// HandleOffer is the function to accept or request for driver
import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/wyiemay/convoy-api/api/models"
)

func HandleOffer(w http.ResponseWriter, req *http.Request) {
	r := render.New(render.Options{})
	vars := mux.Vars(req)
	offerID := vars["id"]
	decoder := json.NewDecoder(req.Body)
	var o models.OfferPutInput
	err := decoder.Decode(&o)
	if err != nil {
		panic(err)
	}
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
	var shipmentID int
	var isShipmentAccepted int
	err = db.QueryRow("SELECT shipments.id, accepted FROM shipments LEFT JOIN offers on shipments.id = offers.shipment_id WHERE offers.id = ? ", offerID).Scan(&shipmentID, &isShipmentAccepted)

	if err != nil {
		log.Fatal(err)
	}
	if isShipmentAccepted == 1 {
		w.WriteHeader(http.StatusAlreadyReported)
		return
	}

	// step 2. update offers table status
	stmt, err := db.Prepare("UPDATE offers set status = ? WHERE id = ?")
	checkErr(err)
	_, err = stmt.Exec(o.Status, offerID)
	checkErr(err)

	// step 3. update shipments table
	if o.Status == "ACCEPT" {
		stmt, err = db.Prepare("UPDATE shipments set accepted = ? WHERE id = ?")
		checkErr(err)
		_, err = stmt.Exec(1, offerID)
		checkErr(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	r.JSON(w, 200, struct{}{})
}
