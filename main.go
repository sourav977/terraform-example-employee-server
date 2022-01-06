package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	ctrl "github.com/sourav977/terraform-example-employee-server/controllers"
)

func init() {
	if os.Getenv("MONGO_CONNECTION_URL") == "" {
		os.Setenv("MONGO_CONNECTION_URL", "mongodb://mongo:27017")
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/getAllEmployees", ctrl.GetAllEmployees).Methods("GET")
	router.HandleFunc("/api/addEmployee", ctrl.AddEmployee).Methods("POST")
	//filter by empID and update. So except empID, all values can be updated
	router.HandleFunc("/api/updateEmployeeByID", ctrl.UpdateEmployeeByEmployeeID).Methods("PUT")
	router.HandleFunc("/api/DeleteEmployeeByID", ctrl.DeleteEmployeeByEmployeeID).Methods("DELETE")
	router.HandleFunc("/healthcheck", ctrl.Healthcheck).Methods("GET")
	router.Use(loggingMiddleware)

	//start server
	log.Fatal(http.ListenAndServe(":8000", router))

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("in Logger")
		defer log.Println("Logger ended")
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		log.Println(r.Header, r.Body)
		next.ServeHTTP(w, r)
	})
}
