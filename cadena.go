package main

import (
	"fmt"
	services "seminarioGo2021/services"
)

/*func ControlarLargoValor(c *services.Cadena){
	if len(c.Valor) != c.Largo{
		fmt.Println("Error, el valor"+ c.Valor + " no es del largo "+ )
	}
}*/

func main() {
	//fmt.Println(cadena)
	c := services.NewCadena("NN", "asdf", 4)      //funcion constructora (go no tiene constructores)
	services.CambiarCadena(&c, "TX", "ABCDEF", 6) // & es la dir de memoria
	fmt.Println(c)

}
