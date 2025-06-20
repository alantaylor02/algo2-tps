package lista

type nodo[T any] struct {
	dato T
	prox *nodo[T]
}

type listaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
	largo   int
}

type iteradorLE[T any] struct {
	actual   *nodo[T]
	anterior *nodo[T]
	lista    *listaEnlazada[T]
}

func crearNodo[T any](elemento T) *nodo[T] {
	return &nodo[T]{elemento, nil}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{nil, nil, 0}
}

func (l listaEnlazada[T]) panicSiVacia() {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
}

func (l listaEnlazada[T]) EstaVacia() bool {
	return l.primero == nil
}

func (l *listaEnlazada[T]) InsertarPrimero(elem T) {
	nodoCreado := crearNodo(elem)
	if l.EstaVacia() {
		l.ultimo = nodoCreado
	}
	nodoCreado.prox = l.primero
	l.primero = nodoCreado
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(elem T) {
	nodoCreado := crearNodo(elem)
	if l.EstaVacia() {
		l.primero = nodoCreado
	} else {
		l.ultimo.prox = nodoCreado
	}
	l.ultimo = nodoCreado
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	l.panicSiVacia()
	datoBorrado := l.primero.dato
	l.primero = l.primero.prox
	l.largo--
	if l.primero == nil {
		l.ultimo = nil
	}
	return datoBorrado
}

func (l listaEnlazada[T]) VerPrimero() T {
	l.panicSiVacia()
	return l.primero.dato
}

func (l listaEnlazada[T]) VerUltimo() T {
	l.panicSiVacia()
	return l.ultimo.dato
}

func (l listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	nodoActual := l.primero
	for nodoActual != nil {
		if !visitar(nodoActual.dato) {
			break
		}
		nodoActual = nodoActual.prox
	}
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorLE[T]{l.primero, nil, l}
}

func (iter iteradorLE[T]) panicSiNoHaySiguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}
func (iter iteradorLE[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter iteradorLE[T]) VerActual() T {
	iter.panicSiNoHaySiguiente()
	return iter.actual.dato
}

func (iter *iteradorLE[T]) Siguiente() {
	iter.panicSiNoHaySiguiente()
	iter.anterior = iter.actual
	iter.actual = iter.actual.prox
}

func (iter *iteradorLE[T]) Insertar(elem T) {
	nodoCreado := crearNodo(elem)
	if iter.actual == nil {
		iter.lista.ultimo = nodoCreado
	}
	if iter.anterior == nil {
		iter.lista.primero = nodoCreado
	} else {
		iter.anterior.prox = nodoCreado
	}
	nodoCreado.prox = iter.actual
	iter.actual = nodoCreado
	iter.lista.largo++
}

func (iter *iteradorLE[T]) Borrar() T {
	iter.panicSiNoHaySiguiente()
	datoBorrado := iter.actual.dato
	if iter.actual.prox == nil {
		iter.lista.ultimo = iter.anterior
	}
	if iter.anterior != nil {
		iter.anterior.prox = iter.actual.prox
	} else {
		iter.lista.primero = iter.actual.prox
	}
	iter.actual = iter.actual.prox
	iter.lista.largo--
	return datoBorrado
}
