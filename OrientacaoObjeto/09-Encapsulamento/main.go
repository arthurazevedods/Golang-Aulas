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
	}
	contaDoSicrano := contas.ContaCorrente{
		Titular:       clienteSicrano,
		NumeroAgencia: 123,
		NumeroConta:   22681,
	}

	poupancaDoFulano := contas.ContaPoupanca{
		Titular:  clienteFulano,
		Operacao: 12,
	}
	fmt.Println(poupancaDoFulano)
	fmt.Println(contaDoFulano.Depositar(225.90))
	fmt.Println(contaDoSicrano.Depositar(2125.10))
	fmt.Println(contaDoSicrano.Sacar(1121.00))

	fmt.Println(poupancaDoFulano.Depositar(10000.50))
	fmt.Println(poupancaDoFulano.ObterSaldo())
	fmt.Println(contaDoFulano.ObterSaldo())
	//fmt.Println(contaDoSicrano.ObterSaldo())

	fmt.Println(poupancaDoFulano.TransferirContaCorrente(1500, &contaDoFulano))

}
