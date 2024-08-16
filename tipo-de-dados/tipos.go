package main

import "fmt"

//Strintg
//Bool
//int | int8, int16, int64, uint8 - aceita até uma faxaetaria de números negativos e/ou positivos tamanho de memoria para o recebimento de informações dentro dessa declaração
//float32, float64
//uint16, uint64 não suporta número negativo tamanho de memoria para o recebimento de informações dentro dessa declaração
//float32, float64
//bytes, runes

func main() {

	var nome = "Samuel" //string
	fmt.Println(nome, "é uma string")

	var isAdmin = true //bool
	fmt.Println(isAdmin, "É um numero boleano")

	var numero1 int32 = 10 //int32
	fmt.Println(numero1)

	var numero2 int = 10 //int
	fmt.Println(numero2)

	var valor1 float32 = 23.25 //Float
	fmt.Println(valor1)

	valor2 := 35.20 //:= para cortar o valor terminado em 0
	fmt.Println(valor2)

	var teste byte = 'A' //byte
	fmt.Println(teste)

	var testeRune rune = '@' //Rune
	fmt.Println(testeRune)
}
