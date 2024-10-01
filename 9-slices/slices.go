package main

import (
	"fmt"
	"slices"
)

func main() {

	// Declaração de uma slice de strings não inicializada
	var frutas []string
	fmt.Println("não inicializado:", frutas, frutas == nil, len(frutas) == 0)

	// Inicializa a slice com 3 elementos (valores padrão vazios)
	frutas = make([]string, 3)
	fmt.Println("vazio:", frutas, "tamanho:", len(frutas), "capacidade:", cap(frutas))

	// Definindo valores individuais na slice
	frutas[0] = "maçã"
	frutas[1] = "banana"
	frutas[2] = "cereja"
	fmt.Println("definido:", frutas)
	fmt.Println("obter:", frutas[2])

	// Exibe o comprimento atual da slice
	fmt.Println("tamanho:", len(frutas))

	// Adicionando mais elementos à slice usando `append`
	frutas = append(frutas, "damasco")
	frutas = append(frutas, "embaúba", "figo")
	fmt.Println("adicionado:", frutas)

	// Criando uma cópia da slice
	copia := make([]string, len(frutas))
	copy(copia, frutas)
	fmt.Println("cópia:", copia)

	// Criando slices a partir de fatias (subconjuntos) da original
	subSlice1 := frutas[2:5]
	fmt.Println("slice 1:", subSlice1)

	subSlice2 := frutas[:5]
	fmt.Println("slice 2:", subSlice2)

	subSlice3 := frutas[2:]
	fmt.Println("slice 3:", subSlice3)

	// Declaração e inicialização de outra slice
	outrosFrutos := []string{"goiaba", "hibisco", "ingá"}
	fmt.Println("declarado:", outrosFrutos)

	// Comparando duas slices
	outrosFrutos2 := []string{"goiaba", "hibisco", "ingá"}
	if slices.Equal(outrosFrutos, outrosFrutos2) {
		fmt.Println("outrosFrutos é igual a outrosFrutos2")
	}

	// Criando uma slice 2D (matriz) e preenchendo com valores
	doisD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		tamanhoInterno := i + 1
		doisD[i] = make([]int, tamanhoInterno)
		for j := 0; j < tamanhoInterno; j++ {
			doisD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", doisD)
}
