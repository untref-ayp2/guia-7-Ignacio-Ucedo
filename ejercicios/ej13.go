package ejercicios

// Escriba un método recursivo para determinar si
// un elemento está en un arreglo utilizando búsqueda binaria.
func BusquedaBinaria(arreglo []int, elemento int) bool {
	if len(arreglo) == 0 {
		return false
	}

	medio := arreglo[len(arreglo)/2]

	if elemento < medio {
		arreglo = arreglo[:len(arreglo)/2]
	} else if elemento > medio {
		arreglo = arreglo[(len(arreglo)/2)+1:]
	} else {
		return true
	}

	return BusquedaBinaria(arreglo, elemento)
}
