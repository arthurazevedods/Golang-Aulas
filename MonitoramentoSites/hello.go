package main

import "fmt"

func main() {
	nome := "Arthur"

	versao := 0.4

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Esse programa está na versão:", versao)

	fmt.Println("MENU")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando digitado foi:", comando)

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exbindo Logs...")
	case 0:
		fmt.Println("Saindo...")
	default:
		fmt.Println("Não conheço este comando")
	}
}
