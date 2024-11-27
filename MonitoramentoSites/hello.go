package main

import (
	"encoding/json"
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

	versao := 0.7

	fmt.Println("\nOlá, sr.", nome)
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

	sites, _ := leSitesDoArquivo()
	for i := 0; i < ciclos_monitoramento; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

}

func exibirLogs() {
	fmt.Println("Exbindo Logs...")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

func leSitesDoArquivo() ([]string, error) {
	// Abrir o arquivo JSON
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório:", err)
	} else {
		fmt.Println("Diretório atual:", dir)
	}
	file, err := os.Open("sites.json")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		return nil, err
	}
	defer file.Close()

	// Decodificar o JSON diretamente do arquivo
	var data struct {
		Sites []string `json:"sites"`
	}

	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return nil, err
	}

	fmt.Println("\nOs sites lidos são:")
	for i, site := range data.Sites {
		fmt.Println(i, " - ", site)
	}
	// Retornar o array de sites
	return data.Sites, nil
}
