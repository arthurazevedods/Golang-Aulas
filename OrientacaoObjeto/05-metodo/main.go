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
func main() {

	contaDoFulano := ContaCorrente{}
	contaDoFulano.titular = "Fulano"
	contaDoFulano.saldo = 500

	// Exibe o saldo inicial da conta
	fmt.Println(contaDoFulano.saldo)

	// Realiza um saque de 400 e exibe o resultado
	fmt.Println(contaDoFulano.Sacar(400)) //Vai exibir a mensagem de retorno

	// Exibe o saldo atualizado após o saque
	fmt.Println(contaDoFulano.saldo)

	// Realiza o depósito de 100
	//fmt.Println(contaDoFulano.Depositar(100))
	status, saldo := contaDoFulano.Depositar(100)
	fmt.Println(status, " - Saldo atual:", saldo)
	// Exibe o saldo atualizado após o depósito
	fmt.Println(contaDoFulano.saldo)
}
