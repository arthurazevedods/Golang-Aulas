package main

import "fmt"

var global = "global"

func main() {
	//var declara uma ou mais variáveis
	//você pode declarar várias de uma vez
	//Go vai inferir o tipo da variável inicializada

	var a = "initial"
	fmt.Println(a)

	// Você declarar múltiplas variáveis
	// do mesmo tipo em uma única linha, separadas por vírgulas
	var b, c int = 1, 2
	fmt.Println(b, c)
	var b1, c1 = 3, 4
	fmt.Println(b1, c1)
	var b2, c2 string = "b2", "b3"
	fmt.Println(b2, c2)
	var b3, c3 = "b3", "b3"
	fmt.Println(b3, c3)
	var b4, c4 = 4, "b4"
	fmt.Println(b4, c4)

	//Variáveis ​​declaradas sem uma inicialização correspondente têm valor zero.
	//Por exemplo, o valor zero para um int é 0
	var numero int
	fmt.Println("Number:", numero)
	var d int
	fmt.Println(d)

	var booleano = true
	fmt.Println(booleano)

	//A sintaxe := é uma forma abreviada de declarar e inicializar uma variável
	//Essa sintaxe só está disponível dentro de funções.
	f := "apple"
	fmt.Println(f)

	//Se global estive sido inicializada como global := "global", iria dar erro
	fmt.Println(global)
}
