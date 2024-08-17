package main

import (
	"fmt"
)

func main() {

	var tarefas []string

	tarefas = append(tarefas, "Tarefa1") //1
	tarefas = append(tarefas, "Tarefa2") //2
	tarefas = append(tarefas, "Tarefa3") //3
	tarefas = append(tarefas, "Tarefa4") //4
	tarefas = append(tarefas, "Tarefa5") //5
	tarefas = append(tarefas, "Tarefa6") //6

	fmt.Println(tarefas)

	fmt.Println(len(tarefas))

	//slicing / cortar

	var tarefa2 []string

	tarefa2 = append(tarefa2, "one", "two", "three", "four", "five", "six", "seven", "eight")

	fmt.Println(tarefa2)
	fmt.Println(tarefa2[1:])  // cortando o indice 0
	fmt.Println(len(tarefa2)) // tamanho da tarefa2
	fmt.Println(tarefa2[:6])  //cortando todos a partir do indice 6

	// tarefa2 = tarefa2[:len(tarefa2)-1]

	// fmt.Println("Resultado final", tarefa2)

	tarefa2 = append(tarefa2[:1], tarefa2[3:]...)

	fmt.Println("cortando o indice 1, 2", tarefa2) //Apresentando sem os indices 1 e 2
}
