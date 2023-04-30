package ejercicios

// Escriba un método recursivo que tome dos
// números enteros y calcule la multiplicación
// entre ellos, usando sólo sumas
func Multiplicar(a, b int) int {
	if b == 0 {
		return 0
	}

	return Multiplicar(a, b-1) + a
}
