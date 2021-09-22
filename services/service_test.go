package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//"github.com/go-playground/assert"
	//"github.com/stretchr/testify/assert"
)

/*func TestCambiarCadena(t *testing.T) {

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
}*/

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
		/*{"TX05ABCDEFG", false, "", 0, ""},
		{"NN020A", false, "", 0, ""},*/
	}

	for _, tt := range tests {
		//var r Cadena

		_, err := NewCadena1(tt.Input)
		if err == nil {
			assert.Equal(t, err == nil, tt.Success)
		}

	}
}
