package main

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	rows := make([]string, numRows)
	curr := 0
	down := false

	for _, char := range s {
		rows[curr] += string(char)

		if curr == 0 || curr == numRows-1 {
			down = !down
		}
		if down {
			curr++
		} else {
			curr--
		}
	}
	return strings.Join(rows, "")
}

func main() {
	result := convert("PAYPALISHIRING", 3)
	println(result)
}
