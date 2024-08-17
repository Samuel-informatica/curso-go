package main

import (
	"fmt"
)

func main() {

	var numeros [4]int

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40

	fmt.Println(numeros)

	var numero = [4]int{10, 20, 30, 40}

	fmt.Println(numero)

	permission := [3]string{"usuario", "admin", "editor"}

	fmt.Println(permission)
}
