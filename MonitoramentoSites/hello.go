package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const delay = 5
const ciclos_monitoramento = 4

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
	for i := 0; i < ciclos_monitoramento; i++ {
		for i, site := range sites {
			resp, _ := http.Get(site)
			if resp.StatusCode == 200 {
				fmt.Println("Site número ", i+1, ". Site:", site, "foi carregado com sucesso!")
			} else {
				fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
			}
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

}

func exibirLogs() {
	fmt.Println("Exbindo Logs...")
}
