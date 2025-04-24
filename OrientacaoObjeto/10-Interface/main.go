package main

import (
	"exercicio/clientes"
	"exercicio/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	clienteFulano := clientes.Titular{
		Nome:      "Fulano",
		CPF:       "70045612389",
		Profissao: "Advogado",
	}

	contaDoFulano := contas.ContaCorrente{
		Titular:       clienteFulano,
		NumeroAgencia: 123,
		NumeroConta:   12567,
	}
	poupancaDoFulano := contas.ContaPoupanca{
		Titular:  clienteFulano,
		Operacao: 12,
	}
	contaDoFulano.Depositar(100)
	poupancaDoFulano.Depositar(520)
	PagarBoleto(&contaDoFulano, 120)
	PagarBoleto(&poupancaDoFulano, 1200)

	fmt.Println(contaDoFulano.ObterSaldo())
	fmt.Println(poupancaDoFulano.ObterSaldo())

	PagarBoleto(&poupancaDoFulano, 120)
	fmt.Println(poupancaDoFulano.ObterSaldo())

	PagarBoleto(&poupancaDoFulano, -120)
	fmt.Println(poupancaDoFulano.ObterSaldo())

	PagarBoleto(&poupancaDoFulano, 1020)
	fmt.Println(poupancaDoFulano.ObterSaldo())

}
