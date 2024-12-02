package main

import "fmt"

// Estrutura que representa uma conta corrente
type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

// Método para realizar saques na conta corrente
// Recebe o valor do saque como parâmetro e retorna uma mensagem de sucesso ou erro
// (c *ContaCorrent) aponta para a conta da pessoa que quer sacar
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	// Verifica se o saque é válido (valor positivo e menor ou igual ao saldo disponível)
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		// Deduz o valor do saque do saldo
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		// Retorna mensagem de erro se o saldo for insuficiente ou o valor inválido
		return "Saldo insuficiente"
	}
}

// Múltiplos retornos
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "O valor do depósito não é válido", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	if valor < c.saldo && valor > 0 {
		c.saldo -= valor
		contaDestino.Depositar(valor)
		return true
	} else {
		return false
	}
}

func main() {

	contaDoFulano := ContaCorrente{titular: "Fulano", saldo: 550.25}
	contaDoSicrano := ContaCorrente{titular: "Sicrano", saldo: 2000.82}

	fmt.Println(contaDoFulano)
	fmt.Println(contaDoSicrano)

	status := contaDoFulano.Transferir(2500.00, &contaDoSicrano)
	fmt.Println(status)
	fmt.Println(contaDoFulano)
	fmt.Println(contaDoSicrano)

	status = contaDoFulano.Transferir(25.00, &contaDoSicrano)
	fmt.Println(status)
	fmt.Println(contaDoFulano)
	fmt.Println(contaDoSicrano)

}
