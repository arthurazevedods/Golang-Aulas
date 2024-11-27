package main

import "fmt"

func main() {
	nome := "Arthur"

	versao := 0.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Esse programa está na versão:", versao)
	/*
		fmt.Println("MENU")
		fmt.Println("1 - Iniciar Monitoramento")
		fmt.Println("2 - Exibir Logs")
		fmt.Println("0 - Sair do Programa")
	*/
	var comando int
	//fmt.Scanf("%d", &comando)
	//fmt.Println("O ponteiro do comando de:", &comando)
	//fmt.Println("O comando digitado foi:", comando)

	fmt.Println("MENU")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
	fmt.Scan(&comando)
	fmt.Println("O comando digitado foi:", comando)

}
