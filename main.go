package main

import (
	"buildingcost/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/calculator-service/v1").Subrouter()

	// Register routes
	api.HandleFunc("/bpa/estimate_calculate", controller.CalculateCost).Methods("POST")

	log.Println("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
