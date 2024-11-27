package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	introducao()
	for {
		comando := menu()

		switchComandos(comando)
	}

}

func introducao() {
	nome := "Arthur"

	versao := 0.4

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Esse programa está na versão:", versao)
}

func menu() int {
	fmt.Println("MENU")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando digitado foi:", comando)
	return comando
}

func switchComandos(comando int) {
	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		exibirLogs()
	case 0:
		fmt.Println("Saindo...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	//slices
	sites := []string{
		"https://httpbin.org/status/200",
		"https://github.com/",
		"https://arthurazevedods.vercel.app/",
		"https://supervisao-e-sinergia.vercel.app/",
	}

	for i := range len(sites) {
		resp, _ := http.Get(sites[i])
		if resp.StatusCode == 200 {
			fmt.Println("Site:", sites[i], "foi carregado com sucesso!")
		} else {
			fmt.Println("Site:", sites[i], "está com problemas. Status Code:", resp.StatusCode)
		}
	}

}

func exibirLogs() {
	fmt.Println("Exbindo Logs...")
}
