# Manipulação de Strings Unicode em Go

Este código é baseado na documentação oficial do Golang e demonstra como trabalhar com strings Unicode no Go, especificamente com a codificação UTF-8. O código exemplo utiliza uma string em tailandês ("สวัสดี", que significa "Olá") e mostra como:

- Contar o número de bytes e runas (caracteres Unicode).
- Iterar sobre os bytes e runas da string.
- Usar a função `utf8.DecodeRuneInString` para decodificar runas manualmente.
- Comparar runas específicas com valores predefinidos.

## Funcionamento do Código

### Principais Componentes

- **`len(s)`**: Retorna o número de bytes que a string ocupa. Em strings Unicode, um caractere pode ocupar mais de 1 byte, por isso o valor retornado pode ser maior que o número de caracteres visíveis.
  
- **`utf8.RuneCountInString(s)`**: Retorna o número real de caracteres (runes) na string. Uma runa pode ocupar mais de 1 byte, e esta função conta corretamente os caracteres, independentemente do tamanho em bytes.

- **Iteração por bytes**: Um loop percorre os bytes da string e imprime o valor hexadecimal de cada byte, permitindo observar a codificação UTF-8.

- **Iteração por runas**: Usando `range`, o código percorre cada runa da string, imprime o valor Unicode (`%#U`) e o índice de início da runa.

- **Decodificação manual com `utf8.DecodeRuneInString`**: O código também demonstra como decodificar manualmente as runas de uma string, obtendo o valor da runa e a largura em bytes de cada caractere.

### Função `examineRune`

A função `examineRune(r rune)` examina cada runa da string e compara com valores predefinidos:
- Se a runa for `'t'`, imprime `"found tee"`.
- Se a runa for `'ส'` (um caractere tailandês), imprime `"found so sua"`.

Esta função mostra como é possível comparar runas individuais, permitindo realizar verificações em strings Unicode de forma simples.

### Estrutura do Código

O código é organizado da seguinte maneira:

- **`main()`**: A função principal, onde a string é definida e as operações de contagem, iteração e decodificação são realizadas.
- **`examineRune(r rune)`**: Função auxiliar para verificar runas específicas e realizar ações baseadas nas comparações.

## Como Rodar o Código

1. Certifique-se de que você tenha o Go instalado. Você pode verificar se o Go está instalado executando o seguinte comando no terminal:

    ```bash
    go version
    ```

2. Clone ou baixe este repositório.

3. Execute o código Go usando o seguinte comando:

    ```bash
    go run main.go
    ```

## Saída Esperada

Quando o código é executado, ele exibe:

1. O número de bytes da string.
2. Os valores hexadecimais dos bytes da string.
3. O número de runas (caracteres Unicode).
4. Os valores Unicode e os índices onde cada runa começa.
5. Decodificação manual de runas, mostrando seus valores e índices.
6. Verificações específicas para as runas `'t'` e `'ส'`.

### Exemplo de Saída

```bash
Len: 18
e0 b8 aa e0 b8 a7 e0 b8 b1 e0 b8 aa e0 b8 94 e0 b8 b5 
Rune count: 6
U+0E2A starts at 0
U+0E27 starts at 3
U+0E31 starts at 6
U+0E2A starts at 9
U+0E14 starts at 12
U+0E35 starts at 15

Using DecodeRuneInString
U+0E2A starts at 0
found so sua
U+0E27 starts at 3
U+0E31 starts at 6
U+0E2A starts at 9
found so sua
U+0E14 starts at 12
U+0E35 starts at 15
