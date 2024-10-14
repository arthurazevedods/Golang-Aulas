package main

import "fmt"

// permite que você passe um número variável de argumentos para ela.
//A definição da função usa ... para indicar que pode aceitar zero ou mais argumentos de um tipo específico.
//A função sum recebe um argumento variádico nums do tipo int.
//Dentro da função, nums é tratado como um slice de inteiros ([]int).
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	// A função imprime os números recebidos e, em seguida,
	//itera sobre eles usando um loop for para calcular a soma total.
	//No final, imprime o total.
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
