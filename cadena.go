package main

import (
	"fmt"
	"seminarioGo2021/services"
)

func main() {
	var e services.Cadena
	e, err := services.NewCadena1("ldfafa")
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println(e)

}
