package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"oakdaokdaokad", "name":"Malsss bangetttt", "price":200000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{} //kalau punya jsonnya tidak terprediksi struktur datanya daripada menggunakan struct, jadi pake map
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P0001",
		"name":  "Malas Bangtetetet",
		"price": 21398123,
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
