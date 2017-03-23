package main

import (
	recovery "github.com/albrow/negroni-json-recovery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/wyiemay/convoy-api/api/controllers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/driver", controllers.CreateDriver).Methods("POST")
	router.HandleFunc("/driver/{id}", controllers.ViewDriverOffers).Methods("GET")
	router.HandleFunc("/shipment", controllers.CreateShipment).Methods("POST")
	router.HandleFunc("/shipment/{id}", controllers.ViewShipment).Methods("GET")
	router.HandleFunc("/offer/{id}", controllers.HandleOffer).Methods("PUT")

	n := negroni.New(negroni.NewLogger())
	n.Use(recovery.JSONRecovery(true))
	n.UseHandler(router)

	n.Run(":8080")
}
