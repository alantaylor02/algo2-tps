package pila

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

const CAP_INICIAL int = 8
const FACTOR_REDIMENSION int = 2

func CrearPilaDinamica[T any]() Pila[T] {

	datos := make([]T, 0, CAP_INICIAL)
	pila := &pilaDinamica[T]{datos, 0}

	return pila
}

func (p *pilaDinamica[T]) redimensionarPila(nuevoTam int) {

	nuevoArr := make([]T, p.cantidad, nuevoTam)
	copy(nuevoArr, p.datos)
	p.datos = nuevoArr
}

func (p pilaDinamica[T]) EstaVacia() bool {

	return p.cantidad == 0
}

func (p pilaDinamica[T]) VerTope() T {

	if p.EstaVacia() {
		panic("La pila esta vacia")
	}

	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(elem T) {

	p.datos = p.datos[:p.cantidad+1]
	p.datos[p.cantidad] = elem
	p.cantidad++

	if p.cantidad == cap(p.datos) {
		p.redimensionarPila(cap(p.datos) * FACTOR_REDIMENSION)
	}
}

func (p *pilaDinamica[T]) Desapilar() T {

	if p.EstaVacia() {
		panic("La pila esta vacia")
	}

	datoDesapilado := p.datos[p.cantidad-1]
	p.cantidad--
	p.datos = p.datos[:p.cantidad]

	if p.cantidad*4 <= cap(p.datos) && cap(p.datos) > CAP_INICIAL {
		p.redimensionarPila(cap(p.datos) / FACTOR_REDIMENSION)
	}

	return datoDesapilado
}
