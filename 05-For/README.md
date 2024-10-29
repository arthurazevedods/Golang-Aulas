# Loops com For

Este código é baseado na documentação oficial do Golang e demonstra como realizar loops com o for

Os códigos a seguir ilustram várias formas de usar loops `for` em Go, incluindo incrementos, iterações e controles de fluxo.

## 1. Incrementando uma Variável

O primeiro loop inicializa uma variável `i` e a incrementa até 3:

```go
i := 1
for i <= 3 {
    i = i + 1
}
fmt.Println("i é igual a:", i)
```
- Inicializa i com 1.
- O loop continua enquanto i for menor ou igual a 3.
- O valor de i é impresso ao final, que será 4 após a conclusão do loop.

## 2. Loop Simples

O segundo loop imprime os números de 0 a 2:

```go
for j := 0; j < 3; j++ {
    fmt.Println(j)
}
```

- Este loop inicializa j como 0 e incrementa até que j seja menor que 3.
- Os valores 0, 1 e 2 são impressos na saída.

## 3. Iterando com `range`

O terceiro loop utiliza o `range` para iterar sobre uma estrutura. O código atual contém um erro, pois `range` deve ser aplicado a um slice ou array:

```go
for i := range [3]int{} { // Corrigido para iterar sobre um slice de 3 elementos
    fmt.Println("range", i)
}
```
- Corrigido para usar range em um array de 3 elementos.
- Os índices 0, 1 e 2 são impressos na saída.

## 4. Loop Infinito
Um loop infinito que é interrompido imediatamente:
```go
for {
    fmt.Println("loop")
    break // Sai do loop após a primeira iteração
}
```
- O loop continua indefinidamente, mas é interrompido pela instrução break, imprimindo "loop" uma vez.

## 5. Filtrando Números Ímpares
O último loop itera sobre um slice de 6 elementos, imprimindo apenas os números ímpares:
```go
for n := range [6]int{} { // Corrigido para iterar sobre um slice de 6 elementos
    if n%2 == 0 {
        continue
    }
    fmt.Println(n)
}
```
- Corrigido para usar range em um array de 6 elementos.
- O loop verifica se n é par; se for, a iteração é pulada, imprimindo apenas os números ímpares.

## Como Executar o código

```bash
go run for.go
```