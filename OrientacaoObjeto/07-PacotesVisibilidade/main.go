package main

import (
	c "exercicio/contas"
	"fmt"
)

func main() {

	contaDoFulano := c.ContaCorrente{Titular: "Fulano", Saldo: 550.25}
	contaDoSicrano := c.ContaCorrente{Titular: "Sicrano", Saldo: 2000.82}

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
