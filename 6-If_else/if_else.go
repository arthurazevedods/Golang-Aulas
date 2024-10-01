package main

import "fmt"

func main() {
	variavel := 7

	if variavel%2 == 0 {
		fmt.Println(variavel, " é par")
	} else {
		fmt.Println(variavel, " é ímpar")
	}

	if variavel%4 == 0 {
		fmt.Println(variavel, " é divisível por 4")
	}

	if variavel == 8 || variavel == 7 {
		fmt.Println("ou a variável é igual a 8 ou é igual a 7")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "É negativo")
	} else if num < 10 {
		fmt.Println(num, "tem um dígito")
	} else {
		fmt.Println(num, "tem múltiplos dígitos")
	}
}
