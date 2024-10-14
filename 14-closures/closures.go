package main

import "fmt"

// Funções Anônimas e Closures

//unções anônimas são funções que não têm um nome.
//Elas são úteis quando você deseja definir uma função inline, sem precisar nomeá-la.
//Além disso, funções anônimas podem capturar variáveis do seu contexto de definição,
//formando closures.

//A função intSeq retorna outra função.
//Essa função anônima é definida dentro do corpo de intSeq.
func intSeq() func() int {
	i := 0 //a variável i é inicializada em 0.

	//A função anônima que é retornada incrementa i a cada chamada e retorna seu valor.
	//Ou seja, toda vez que essa função for chamada, ela irá incrementar i.
	return func() int {
		i++
		return i
	}
}

func main() {

	//chamamos intSeq e atribuimos o resultado (que é uma função) à variável nextInt.
	//Ou seja, toda vez que chamarmos nextInt, o i vai ser incrementado
	nextInt := intSeq()

	//Essa função nextInt captura seu próprio valor de i,
	//que será atualizado cada vez que nextInt for chamada.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
