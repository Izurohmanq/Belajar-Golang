package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ini kalau manual banget
func TestHelloWorld(t *testing.T) {
	result := HelloWorld("asep")

	if result != "Hello asep" {
		// error
		t.Fail()
	}

	fmt.Println("TestHelloWorld Done")
}

// ini kalau pake import dari library orang(Testify), karena golang tidak menyediakan assertion
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("asep")

	assert.Equal(t, "Hello asep", result, "Result must be 'Hello asep'")
	fmt.Println("Test with assert Done")

}

// ini pake required test
func TestHelloWorldrequired(t *testing.T) {
	result := HelloWorld("asep")

	require.Equal(t, "Hello asep", result, "Result must be 'Hello asep'")
	fmt.Println("Test with required Done")

}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run on MacOS")
	}
	fmt.Println("Test Skip Done") // ini langsung gak dijalankan
}

func TestSubTest(t *testing.T) {
	t.Run("asep", func(t *testing.T) {
		result := HelloWorld("asep")
		assert.Equal(t, "Hello asep", result, "Result must be 'Hello asep'")
		fmt.Println("Test with assert 1 Done")
	})
	t.Run("Padli", func(t *testing.T) {
		result := HelloWorld("Padli")
		assert.Equal(t, "Hello Padli", result, "Result must be 'Hello Padli'")
		fmt.Println("Test with assert 2 Done")
	})
}

// hanya jalan di satu package, kalau ada package lain, buat lagi
func TestMain(m *testing.M) {
	//before unit test

	fmt.Println("Before unit Test")

	m.Run()

	fmt.Println("After unit Test")
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Asep",
			request:  "Asep",
			expected: "Hello Asep",
		},
		{
			name:     "Padli",
			request:  "Padli",
			expected: "Hello Padli",
		},
		{
			name:     "Ujang",
			request:  "Ujang",
			expected: "Hello Ujang",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
