package main

import "fmt"

// Função que recebe um valor inteiro e tenta modificar seu valor para 0
// No entanto, essa alteração só afeta a cópia local da variável (passagem por valor)
func zerolocal(ival int) {
	ival = 0
	fmt.Println("zerolocal dentro da função:", ival)
}

// Função que recebe um ponteiro para um inteiro e altera o valor no endereço de memória
// Nesse caso, a alteração afeta a variável original (passagem por referência)
func zeroponteiro(iptr *int) {
	*iptr = 0 // Dereferência do ponteiro para modificar o valor no endereço apontado
}

func main() {
	i := 1                     // Define a variável i com o valor 1
	fmt.Println("initial:", i) // Exibe o valor inicial de i

	zerolocal(i)                 // Passa uma cópia do valor de i (passagem por valor)
	fmt.Println("zerolocal:", i) // O valor de i não é alterado, pois apenas uma cópia foi passada

	zeroponteiro(&i)                // Passa o endereço de i (passagem por referência)
	fmt.Println("zeroponteiro:", i) // O valor de i é alterado para 0, pois foi modificado diretamente na memória

	fmt.Println("pointer:", &i) // Exibe o endereço de memória de i
}
