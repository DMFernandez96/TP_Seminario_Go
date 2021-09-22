package main

import (
	"fmt"
	"seminarioGo2021/services"
)

/*func ControlarLargoValor(c *services.Cadena){
	if len(c.Valor) != c.Largo{
		fmt.Println("Error, el valor"+ c.Valor + " no es del largo "+ )
	}
}*/

func main() {
	var e services.Cadena
	e, err := services.NewCadena1("ldfafa")
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println(e)

}
