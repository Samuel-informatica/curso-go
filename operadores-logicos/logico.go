package main

import (
	"fmt"
)

//and (&&)
//or (||)
//negação (!)

func main() {

	estoque := true
	vendasLiberadas := true

	fmt.Println("AND:", estoque && vendasLiberadas) // ambos os dois precisam ser verdadeiro para retornar algum resultado (true), caso contrario vai retornar como false e não dará continuidade.

	estoque2 := true
	vendasLiberadas2 := false

	fmt.Println("OR:", estoque2 || vendasLiberadas2) // apenas um dos dois precisa ser verdadeiro para retornar algum resultado (true), caso ambos estejam como false, retornará false

	freteGratis := false

	fmt.Println("Negação:", !freteGratis) //irá inverter o resultado, que nem nesse exmplo o freteGratis já vem como false, utilizando a negação, atuomaticamente vai se tornar como true

	nome1 := "Samuel"
	nome2 := "Leticia"

	fmt.Println("Diferente:", nome1 != nome2) //Estou comparando se o nome1 é diferente do nome2, nesse exemplo são diferentes, retornaria como true, caso fossem iguais, retornaria como false
}
