package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("Escreva ", i, " as ")
	switch i {
	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("três")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("É final de semana")
	default:
		fmt.Println("É dia de semana")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Antes de meio-dia")
	default:
		fmt.Println("Depois de meio-dia")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("É um booleano")
		case int:
			fmt.Println("É um int")
		case string:
			fmt.Println("É uma string")
		default:
			fmt.Printf("Não reconheço o tipo %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("Olá")
	whatAmI(44.4)

	cardapio()
}

func cardapio() {
	var prato string
	fmt.Println("Digite seu prato preferido...")
	fmt.Println("P - Pizza")
	fmt.Println("H - Hambúrguer")
	fmt.Println("B - Bife com fritas")
	fmt.Println("S - Salada Caesar")
	fmt.Println("F - Salada de Frutas")
	fmt.Println("E - Estrogonofe")
	fmt.Println("O - Outros")
	fmt.Scan(&prato)

	switch prato {
	case "B":
		fmt.Println("Com batatas Palito ou Noisete?")
	case "H":
		fmt.Println("Com Queijo ou com Ovo?")
	case "P":
		fmt.Println("Calabresa ou Napolitana?")
	case "S":
		fmt.Println("Alface ou Rúcula?")
	case "F":
		fmt.Println("Kiwi ou Frango?")
	case "E":
		fmt.Println("Carne ou Frango?")
	case "O":
		fmt.Println("Não gostou de nosso cardápio?")
	default:
		fmt.Println("Não entendi seu paladar...")
	}
}
