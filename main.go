package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	ctrl "github.com/sourav977/mongo-backend/controllers"
)

func init() {
	os.Setenv("MONGO_CONNECTION_URL", "mongodb://localhost:27017")
}
func main() {
	fmt.Println(os.Getenv("MONGO_CONNECTION_URL"))
	router := mux.NewRouter()
	router.HandleFunc("/api/getAllEmployees", ctrl.GetAllEmployees).Methods("GET")
	router.HandleFunc("/api/addEmployee", ctrl.AddEmployee).Methods("POST")
	router.HandleFunc("/healthcheck", ctrl.Healthcheck).Methods("GET")
	router.HandleFunc("/readiness", ctrl.Readiness).Methods("GET")

	//start server
	log.Fatal(http.ListenAndServe(":8000", router))

}
