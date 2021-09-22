package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCadenaParser(t *testing.T) {
	var tests = []struct {
		Input   string // input string in order to be parsed
		Success bool   // paser result
		Tipo    string // the input type
		Largo   int    // value length
		Valor   string // the input value

	}{
		{"NN03123", true, "NN", 3, "123"},
		{"TX09ABCDEFGHI", true, "NN", 9, "ABCDEFGHI"},
		{"TX05ABCDEFG", false, "", 0, ""},
		{"NN020A", false, "", 0, ""},
		{"ER020A", false, "", 0, ""},
	}

	for _, tt := range tests {

		_, err := NewCadena(tt.Input)
		if err == nil {
			assert.Equal(t, err == nil, tt.Success)
		}

	}
}
