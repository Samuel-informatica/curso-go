package main

import (
	"fmt"
)

// = - Atribuição
// == - Comparação
// >= - Maior ou igual
// > maior que
// < - Menor que
// <= - Menor ou igual
// len - pega quantidade de caracteres do nome - tamanho da string

func main() {

	//==
	valor1 := 10
	valor2 := 10

	fmt.Println(valor1 == valor2) //verdadeiro

	//>=
	valor3 := 10
	valor4 := 10

	fmt.Println(valor3 >= valor4) //verdadeiro

	//>
	valor5 := 10
	valor6 := 10

	fmt.Println(valor5 > valor6) //falso

	//<
	valor7 := 10
	valor8 := 10

	fmt.Println(valor7 < valor8) //falso

	//<=
	valor9 := 10
	valor10 := 10

	fmt.Println(valor9 <= valor10) //verdadeiro

	//len
	nome1 := "Samuel"
	nome2 := "Leticia"

	fmt.Println(len(nome1), len(nome2)) // quantidade de caracteres em cada nome - tamanho da string
}
