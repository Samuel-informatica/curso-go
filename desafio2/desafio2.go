package main

import (
	"fmt"
)

func main() {
	// Foormula: j = p * i * n
	//j = juros
	//p = valor presente capital
	//i = taxa de juros
	//n = numero de meses

	p := 2000.59
	i := 0.03
	n := 18

	var j float64 = float64(p) * i * float64(n)
	fmt.Printf("Resultado finado dos juros é: %.2f", j)

}
