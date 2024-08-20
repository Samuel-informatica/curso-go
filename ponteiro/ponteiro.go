// package main

// import "fmt"

// // Um ponteiro é uma variável que armazena o endereço de memória de outra variável.

// func main() {
// 	// var numero int = 60
// 	// var p *int = &numero

// 	// fmt.Println(numero)
// 	// fmt.Println("Endereço na memoria da var numero:", &numero)

// 	// fmt.Println("Valor do ponteiro P:", p)
// 	// fmt.Println("Valor apontado por p:", *p)
// 	// // Imprimir o valor no endereço armazenado em P (dereferência)

// 	// numero = 10

// 	// fmt.Println("Valor apontado por p ATUALIZADO:", *p)
// }
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Um ponteiro é uma variável que armazena o endereço de memória de outra variável.

package main

import "fmt"

// Um ponteiro é uma variável que armazena o endereço de memória de outra variável.

func somar(num *int) {
	*num = *num + 1
}

func subtrair(num *int) {
	*num = *num - 1
}

func main() {
	numero := 10
	fmt.Println("Valor inicial da variavel numero:", numero)
	fmt.Println("Endereço de memoria da variavel numero", &numero)

	somar(&numero)
	somar(&numero)

	fmt.Println("Valor atual da variavel numero:", numero)

	subtrair(&numero)

	fmt.Println("Valor atual da variavel numero após subtrair:", numero)

}
