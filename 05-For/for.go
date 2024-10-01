package main

import "fmt"

func main() {
	// Inicializa uma variável 'i' com o valor 1 e incrementa até 3
	i := 1
	for i <= 3 {
		// Incrementa 'i' em 1 em cada iteração
		i = i + 1
	}
	fmt.Println("i é igual a: ", i)

	// Loop que imprime os números de 0 a 2
	fmt.Println("For do J:")
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// Itera sobre um slice implícito de 3 elementos (índices 0, 1, 2)
	fmt.Println("For com range:")
	for i := range 3 {
		fmt.Println("range", i)
	}

	// Loop infinito que é interrompido imediatamente
	for {
		fmt.Println("loop")
		break // Sai do loop após a primeira iteração
	}

	// Itera sobre um slice de 6 elementos, imprimindo apenas os números ímpares
	for n := range 6 {
		// Pula para a próxima iteração se 'n' for par
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
