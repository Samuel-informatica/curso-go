package main

import (
	"fmt"
)

func main() {

	//Maps
	// Declaração inicial

	var capitais map[string]string = make(map[string]string)
	//Chave       valor
	capitais["França"] = "Paris"

	fmt.Println(capitais) //acessa todas capitais

	fmt.Println(capitais["França"]) //acessa uma capital especifica

	//Conseguimos mudar os valores de cada map, basta acessar a sua respectiva chave e adicionar novo valor, exemplo abaixo

	capitais["França"] = "Tóquio"

	fmt.Println(capitais)
}
