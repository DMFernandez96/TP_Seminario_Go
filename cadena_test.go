package main

import (
	services "seminarioGo2021/services"
	"testing"
	//"github.com/go-playground/assert"
	//"github.com/go-playground/assert/v2"
)

func TestCambiarCadena(t *testing.T) {

	c := services.NewCadena("TX", "ABCDEF", 6)
	CambiarCadena(&c, "TX", "ABC", 2)
	if c.Value != "ABC" {
		t.Error("Los valores bno coinciden")
	}
	//assert.Equal(t, c.Tipo, "NN")
	//assert.Equal(t, c.Tipo, "NN", "el tipo no coincide")
}
