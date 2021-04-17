package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/sourav977/mongo-backend/helper"
	"github.com/sourav977/mongo-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	var employes []models.Employee
	collection := helper.ConnectToDB()

	cur, err := collection.Find(context.TODO(), bson.M{})
	defer cur.Close(context.TODO())
	if err != nil {
		helper.SetError(err, http.StatusInternalServerError, w)
		return
	}

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

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func Readiness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
