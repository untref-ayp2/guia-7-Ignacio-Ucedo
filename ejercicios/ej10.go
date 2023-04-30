package ejercicios

import "fmt"

// Escribir un método recursivo que dado un número
// entero decimal devuelva su equivalente en
// binario en forma de string
func DecimalABinario(n int) string {

	if n == 0 {
		return fmt.Sprint(0)
	}

	if n < 2 {
		return fmt.Sprint(1)
	}

	resto := n % 2

	return DecimalABinario(n/2) + fmt.Sprint(resto)
}
