package ejercicios

// Sean as y bs dos listas de enteros de tamaño n.
// Escribir una función que reciba como parámetros
// as, bs y un entero x y decida si x se puede
// escribir como suma de un elemento de as más un
// elemento de bs. O(nlogn) pero solo funciona con bs ordenado
func SumaElementos(as, bs []int, x int) bool {

	if len(as) == 0 || len(bs) == 0 {
		return false
	}

	medio := bs[len(bs)/2]

	if medio > x-as[0] {
		bs = bs[:len(bs)/2]
	} else if medio < x-as[0] {
		bs = bs[(len(bs)/2)+1:]
	} else {
		return true
	}

	return SumaElementos(as[1:], bs, x)
}
