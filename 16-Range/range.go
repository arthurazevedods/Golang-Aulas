package main

import "fmt"

func main() {
	// Criando um slice de inteiros com os valores 2, 3 e 4
	nums := []int{2, 3, 4}
	soma := 0

	// Usando um loop 'for' com 'range' para iterar sobre o slice 'nums'
	// O 'range' retorna dois valores: o índice (que não estamos usando, por isso '_') e o valor do elemento
	for _, num := range nums {
		soma += num // Acumulando os valores de 'nums' na variável 'soma'
	}
	fmt.Println("soma:", soma) // Exibe a soma dos elementos (2 + 3 + 4 = 9)

	// Iterando novamente sobre 'nums', desta vez utilizando o índice 'i'
	for i, num := range nums {
		// Se o valor do elemento for igual a 3, imprime o índice correspondente
		if num == 3 {
			fmt.Println("índice:", i)
		}
	}

	// Criando um mapa (dicionário) com pares chave-valor
	frutas := map[string]string{"a": "maçã", "b": "banana"}

	// Iterando sobre o mapa, onde 'chave' representa a chave e 'v' o valor
	for chave, v := range frutas {
		// Exibe cada par chave-valor do mapa no formato 'chave -> valor'
		fmt.Printf("%s -> %s\n", chave, v)
	}

	// Iterando apenas sobre as chaves do mapa
	for chave := range frutas {
		// Exibe as chaves presentes no mapa
		fmt.Println("chave:", chave)
	}

	// Iterando sobre uma string, onde 'i' é o índice e 'c' é o código Unicode (rune) de cada caractere
	for i, c := range "go" {
		// Exibe o índice e o valor Unicode de cada caractere da string "go"
		fmt.Println(i, c)
	}
}
