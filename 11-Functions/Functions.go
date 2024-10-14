package main

import "fmt"

//Variável - tipo de dado
// O int antes das chaves indica o retorno da função
func plus(a int, b int) int {
	return a + b
}

//O tipo de dado no final dos parametros, indica o tipo de dado das variáveis que estão antes dela
// variável, variável, variável e o tipo de dado delas
func plusPlus(a, b, c int) int {
	return a + b + c
}

func plusPlusFloat(a, b, c float64) float64 {
	return a + b + c
}

func texto(str string) string {
	return "Este é o texto:" + str
}

func vazio() {
	fmt.Println("Vazio")
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	res2 := plusPlusFloat(1.1, 2.2, 3.3)
	fmt.Println("1.1+2.2+3.3 =", res2)

	fmt.Println(texto("Exemplo de string"))

	vazio()
}
