package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}
type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Hobbies    []string
	Addresses  []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Asep",
		MiddleName: "Kayu",
		LastName:   "Mahoni",
	}

	bytes, _ := json.Marshal(customer) //buat encode
	fmt.Println(string(bytes))
}
