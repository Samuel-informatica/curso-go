package main

import "fmt"

func calcularMedia(notas Aluno1) float64 {
	resultado := (notas.Nota1 + notas.Nota2) / 2
	return resultado
}

func Aprovacao(ResultadoAluno2 float64) string {
	if ResultadoAluno2 >= 7 {
		return "Aluno Aprovado" // Resultado
	}
	return "Aluno Reprovado"
}

type Aluno1 struct {
	Nome  string
	Nota1 float64
	Nota2 float64
}

func main() {

	//Aluno1
	aluno1 := Aluno1{Nome: "Samuel", Nota1: 6.0, Nota2: 9.5}

	resultado_final := calcularMedia(aluno1)

	fmt.Printf("Media final do aluno %s: %.2f \n", aluno1.Nome, resultado_final)

	//Alunos 2
	aluno2 := Aluno1{Nome: "Leticia", Nota1: 3.5, Nota2: 6.5}

	ResultadoAluno2 := calcularMedia(aluno2)

	status := Aprovacao(ResultadoAluno2)

	fmt.Printf("Media final do aluno %s: %.2f - %s \n", aluno2.Nome, ResultadoAluno2, status)

	//------------------------------------------------------------------------------------------------------------------------//

	// COM FOR

	fmt.Println("Calcular com FOR")
	alunos := map[int]Aluno1{
		1: {Nome: "Samuel", Nota1: 6.0, Nota2: 9.5},
		2: {Nome: "Leticia", Nota1: 3.5, Nota2: 6.5},
		3: {Nome: "Dudu", Nota1: 1.5, Nota2: 3.5},
		4: {Nome: "KK", Nota1: 5.5, Nota2: 8.5},
		5: {Nome: "Jão", Nota1: 8.5, Nota2: 3.5},
	}

	i := 1
	for i <= len(alunos) {

		aluno := alunos[i]

		ResultadoAluno := calcularMedia(aluno)

		status := Aprovacao(ResultadoAluno)

		fmt.Printf("Media final do aluno %s: %.2f - %s \n", aluno.Nome, ResultadoAluno, status)
		i++
	}
}
