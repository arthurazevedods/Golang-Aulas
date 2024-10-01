package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		i = i + 1
	}

	fmt.Println("i é igual a: ", i)

	fmt.Println("For do J:")
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	fmt.Println("For com range:")
	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
