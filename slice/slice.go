package main

import (
	"fmt"
)

func main() {

	//slice

	var lista []int //declarando slice inteiro
	//onde desejo colocar -------depois da vírgula é o que quero colocar dentro da lista
	lista = append(lista, 10, 20, 30, 40)
	fmt.Println(lista)

	//slice a partir de um array

	MyArray := [5]int{10, 10, 10, 10}

	mySlices := MyArray[0:3] // a ultima posição não é pego

	fmt.Println("Meu slice a partir de um Array:", mySlices)

	nomes := make([]string, 0) // 0 quer dizer que não quero que seja iniciado com nada

	fmt.Println(nomes)

	nomes = append(nomes, "Samuel", "Leticia")
	fmt.Println(nomes)
}
