package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	//vals retorna dois valores
	//o 3 será guardado em a
	//o 7 será guardado em b
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
