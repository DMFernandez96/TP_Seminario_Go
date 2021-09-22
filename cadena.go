package main

import (
	"fmt"
	"seminarioGo2021/services"
)

func main() {
	var e services.Cadena
	e, err := services.NewCadena("TX02AB")
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println(e)

}
