package main

import (
	"fmt"
	"strconv" // conversão de string pra int e int para string
)

// tipo de conversão
func main() {

	// inteeiro paara float

	var valor1 int = 45
	var valor1Convertido = float64(valor1)
	fmt.Println("Valor foi convertido de int para Floar", valor1Convertido)

	//converter int para string

	var valor2 int = 50
	var valorString string = strconv.Itoa(valor2)
	fmt.Println("Valor foi convertido de int para String", valorString)

	//converter string para int

	var valor3 string = "200.4"

	valorInteiro, err := strconv.Atoi(valor3) //converter string para int

	if err != nil {
		fmt.Println("Deu erro na hora de converter a string para int")
	} else {
		fmt.Println("Valor String foi convertido para int", valorInteiro)
	}

	//Converter String para float

	var receitaFederal = "1234.65"

	RecebeValorFloat, error := strconv.ParseFloat(receitaFederal, 64)

	//se o error for igual ou diferente de nil
	if error != nil {
		fmt.Println("Deu algum erro na conversão", error)
	} else {
		fmt.Println("Valor em Float", RecebeValorFloat)
	}
}
