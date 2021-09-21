package services

import (
	"fmt"
	"strconv"
)

type Cadena struct {
	Tipo  string
	Largo int
	Valor string
}

func NewCadena(tipo, valor string, largo int) Cadena {
	return Cadena{tipo, largo, valor} //genera una instancia de la estructura Cadena
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

func NewCadena1(cad string) Cadena {

	var c Cadena
	c.Tipo = cad[:2]
	r, err := strconv.Atoi(cad[2:4])
	c.Largo = r
	if err != nil {
		panic(err)
	}
	c.Tipo = cad[4:]
	return Cadena()
}

/*func NewCadena2(s string) Cadena {
	var n = strconv.Atoi(s[2])
	return Cadena{s[:2], n, s[4:]} //genera una instancia de la estructura Cadena
}
/*func CambiarLargo(c string) int{

	for
	return strconv.Atoi(c[2])
}*/
