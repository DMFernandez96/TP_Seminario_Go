package services

import (
	"fmt"
)

type Cadena struct {
	Tipo  string
	Largo int
	Valor string
}

func NewCadena(tipo, valor string, largo int) Cadena {
	return Cadena{tipo, largo, valor} //genera una instancia de la estructura Cadena
}

func NewCadena1(cad string) {
	var cadena Cadena
	cadena.Tipo = cad[:2]
	/*cadena.Largo = strconv.Atoi(cad[2:4])
	cadena.Largo = cad[2:4]
	cadena.Tipo = cad[4:]
	if*/
}

func CambiarCadena(c *Cadena, tipo, valor string, largo int) {
	c.Tipo = tipo
	c.Valor = valor
	c.Largo = largo
}

func ControlarLargoValor(c *Cadena) bool {
	r := true
	if len(c.Valor) != c.Largo {
		fmt.Println("Error, el valor" + c.Valor + " no es del largo correspondiente")
		r = false
	}
	return r
}
