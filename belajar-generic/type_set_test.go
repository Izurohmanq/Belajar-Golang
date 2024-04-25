package belajargeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int

type Number interface {
	~int | int8 | int16 | int64 | float32 | float64
}

func Min[T Number](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestMin(t *testing.T) {
	assert.Equal(t, 100, Min[int](100, 200))
	assert.Equal(t, int64(100), Min[int64](int64(100), int64(200)))
	assert.Equal(t, Age(100), Min[Age](Age(100), Age(200))) //tambahin tilde ~ kalau mau penulisannya seperti ini
}

func TestMinTypeInference(t *testing.T) { // lebih singkat saja, karena golang bisa langsung bisa baca type parameter
	assert.Equal(t, 100, Min(100, 200))
	assert.Equal(t, int64(100), Min(int64(100), int64(200)))
	assert.Equal(t, Age(100), Min(Age(100), Age(200)))
}
