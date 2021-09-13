package main

import (
	"fmt"
	"seminarioGo2021/services/cadena"
)

func main() {
	service, err := cadena.NewCadenaService()
	if err != nil {
		panic(err) //muere la ejecución
	}

	cadena, err := service.Login("tipo", "valor", 6)
	if err != nil {
		panic(err)
	}

	fmt.Println(cadena)
}
