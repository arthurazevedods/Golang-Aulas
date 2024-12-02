package main

import (
	"exercicio/clientes"
	"exercicio/contas"
	"fmt"
)

func main() {
	clienteFulano := clientes.Titular{
		Nome:      "Fulano",
		CPF:       "70045612389",
		Profissao: "Advogado",
	}

	clienteSicrano := clientes.Titular{
		Nome:      "Sicrano",
		CPF:       "65542372309",
		Profissao: "Engenheiro",
	}

	contaDoFulano := contas.ContaCorrente{
		Titular:       clienteFulano,
		NumeroAgencia: 123,
		NumeroConta:   12567,
		Saldo:         550.25,
	}
	contaDoSicrano := contas.ContaCorrente{
		Titular:       clienteSicrano,
		NumeroAgencia: 123,
		NumeroConta:   22681,
		Saldo:         2000.82,
	}

	fmt.Println(contaDoFulano)
	fmt.Println(contaDoSicrano)
}
