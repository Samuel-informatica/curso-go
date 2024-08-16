package main

import (
	"fmt"
)

func main() {

	//Desafio 1 - Calcular a Idade atual

	AnoNascimento := 1999
	AnoAtual := 2024

	IdadeAtual := AnoAtual - AnoNascimento
	fmt.Println("A idade atual seria", IdadeAtual, "anos")

	//Desafio 2 - Converter horas para minutos

	qtdHoras := 2
	const minutos = 60
	MinConvertido := qtdHoras * minutos
	fmt.Println("Horas para minutos", MinConvertido)
}
