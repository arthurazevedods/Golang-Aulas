# Estruturas Condicionais e Loops em Go

Este código demonstra o uso de estruturas condicionais e loops em Go, com exemplos que ilustram a lógica de controle de fluxo.

Os códigos a seguir mostram como usar `if`, `else`, e `for` para executar diferentes operações.

## 1. Incrementando uma Variável

O primeiro trecho inicializa uma variável `i` e a incrementa até 3:

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
fmt.Println("For do J:")
for j := 0; j < 3; j++ {
    fmt.Println(j)
}
``` 
- Este loop inicializa j como 0 e incrementa até que j seja menor que 3.
- Os valores 0, 1 e 2 são impressos na saída.