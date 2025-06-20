package lista

type IteradorLista[T any] interface {

	// VerActual muestra el elemento al que apunta el iterador en ese momento. Si el iterador termino la iteracion
	// entra en panico con el mensaje "El iterador termino de iterar"
	VerActual() T

	// HaySiguiente verifica si el iterador esta parado sobre un elemento o sobre NIL.
	HaySiguiente() bool

	// Siguiente mueve la posicion a la que apunta el iterador a la posicion posterior. Si el iterador termino la iteracion
	// entra en panico con el mensaje "El iterador termino de iterar"
	Siguiente()

	// Insertar coloca un elemento en la posicion en la que apunte el iterador.
	Insertar(T)

	// Borrar elimina el elemento de la lista al que apunte el iterador en ese momento. Si el iterador termino la iteracion
	// entra en panico con el mensaje "El iterador termino de iterar"
	Borrar() T
}

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos insertados, false en caso contrario
	EstaVacia() bool

	// InsertarPrimero agrega un nuevo elemento a la lista al principio
	InsertarPrimero(T)

	// InsertarUltimo agrega un nuevo elemento a la lista al final
	InsertarUltimo(T)

	// BorrarPrimero saca el primer elemento de la lista y lo devuelve. Si la lista está vacía, entra
	// en pánico con el mensaje "La lista esta vacia"
	BorrarPrimero() T

	// VerPrimero obtiene el valor del primero de la lista. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	VerPrimero() T

	// VerUltimo obtiene el valor del ultimo de la lista. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos que tiene la lista.
	Largo() int

	// Iterar permite ir por todos los elementos de la lista aplicándole a cada uno la funcion "visitar" pasada
	// por parámetro, hasta que la misma devuelva false o la lista se haya iterado por completo.
	Iterar(visitar func(T) bool)

	// Iterador devuelve un Struct que cumple con la interfaz del iterador externo.
	Iterador() IteradorLista[T]
}
