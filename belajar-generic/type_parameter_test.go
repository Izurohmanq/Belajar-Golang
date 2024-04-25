package belajargeneric

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T { //any itu == interface{}
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	var resultString string = Length[string]("asep")
	assert.Equal(t, "asep", resultString)

	var resultNumber int = Length[int](123)
	assert.Equal(t, 123, resultNumber)

}
