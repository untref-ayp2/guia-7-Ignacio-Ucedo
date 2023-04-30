package ejercicios

// Escribir un método recursivo que devuelva el
// cociente y el resto de la división entera mediante
// restas sucesivas
func DivisionEntera(dividendo, divisor int) (cociente, resto int) {
	if dividendo < divisor {
		return 0, dividendo
	}

	c, _ := DivisionEntera(dividendo-divisor, divisor)
	c++
	return c, 0

}
