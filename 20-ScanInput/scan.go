package main

import "fmt"

func main() {
	var nome string
	fmt.Println("Qual o seu nome?")
	fmt.Scanf("%s", &nome)

	fmt.Println("Olá, sr.", nome)

	/*
		fmt.Println("MENU")
		fmt.Println("1 - Comando")
		fmt.Println("2 - Comando")
		fmt.Println("0 - Comando")
	*/

	var comando int
	//fmt.Scanf("%d", &comando)
	//fmt.Println("O ponteiro do comando de:", &comando)
	//fmt.Println("O comando digitado foi:", comando)

	fmt.Println("MENU")
	fmt.Println("1 - Comando")
	fmt.Println("2 - Comando")
	fmt.Println("0 - Comando")
	fmt.Scan(&comando)
	fmt.Println("O ponteiro do comando de:", &comando)
	fmt.Println("O comando digitado foi:", comando)

}
