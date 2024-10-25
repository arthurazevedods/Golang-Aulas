package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// Definindo uma string constante com caracteres tailandeses.
	const s = "สวัสดี" // "Olá" em tailandês.

	// Exibe o comprimento da string em bytes.
	fmt.Println("Len:", len(s))
	// 'len' retorna o número de bytes, não o número de caracteres.
	// Como os caracteres Unicode podem usar mais de 1 byte, o valor retornado
	// pode ser maior do que o número de caracteres visíveis.

	// Itera sobre os bytes da string e imprime o valor hexadecimal de cada byte.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// Exibe a quantidade de runas (caracteres Unicode) na string.
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	// A função `utf8.RuneCountInString` conta o número real de caracteres,
	// considerando que cada runa pode ter mais de 1 byte.

	// Itera sobre as runas (caracteres) da string.
	for idx, runeValue := range s {
		// Exibe o valor Unicode do caractere e o índice em que ele começa.
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	// Loop manual usando DecodeRuneInString para decodificar runas.
	for i, w := 0, 0; i < len(s); i += w {
		// Decodifica a próxima runa da string, retornando seu valor e largura em bytes.
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		// Exibe o valor da runa e o índice onde ela começa.
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		// Atualiza `w` com a largura em bytes da runa, para avançar no loop.
		w = width

		// Chama a função examineRune para examinar cada runa.
		examineRune(runeValue)
	}
}

// Função que examina uma runa e verifica se ela corresponde a certos valores.
func examineRune(r rune) {
	// Verifica se a runa é o caractere 't' (latino).
	if r == 't' {
		fmt.Println("found tee")
		// Verifica se a runa é 'ส' (uma letra do alfabeto tailandês).
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
