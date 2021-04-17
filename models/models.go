package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	EmpID    string             `json:"empID,omitempty" bson:"empID,omitempty"`
	FullName string             `json:"fullName" bson:"fullName,omitempty"`
	Company  *Company           `json:"company" bson:"company,omitempty"`
}

type Company struct {
	CompanyName    string `json:"companyName,omitempty" bson:"companyName,omitempty"`
	CompanyAddress string `json:"companyAddress,omitempty" bson:"companyAddress,omitempty"`
}
