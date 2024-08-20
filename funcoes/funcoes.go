package main

import (
	"errors"
	"fmt"
)

// func boasvindas() {
// 	fmt.Println("funcionado")

// }

// a função está aguardando um paramentro
func boasvindas(nome string) {
	fmt.Printf("Bem vindo %s \n", nome)

}

// é preciso sempre declarar que tipo de retorno é essa função
func calcular(num1 int, num2 int) int {
	resultado := num1 + num2
	return resultado
}

// strcuts
type Usuario struct {
	Nome  string
	Senha string
}

// user vai ter o comportamento do structs e senha é uma variável para recerber info de senha
func auth(user Usuario, senha string) (string, error) {

	if user.Senha != "1123" { // se a senha que está armazenado dentro do struct Senha for diferente de "123" terá o retorno de senha incorreta
		return "", errors.New("Password icorrect")
	}
	//caso o retorno da senha for igual vai retornar o Nome salvo dentro do struct, caso contrario vai retornar erro
	return user.Nome, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//Desafio

func UserNew(username string) (string, bool) {
	if username != "" {
		return username, true
	}

	return "Usuário sem username", false
}

func main() {

	boasvindas("Samuel")

	resultadoDaFuncao := calcular(123, 40)

	fmt.Printf("O resultado final é: %d  \n", resultadoDaFuncao)

	//userMandaStructs variavel criado para receber as info e mandar para o struct
	userMandaStructs := Usuario{Nome: "samuel@gmail.com", Senha: "d1123"}

	//emailuser, err, duas variaveis criados para receber os retornos da funcao
	emailuser, err := auth(userMandaStructs, "213") // chamando minha funcao e mando as duas informações dentro dos parenteses

	//se o erro for diferente de nil vamos receber esse retorno abaixo
	if err != nil {
		fmt.Println(err)
	} else { //caso contrario o retorno será esse abaixo
		fmt.Printf("Usuario logado com email: %s\n", emailuser)
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	//Desafio

	Usuario11 := "Samuel"

	resultadoDesafio, status := UserNew(Usuario11)

	fmt.Printf("Õ resultado final é: %s, %t \n", resultadoDesafio, status)

	Usuario12 := ""

	resultadoDesafio2, status := UserNew(Usuario12)

	fmt.Printf("Õ resultado final é: %s, %t \n", resultadoDesafio2, status)

}
