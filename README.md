# curso-go
Novo curso de programação, linguagem Go Lang com professor Matheus Fraga 

# Curso de Go Lang

## Seção 1

- É uma linguagem aberta, criado pelo Google
- É conhecida pela sua simplicidade, eficiência
- É uma linguagem altamente escalável
- É uma linguagem altamente compilada

- Simplicidade e Legibilidade
- Sintaxe limpa e fácil de entender
- Extrema Performance (comparavel )

Teste realizado para ver se estava funcionando o código 

```
package main

import "fmt"

func main() {
	fmt.Println("Rodando")
}
```

- **fmt** - é um pacote de print utilizado para essa linguagem
- **Println** - Para mostrar o retorno em tela no terminal, porém essa utilização é para  quebra de linha, ou seja, caso tiver outro print  abaixo, não vai aparecer na mesma linha e sim abaixo
- **Printf** - formatação de variaveis de nome, numero, entre outros 
- **main** - pode ter apenas um por arquivo


- **%.1f** - Define quantidade de decimais que vai parecer para o número após o ponto (.), nesse caso iria aparecer 30.1
- **%s** - Formatação para texto, nesse caso usamos quando queremos que uma variável de nome no meio do texto, exemplo fmt.Printf("Olá %s seja bem vindo ao Curso de GO", nome)

## Seção 2 - Fundação GO

var - É uma variável mutável, a var ela pode ser modificada a qualquer momento do código, porém só será aceito a tipagem que foi declarada a ela, se eu declarei meu var como string não consigo inserir valor do tipo int ou vice-verso.

const - Não pode ser alterado o valor após a inicialização.