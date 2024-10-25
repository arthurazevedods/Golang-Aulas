package main

import "fmt"

// Função recursiva para calcular o fatorial de um número
func fact(n int) int {
	// Caso base: fatorial de 0 é 1
	if n == 0 {
		return 1
	}
	// Caso recursivo: n * fatorial de n-1
	return n * fact(n-1)
}

func main() {
	// Exibe o fatorial de 7
	fmt.Println(fact(7))

	// Declara uma variável fib que armazena uma função anônima para calcular Fibonacci
	var fib func(n int) int

	// Define a função recursiva para calcular o enésimo número de Fibonacci
	fib = func(n int) int {
		// Caso base: fib(0) = 0 e fib(1) = 1
		if n < 2 {
			return n
		}
		// Caso recursivo: fib(n-1) + fib(n-2)
		return fib(n-1) + fib(n-2)
	}

	// Exibe o n número da sequência de Fibonacci
	fmt.Println(fib(1))

	// Vendo os números do cálculo fatorial com um For
	for i := range 8 {
		fmt.Println("O fatorial de ", i, " é:", fact(i))
	}
	fmt.Println()
	// Vendo os números de fibonacci com um For
	for i := range 8 {
		fmt.Println("Número ", i, " da sequência de Fibonacci:", fib(i))
	}

}
