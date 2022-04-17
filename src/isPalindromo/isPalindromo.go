package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var lowerText = strings.ToLower(text)
	var textReverse string
	for i := len(lowerText) - 1; i >= 0; i-- {
		textReverse += string(lowerText[i])
	}
	if lowerText == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func main() {
	isPalindromo("Ama")
}
