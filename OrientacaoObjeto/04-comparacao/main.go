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

	contaArthur2 := ContaCorrente{
		titular:       "Arthur",
		numeroAgencia: 589,
		numeroConta:   214568,
		saldo:         325.5,
	}

	fmt.Println(contaArthur)
	fmt.Println(contaArthur2)
	fmt.Println(contaArthur == contaArthur2) //true
	contaArthur2.numeroAgencia = 140
	fmt.Println(contaArthur == contaArthur2) //false

	var contaDoFulano *ContaCorrente
	contaDoFulano = new(ContaCorrente)
	contaDoFulano.titular = "Fulano"
	contaDoFulano.saldo = 500

	var contaDoFulano2 *ContaCorrente
	contaDoFulano2 = new(ContaCorrente)
	contaDoFulano2.titular = "Fulano"
	contaDoFulano2.saldo = 500

	fmt.Println(contaDoFulano == contaDoFulano2) //false
	//é falso por terem endereços diferentes
	fmt.Println(contaDoFulano)
	fmt.Println(contaDoFulano2)

	fmt.Println(&contaDoFulano)  //endereço
	fmt.Println(&contaDoFulano2) //endereço

	fmt.Println(*contaDoFulano == *contaDoFulano2)

}
