package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Asep",
		MiddleName: "Kayu",
		LastName:   "Jati",
		Hobbies:    []string{"ngoding", "Malassss"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Asep","MiddleName":"Kayu","LastName":"Jati","Hobbies":["ngoding","Malassss"]}`
	jsonBytes := []byte(jsonString) //di sini intinya

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer) // terus untuk decodenya

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "ASEPPPP",
		Addresses: []Address{
			{
				Street:     "Jalan Gak ada",
				Country:    "malas",
				PostalCode: "92310391",
			},
		},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"ASEPPPP","MiddleName":"","LastName":"","Hobbies":null,"Addresses":[{"Street":"Jalan Gak ada","Country":"malas","PostalCode":"92310391"}]}`
	jsonBytes := []byte(jsonString) //di sini intinya

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer) // terus untuk decodenya

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}
