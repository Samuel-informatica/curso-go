package main

import (
	"fmt"
)

func main() {

	//for

	// enquanto o i for menor que 5 ou o numero declarado nessa parte, o for estará executando
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	numero := 0

	for numero < 10 {
		fmt.Println("For", numero)
		numero++
	}

	for i := 0; i < 10; i++ {
		fmt.Println("condition in if", i)

		if i == 4 {
			fmt.Println("for igual a 4")
			continue
		}

		if i == 7 {
			fmt.Println("for parado aqui")
			break // o break acaba não executando o restante, mesmo que o i não tenha finalizado de fato
		}
		fmt.Println("Finalizando")
	}

}
