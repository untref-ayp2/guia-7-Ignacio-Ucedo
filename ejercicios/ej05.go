package ejercicios

// Escribir un método que calcule recursivamente cuántos
// elementos hay en una pila y no altere el contenido de
// la misma. La pila sólo tiene los métodos Push, Pop y isEmpty.
func CantidadDeElementos(pila Stack) int {
	if !pila.IsEmpty() {
		pila.Pop()
	} else {
		return 0
	}

	return CantidadDeElementos(pila) + 1
}
