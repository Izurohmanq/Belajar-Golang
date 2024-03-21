package belajargolangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "AOKAOKAOK",
		MiddleName: "malasss",
		LastName:   "oakdoawkdoawkdaokd",
	}

	encoder.Encode(customer)
}
