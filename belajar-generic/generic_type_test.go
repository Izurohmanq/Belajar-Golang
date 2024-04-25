package belajargeneric

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestBagString(t *testing.T) {
	names := Bag[string]{"Eko", "Asep", "mahia"}
	PrintBag(names)
}
func TestBagInt(t *testing.T) {
	names := Bag[int]{1, 2, 3, 4, 5}
	PrintBag(names)
}
