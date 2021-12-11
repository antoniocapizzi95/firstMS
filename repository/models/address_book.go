package models

type AddressBook struct {
	Persons []Person `json:"persons" bson:"persons"`
}

type Person struct {
	Name   string `json:"name" bson:"name"`
	Number uint64 `json:"number" bson:"number"`
}
