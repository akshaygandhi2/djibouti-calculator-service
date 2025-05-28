package main

import (
	"buildingcost/config"
	"buildingcost/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	config.LoadEnv()
	r := mux.NewRouter()

	api := r.PathPrefix("/calculator-service/v1").Subrouter()

	// Register routes
	api.HandleFunc("/{code}/estimate_calculate", controller.CalculateCost).Methods("POST")
	api.HandleFunc("/{code}/demand", controller.GeneratetDemands).Methods("POST")

	log.Println("Server listening on http://localhost:8085")
	log.Fatal(http.ListenAndServe(":8085", r))
}
