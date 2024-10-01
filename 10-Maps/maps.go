package main

import (
	"fmt"
	"maps"
)

func main() {

	// inicializando um mapa vazio com chaves do tipo string e valores do tipo int
	mapa := make(map[string]int)

	// Atribuindo valores ao mapa
	mapa["chave1"] = 7
	mapa["chave2"] = 13

	// imprimindo o mapa completo
	fmt.Println("mapa:", mapa)

	// recuperando o valor associado à "chave1" e guardando em valor1
	valor1 := mapa["chave1"]
	fmt.Println("valor1:", valor1)

	// tentando recuperar um valor para uma chave que não existe ("chave3")
	valor3 := mapa["chave3"]
	fmt.Println("valor3:", valor3)

	// imprimindo o tamanho do mapa
	fmt.Println("tamanho:", len(mapa))

	// removendo uma chave do mapa
	delete(mapa, "chave2")
	fmt.Println("mapa após delete:", mapa)

	// limpando o mapa (remove todas as chaves)
	clear(mapa)
	fmt.Println("mapa após clear:", mapa)

	// verificando se uma chave está presente no mapa (retorna um bool)
	_, presente := mapa["chave2"]
	fmt.Println("presente:", presente)

	// inicializando um novo mapa com valores predefinidos
	novoMapa := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("novo mapa:", novoMapa)

	// comparando dois mapas para verificar se são iguais
	novoMapa2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(novoMapa, novoMapa2) {
		fmt.Println("novoMapa é igual a novoMapa2")
	}
}
