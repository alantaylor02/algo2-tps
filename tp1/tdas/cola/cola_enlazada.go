package cola

type nodo[T any] struct {
	dato T
	prox *nodo[T]
}

type colaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
}

func CrearColaEnlazada[T any]() Cola[T] {

	return &colaEnlazada[T]{nil, nil}
}

func (c colaEnlazada[T]) EstaVacia() bool {

	return c.primero == nil
}

func (c colaEnlazada[T]) VerPrimero() T {

	if c.EstaVacia() {
		panic("La cola esta vacia")
	}

	return c.primero.dato
}

func (c *colaEnlazada[T]) Encolar(elem T) {

	nodo := &nodo[T]{elem, nil}

	if c.EstaVacia() {
		c.primero = nodo
	} else {
		c.ultimo.prox = nodo
	}
	c.ultimo = nodo
}

func (c *colaEnlazada[T]) Desencolar() T {

	if c.EstaVacia() {
		panic("La cola esta vacia")
	}

	datoDesencolado := c.primero.dato
	c.primero = c.primero.prox

	if c.primero == nil {
		c.ultimo = nil
	}

	return datoDesencolado
}
