package main

import "fmt"

// Interface para representar veiculos
// VelocidadeMaxima, MeioDeLocomoção

type Veiculo interface {
	VelocidadeMaxima() int
	MeioDeLocomocao() string
	MostraNome() string
}

type Carro struct {
	Nome    string
	MaxVelo int
}

func (c Carro) VelocidadeMaxima() int {
	return c.MaxVelo
}

func (c Carro) MeioDeLocomocao() string {
	return "Rodas"
}

func (c Carro) MostraNome() string {
	return c.Nome
}

func (t Trem) VelocidadeMaxima() int {
	return t.MaxVelo
}

func (t Trem) MeioDeLocomocao() string {
	return "Trilhos"
}

func (t Trem) MostraNome() string {
	return t.Nome
}

func DescreverVeiculo(v Veiculo) {
	fmt.Println("Veiculo atual:", v.MostraNome())
	fmt.Printf("Meio de locomoção: %s, Velocidade Máxima: %d km/h \n", v.MeioDeLocomocao(), v.VelocidadeMaxima())
}

type Trem struct {
	Nome    string
	MaxVelo int
}

type Barco struct {
	Nome    string
	MaxVelo int
}

func (b Barco) VelocidadeMaxima() int {
	return b.MaxVelo
}

func (b Barco) MeioDeLocomocao() string {
	return "Motor"
}

func (b Barco) MostraNome() string {
	return b.Nome
}

func main() {

	meuCarro := Carro{MaxVelo: 130, Nome: "Gol 1.0"}
	meuTrem := Trem{MaxVelo: 90, Nome: "Trem GO"}
	meuBrarco := Barco{MaxVelo: 100, Nome: "Titanic"}

	DescreverVeiculo(meuCarro)
	DescreverVeiculo(meuTrem)
	DescreverVeiculo(meuBrarco)

}
