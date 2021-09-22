package services

import (
	"errors"
	"strconv"
)

type Cadena struct {
	Tipo  string
	Largo int
	Valor string
}

func ControlarLargoValor(c *Cadena) (string, error) {
	if len(c.Valor) != c.Largo {
		return c.Valor, errors.New("Error, el valor" + c.Valor + " no es del largo correspondiente")
	}
	return c.Valor, nil
}

func NewCadena1(cad string) (Cadena, error) {

	var c Cadena
	if len(cad) > 0 {
		c.Tipo = cad[:2]
		_, errTipo := comprobarTipo(c.Tipo)
		r, errLargo := strconv.Atoi(cad[2:4])
		c.Largo = r
		c.Valor = cad[4:]
		_, errValor := ControlarLargoValor(&c)
		if errTipo == nil && errLargo == nil && errValor == nil {
			return c, nil
		}
	}
	return Cadena{"", 0, ""}, errors.New("Error, cadena no valida")
}

func comprobarTipo(s string) (string, error) {
	if s == "TX" || s == "NN" {
		return s, nil
	}
	return s, errors.New("Tipo de cadena no valido")
}
