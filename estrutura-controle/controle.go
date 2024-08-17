package main

import (
	"fmt"
)

//if - else - else if
//
//

func main() {

	nota1 := 7.30
	nota2 := 4.66

	media := (nota1 + nota2) / 2

	if media >= 6 {
		fmt.Printf("Você passou de ano e sua média ficou: %.2f", media)
	} else {
		fmt.Printf("Você não passou de ano e sua média ficou: %.2f", media)
	}

	//Switch
	DiaSemana := 7

	switch DiaSemana {

	case 1:
		fmt.Println("Segunda-feria")
		break

	case 2:
		fmt.Println("Terça-feria")
		break

	case 3:
		fmt.Println("Quarta-feria")
		break

	case 4:
		fmt.Println("Quinta-feria")
		break

	case 5:
		fmt.Println("Sexta-feria")
		break

	case 6:
		fmt.Println("Sabado")
		break

	case 7:
		fmt.Println("Domingo")
		break
	}
}
