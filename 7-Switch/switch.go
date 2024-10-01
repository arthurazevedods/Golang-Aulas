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
}
