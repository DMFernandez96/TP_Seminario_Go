package services

import (
	//services "seminarioGo2021/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCadena(t *testing.T) {
	c := NewCadena("TX", "ABC", 3)
	if len(c.Valor) != c.Largo {
		t.Error("El largo no coincide con el largo del valor de la cadena")
	}
}

func TestControlarLargoValor(t *testing.T) {
	c := NewCadena("TX", "ABC", 3)
	if ControlarLargoValor(&c) == false {
		t.Error("El largo no coincide con el largo del valor de la cadena")
	}
}

func TestControlarLargoValorFalse(t *testing.T) {
	c := NewCadena("TX", "ABCDEFG", 3)
	if ControlarLargoValor(&c) == false {
		t.Error("El largo no coincide con el largo del valor de la cadena")
	}
}

func TestCambiarCadena(t *testing.T) {

	c := NewCadena("TX", "ABCDEF", 6)
	CambiarCadena(&c, "TX", "ABC", 2)
	/*if c.Valor != "ABC" {
		t.Error("Los valores no coinciden")
	}*/
	assert.Equal(t, c.Valor, "ABC", "El valor no coincide")
	//assert.Equal(t, c.Tipo, "NN")
	//assert.Equal(t, c.Tipo, "NN", "El tipo no coincide")
}

/*func TestIngresarCadena(t *testing.T) {
	c := IngresarCadena("TX", "ABCDEF", 6)
	if c != "TX06ABCDEF" {
		t.Error("Error en Ingresar cadena")
	}
}*/
