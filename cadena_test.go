package main

import (
	services "seminarioGo2021/services"
	"testing"

	"github.com/stretchr/testify/assert"
	//"github.com/go-playground/assert"
	//"github.com/stretchr/testify/assert"
)

func TestCambiarCadena(t *testing.T) {

	c := services.NewCadena("TX", "ABCDEF", 6)
	services.CambiarCadena(&c, "TX", "ABC", 2)
	if c.Valor != "ABC" {
		t.Error("Los valores no coinciden")
	}
	assert.Equal(t, c.Valor, "ABC", "El valor no cambió")
	//assert.Equal(t, c.Tipo, "NN")
	//assert.Equal(t, c.Tipo, "NN", "El tipo no coincide")
}
func TestValor(t *testing.T) {
	c := services.NewCadena("TX", "ABCDEF", 6)
	if len(c.Valor) != c.Largo {
		t.Error("El valor no coincide con el largo de la cadena")
	}
}

/*func TestCadenaParser(t *testing.T) {
	var tests = []struct {
		Input   string // input string in order to be parsed
		Success bool   // paser result
		Tipo    string // the input type
		Valor   string // the input value
		Largo   int    // value length
	}{
		{"NN03123", true, "NN", "123", 3},
		{"TX09ABCDEFGHI", true, "NN", "ABCDEFGHI", 9},
		{"TX05ABCDEFG", false, "", "", 0},
		{"NN020A", false, "", "", 0},
	}

	for _, tt := range tests {

		r, err := services.NewCadena([]byte(tt.Input))
		assert.Equal(t, err == nil, tt.Success)
		//assert.Equal(t, r.Tipo, "NN", "El tipo no coincide")
		assert.Equal(t, len(Valor) == Largo, tt.Success)
	}
}*/
