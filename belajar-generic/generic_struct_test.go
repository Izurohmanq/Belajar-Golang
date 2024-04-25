package belajargeneric

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First, Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "asep",
		Second: "Kayu",
	}

	fmt.Println(data)
}
func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "asep",
		Second: "Kayu",
	}

	assert.Equal(t, "Budy", data.ChangeFirst("Budy"))
	assert.Equal(t, "Hello Eko", data.SayHello("Eko"))
}
