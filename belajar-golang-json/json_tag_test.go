package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"` //gakbole ada spasi, jadi => json:"apagitu" jangan kyk json: "apagitu"
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "102301293",
		Name:     "Apple malas",
		ImageURL: "http://google.com",
	}
	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"102301293","name":"Apple malas","image_url":"http://google.com"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	json.Unmarshal(jsonBytes, product)

	fmt.Println(product)
}
