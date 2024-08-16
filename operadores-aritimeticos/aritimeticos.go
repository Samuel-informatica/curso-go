package main

import (
	"fmt"
)

func main() {
	//adição

	var valor1 = 10
	var valor2 = 7

	resultado := valor1 + valor2
	fmt.Println("Resultado da adição:", resultado)

	//subtração

	var valor3 = 4
	var valor4 = 2

	resultadoSubtracao := valor3 - valor4

	fmt.Println("Resultado da subtração:", resultadoSubtracao)

	//Multiplicação

	var valor5 = 4
	var valor6 = 2

	resultadoMultiplicacao := valor5 * valor6

	fmt.Println("Resultado da Multiplicação:", resultadoMultiplicacao)

	//Divisão

	var valor7 = 4
	var valor8 = 2

	resultadoDivisao := valor7 / valor8

	fmt.Println("Resultado da Divisão:", resultadoDivisao)

	//Resto da divisão

	var valor9 = 10
	var valor10 = 3

	resultadoRestoDivisao := valor9 % valor10

	fmt.Println("Resultado do resto da divisão:", resultadoRestoDivisao)
}
