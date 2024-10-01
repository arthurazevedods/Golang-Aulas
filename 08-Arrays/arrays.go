package main

import "fmt"

func main() {
	// declaração de um array de inteiros com 5 posições (valores padrão são 0)
	var a [5]int
	fmt.Println("emp:", a)

	//define o último elemento da lista a como sendo 100
	a[4] = 100
	fmt.Println("set:", a)    // imprime o array com o valor atualizado
	fmt.Println("get:", a[4]) // recupera e imprime o valor do índice 4

	//imprime o tamanho da lista
	fmt.Println("len:", len(a))

	// declaração e inicialização de um array de inteiros com 5 valores
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// atualização dos valores usando a sintaxe de comprimento implícito
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// usando a inicialização de índice:
	// o elemento no índice 0 recebe 100
	// o elemento no índice 3 recebe 400,
	// o elemento no índice 4 recebe 500,
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// declaração de um array 2D (2x3) e preenchimento com valores baseados na soma dos índices
	var doisD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			doisD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", doisD) // imprime o array 2D preenchido

	//inicializa um array 2D diretamente com valores
	doisD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", doisD) // imprime o array 2D inicializado
}
