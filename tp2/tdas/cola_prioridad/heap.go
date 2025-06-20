package cola_prioridad

const (
	CAP_INICIAL        int = 16
	FACTOR_REDIMENSION int = 2
	CANT_MINIMA        int = 4
)

type heap[T comparable] struct {
	datos    []T
	cantidad int
	cmp      func(T, T) int
}

func CrearHeap[T comparable](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	datos := make([]T, CAP_INICIAL, CAP_INICIAL)
	return &heap[T]{datos, 0, funcion_cmp}
}

func CrearHeapArr[T comparable](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	copiaArreglo := make([]T, len(arreglo))
	copy(copiaArreglo, arreglo)
	heapify(copiaArreglo, funcion_cmp)
	return &heap[T]{copiaArreglo, len(copiaArreglo), funcion_cmp}
}

func (h *heap[T]) redimensionarHeap(nuevoTam int) {
	nuevoArr := make([]T, nuevoTam, nuevoTam)
	copy(nuevoArr, h.datos)
	h.datos = nuevoArr
}

func (h heap[T]) EstaVacia() bool {
	return h.cantidad == 0
}

func upheap[T comparable](arr []T, i int, cmp func(T, T) int) {
	actual := i
	padre := (i - 1) / 2
	if padre >= 0 && cmp(arr[actual], arr[padre]) > 0 {
		actual = padre
		arr[actual], arr[i] = arr[i], arr[actual]
		upheap(arr, actual, cmp)
	}
}

func (h *heap[T]) Encolar(dato T) {
	if h.cantidad == 0 {
		h.redimensionarHeap(CAP_INICIAL)
	}
	if h.cantidad == cap(h.datos) {
		h.redimensionarHeap(cap(h.datos) * FACTOR_REDIMENSION)
	}
	h.datos[h.cantidad] = dato
	h.cantidad++
	upheap(h.datos, h.cantidad-1, h.cmp)
}

func (h heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	return h.datos[0]
}

func downheap[T comparable](arr []T, i int, ultimo int, cmp func(T, T) int) {
	if i >= ultimo {
		return
	}
	padre := i
	hijoIzq := 2*i + 1
	hijoDer := 2*i + 2
	if hijoIzq < ultimo && cmp(arr[hijoIzq], arr[padre]) > 0 {
		padre = hijoIzq
	}
	if hijoDer < ultimo && cmp(arr[hijoDer], arr[padre]) > 0 {
		padre = hijoDer
	}
	if padre != i {
		arr[i], arr[padre] = arr[padre], arr[i]
		downheap(arr, padre, ultimo, cmp)
	}
}

func heapify[T comparable](arr []T, cmp func(T, T) int) {
	for i := len(arr) - 1; i >= 0; i-- {
		downheap(arr, i, len(arr), cmp)
	}
}

func (h *heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	datoDesencolado := h.datos[0]
	h.datos[0], h.datos[h.cantidad-1] = h.datos[h.cantidad-1], h.datos[0]
	h.cantidad--
	downheap(h.datos, 0, h.cantidad, h.cmp)
	if h.cantidad*CANT_MINIMA <= cap(h.datos) && cap(h.datos) > CAP_INICIAL {
		h.redimensionarHeap(cap(h.datos) / FACTOR_REDIMENSION)
	}
	return datoDesencolado
}

func (h heap[T]) Cantidad() int {
	return h.cantidad
}

func HeapSort[T comparable](elementos []T, funcion_cmp func(T, T) int) {
	heapify[T](elementos, funcion_cmp)
	for ultimo_relativo := len(elementos) - 1; ultimo_relativo > 0; ultimo_relativo-- {
		elementos[0], elementos[ultimo_relativo] = elementos[ultimo_relativo], elementos[0]
		downheap(elementos, 0, ultimo_relativo, funcion_cmp)
	}
}
