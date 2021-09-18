package main

import (
	"fmt"
	services "seminarioGo2021/services"
)

func CambiarCadena(c *services.Cadena, tipo, value string, length int) {
	c.Tipo = tipo
	c.Value = value
	c.Length = length
}

func main() {
	/*service, err := model.NewCadenaService()
	if err != nil {
		panic(err) //muere la ejecución
	}

	cadena, err := service.Login("tipo", "valor", 6)
	if err != nil {
		panic(err)
	}
	*/
	//fmt.Println(cadena)
	c := services.NewCadena("NN", "asdf", 4) //funcion constructora (go no tiene constructores)
	CambiarCadena(&c, "TX", "ABCDEF", 6)     // & es la dir de memoria
	fmt.Println(c)
}
