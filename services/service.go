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

func NewCadena1(cad string) Cadena {
	var cadena Cadena
	/*if len(cad) > 0{
	cadena.Tipo = cad
	cadena.Largo = retornaLargoCadena(cad)
	cadena.Valor = ""*/
	cadena.Largo = strconv.Atoi(cad[2])
	cadena.Tipo = cad[4:]
	//if
	//}else {
	//retorna el error por ser una cadena vacia
	//}
	return Cadena{cadena.Tipo, cadena.Largo, cadena.Valor}
}

/*func NewCadena2(s string) Cadena {
	var n = strconv.Atoi(s[:2])

	return Cadena{s[:2], n, s[4:]}
}*/

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

/*func retornaLargoCadena(c string){
	var cant = 0
	for i, _ = range c {
		if i = 2{
			cant = strconv.Atoi(c[i])
		}
	}
}*/
