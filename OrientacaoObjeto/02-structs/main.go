package main

import "fmt"

func main() {
	type ContaCorrente struct {
		titular       string
		numeroAgencia int
		numeroConta   int
		saldo         float64
	}

	contaArthur := ContaCorrente{
		titular:       "Arthur",
		numeroAgencia: 589,
		numeroConta:   214568,
		saldo:         325.5,
	}
	contaAzevedo := ContaCorrente{
		titular:       "Azevedo",
		numeroAgencia: 589,
		numeroConta:   256872,
		saldo:         501.25,
	}

	fmt.Println(contaArthur)
	fmt.Println(contaAzevedo)

	contaDoFulano := ContaCorrente{titular: "Fulano", saldo: 125.5}

	fmt.Println(contaDoFulano)
}
