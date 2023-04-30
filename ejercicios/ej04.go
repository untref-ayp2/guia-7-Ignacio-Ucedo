package ejercicios

// Escriba un método recursivo que tome una cadena y
// devuelva como resultado la cadena invertida.
func Invertir(cadena string) string {
	if len(cadena) == 1 {
		return cadena
	}

	return Invertir(cadena[1:]) + string(cadena[0])
}
