package main

import (
	"fmt"
)

type Usuario struct {
	Nome      string
	Idade     int
	Estado    string
	Permissao string
}

func main() {

	usuario := map[int]Usuario{
		1: {Nome: "Ciudad", Idade: 18, Estado: "RS", Permissao: "admin"},
		2: {Nome: "Samuel", Idade: 40, Estado: "SC", Permissao: "master"},
		3: {Nome: "Ciudad", Idade: 28, Estado: "RJ", Permissao: "master"},
		4: {Nome: "Ciudad", Idade: 15, Estado: "SP", Permissao: "master"},
		5: {Nome: "Leticia", Idade: 27, Estado: "RS", Permissao: "admin"},
	}
	fmt.Println(usuario)
	fmt.Println(len(usuario))

	i := 1
	for i <= len(usuario) {
		fmt.Printf("Nome: %s, idade: %d, Estado: %s, Permissão: %s \n", usuario[i].Nome, usuario[i].Idade, usuario[i].Estado, usuario[i].Permissao)
		i++
	}

	// for _, usuario := range usuario {
	// 	if usuario.Estado == "RS" {
	// 		fmt.Println("Pessoas do RS:", usuario.Nome)
	// 	}
	// }

}
