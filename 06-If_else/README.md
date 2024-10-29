# Estruturas Condicionais em Go

Este código é baseado na documentação oficial do Golang e demonstra como usar estruturas condicionais em Go, incluindo `if` e `else`.

Os códigos a seguir ilustram várias formas de utilizar condições e como elas influenciam o fluxo do programa.

## 1. Verificando se um Número é Par ou Ímpar

O primeiro trecho verifica se uma variável é par ou ímpar:

```go
variavel := 7

if variavel%2 == 0 {
    fmt.Println(variavel, "é par")
} else {
    fmt.Println(variavel, "é ímpar")
}
```

- Inicializa a variável com 7.
- Se variavel é divisível por 2, imprime que é par.
- Caso contrário, imprime que é ímpar.
## 2. Verificando Divisibilidade por 4
O segundo trecho verifica se a variável é divisível por 4:

```go
if variavel%4 == 0 {
    fmt.Println(variavel, "é divisível por 4")
}
```
Se variavel é divisível por 4, imprime que é divisível.
## 3. Verificando Igualdade com 7 ou 8
O terceiro trecho verifica se a variável é igual a 7 ou 8:

```go
if variavel == 8 || variavel == 7 {
    fmt.Println("ou a variável é igual a 8 ou é igual a 7")
}
```
Se variavel é igual a 7 ou 8, imprime a mensagem correspondente.
##  4. Usando Declaração de Variável com if
O quarto trecho utiliza uma declaração de variável dentro de uma condição:

```go
if num := 9; num < 0 {
    fmt.Println(num, "é negativo")
} else if num < 10 {
    fmt.Println(num, "tem um dígito")
} else {
    fmt.Println(num, "tem múltiplos dígitos")
}
```
- Inicializa num como 9.
- Se num é negativo, imprime que é negativo.
- Se num tem menos de 10, imprime que tem um dígito.
- Caso contrário, imprime que tem múltiplos dígitos.
## Como Executar o Código
Para executar este código, use o seguinte comando no terminal:

bash
go run if_else.go