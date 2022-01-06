package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/sourav977/terraform-example-employee-server/helper"
	"github.com/sourav977/terraform-example-employee-server/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	var employes []models.Employee
	collection := helper.ConnectToDB()

	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		helper.SetError(err, http.StatusInternalServerError, w)
		return
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var emp models.Employee
		err := cur.Decode(&emp)
		if err != nil {
			log.Fatal(err)
		}
		employes = append(employes, emp)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employes)
}

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	collection := helper.ConnectToDB()
	//decode req body into emp
	_ = json.NewDecoder(r.Body).Decode(&emp)

	result, err := collection.InsertOne(context.TODO(), emp)

	if err != nil {
		helper.SetError(err, http.StatusInternalServerError, w)
		return
	}
	//json returntype
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func UpdateEmployeeByEmployeeID(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	collection := helper.ConnectToDB()
	//decode req body into emp
	_ = json.NewDecoder(r.Body).Decode(&emp)
	//except employeeID, all values can be updated
	filter := bson.M{"empID": emp.EmpID}
	result, err := collection.ReplaceOne(context.TODO(), filter, emp)

	if err != nil {
		helper.SetError(err, http.StatusInternalServerError, w)
		return
	}
	//json returntype
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func DeleteEmployeeByEmployeeID(w http.ResponseWriter, r *http.Request) {
	empID := r.URL.Query().Get("empID")
	collection := helper.ConnectToDB()
	filter := bson.M{"empID": empID}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		helper.SetError(err, http.StatusInternalServerError, w)
		return
	}
	//json returntype
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
