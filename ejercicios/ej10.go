package ejercicios

import (
	"strconv"
)

// Escribir un método recursivo que dado un número
// entero decimal devuelva su equivalente en
// binario en forma de string
func DecimalABinario(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 0 {
		return "0"
	}
	if n > 1 {
		return "" + string(DecimalABinario(n/2)+strconv.Itoa(n%2))
	}
	return "0"
}
