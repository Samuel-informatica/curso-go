package main

import (
	"fmt"
)

type Pessoa struct { // seria tipo objketos do js
	Nome          string
	Idade         int
	Nacionalidade string
}

func main() {

	//Structs
	myPessoa := Pessoa{Nome: "Samuel", Idade: 25, Nacionalidade: "Brasileiro"}
	fmt.Println(myPessoa)      // mostrando todas as informações da myPessoa
	fmt.Println(myPessoa.Nome) // mostrando uma informação especifico da myPessoa

	myPessoa.Nome = "myPessoa" //mudando apenas o nome de myPessoa
	fmt.Println(myPessoa)

	//map com struct

	pessoa3 := map[int]Pessoa{
		1: {Nome: "Samuel Maas", Idade: 30, Nacionalidade: "Argentino"},
		2: {Nome: "Samuel Maas Anton", Idade: 40, Nacionalidade: "Mexicano"},
	}

	fmt.Println("Todos", pessoa3)                    // acessando todos do map com structs
	fmt.Println("Uma pessoa especifica", pessoa3[1]) // acessando diretamento um nome especifico dentro do map
}
