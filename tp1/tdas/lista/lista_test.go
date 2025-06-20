package lista_test

import (
	"github.com/stretchr/testify/require"
	TDALista "tdas/lista"
	"testing"
)

// Probamos las condiciones que demuestran el correcto funcionamiento de una lista vacia.
func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}

func TestConUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 1, lista.VerUltimo())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 1, lista.VerUltimo())
	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}

// Pruebo que puedo insertar y borrar correctamente, cadenas. Se comprueba que sea cumplan las condiciones correctas en todos los casos.
func TestInsertarBorrarEnteros(t *testing.T) {
	listadeEnteros := TDALista.CrearListaEnlazada[int]()
	const maximoPrueba = 10
	for i := 0; i < maximoPrueba; i++ {
		listadeEnteros.InsertarPrimero(i)
		require.EqualValues(t, i, listadeEnteros.VerPrimero())
		require.EqualValues(t, i, listadeEnteros.VerUltimo())
		require.EqualValues(t, i, listadeEnteros.BorrarPrimero())
	}
	require.PanicsWithValue(t, "La lista esta vacia", func() { listadeEnteros.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listadeEnteros.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listadeEnteros.VerUltimo() })
}

// Pruebo que puedo insertar y borrar correctamente, enteros. Se comprueba que sea cumplan las condiciones correctas en todos los casos.
func TestInsertarBorrarCadenas(t *testing.T) {
	listaDeCadenas := TDALista.CrearListaEnlazada[string]()
	cadenas := []string{"manuel", "carlos", "sandro", "braulio", "felipe"}
	for i := range cadenas {
		listaDeCadenas.InsertarPrimero(cadenas[i])
		require.EqualValues(t, cadenas[i], listaDeCadenas.VerPrimero())
		require.EqualValues(t, cadenas[i], listaDeCadenas.VerUltimo())
		require.EqualValues(t, cadenas[i], listaDeCadenas.BorrarPrimero())
	}
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaDeCadenas.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaDeCadenas.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaDeCadenas.VerUltimo() })
}

// Prueba de caso borde: primero se insertan MUCHOS elementos en el Primera posicion
// validando que sean correctos y luego se borran todos a la vez,
// Al final la lista debe quedar vacía.
func TestVolumenInsertarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	var cantidadElegida = 1000
	for x := 0; x < cantidadElegida; x++ {
		lista.InsertarPrimero(x)
		require.EqualValues(t, x, lista.VerPrimero())
	}
	require.EqualValues(t, 999, lista.VerPrimero())
	require.EqualValues(t, 0, lista.VerUltimo())
	for x := cantidadElegida - 1; x >= 0; x-- {
		require.EqualValues(t, x, lista.BorrarPrimero())
	}
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

// Prueba de caso borde: primero se insertan MUCHOS elementos en el Ultima posicion
// validando que sean correctos y luego se borran todos a la vez,
// Al final la lista debe quedar vacía.
func TestVolumenInsertarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	var cantidadElegida = 1000
	for x := 0; x < cantidadElegida; x++ {
		lista.InsertarUltimo(x)
		require.EqualValues(t, x, lista.VerUltimo())
	}
	require.EqualValues(t, 0, lista.VerPrimero())
	require.EqualValues(t, 999, lista.VerUltimo())
	for x := 0; x < cantidadElegida; x++ {
		require.EqualValues(t, x, lista.BorrarPrimero())
	}
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestItExtInsertarSobreVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(1)
	require.EqualValues(t, 1, iter.VerActual())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 1, lista.VerUltimo())
}
func TestItExtInsertarPrincipioSobreNoVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(2)
	lista.InsertarUltimo(3)
	iter := lista.Iterador()
	iter.Insertar(1)
	require.EqualValues(t, 1, iter.VerActual())
	require.EqualValues(t, 1, lista.VerPrimero())
}

func TestItExtInsertarFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	const prueba10elem int = 10
	for x := 0; x < prueba10elem; x++ {
		lista.InsertarUltimo(x)
	}
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	iter.Insertar(11)
	require.EqualValues(t, 11, iter.VerActual())
	require.EqualValues(t, 11, lista.VerUltimo())
}

func TestItExtInsertarMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	const prueba10elem int = 10
	for x := 0; x < prueba10elem; x++ {
		lista.InsertarUltimo(x)
	}
	iter := lista.Iterador()
	for x := 0; x < 5; x++ {
		if iter.HaySiguiente() {
			iter.Siguiente()
		}
	}
	iter.Insertar(11)
	require.EqualValues(t, 11, iter.VerActual())
	require.EqualValues(t, 9, lista.VerUltimo())  // Aca comprobamos que el iterador NO inserto al final.
	require.EqualValues(t, 0, lista.VerPrimero()) // Aca comprobamos que el iterador NO inserto al principio.
}

func TestItExtBorrarSobreVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() }) //El iterador no puede borrar por que la lista esta vacia.
}
func TestItExtBorrarPrincipioSobreNoVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iter := lista.Iterador()
	require.EqualValues(t, 1, iter.Borrar())
	require.EqualValues(t, 2, lista.VerPrimero())
}

func TestItExtBorrarFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iter := lista.Iterador()
	for x := 0; x < 2; x++ {
		if iter.HaySiguiente() {
			iter.Siguiente()
		}
	}
	require.EqualValues(t, 3, iter.Borrar())
	require.EqualValues(t, 2, lista.VerUltimo())
}

func TestItExtBorrarSobreListaUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	iter := lista.Iterador()
	iter.Borrar()
	require.True(t, lista.EstaVacia())
}

func TestItExtComprobarEliminacionELementoMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iter := lista.Iterador()
	for x := 0; x < 1; x++ {
		if iter.HaySiguiente() {
			iter.Siguiente()
		}
	}
	require.EqualValues(t, 2, iter.Borrar())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())
}

func TestItExtTerminoIteracion(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() }) //Ya que estaria parado en el NIL
}

func TestItExtVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	var cantidadElegida = 1000
	for x := 0; x < cantidadElegida; x++ {
		lista.InsertarUltimo(x)
	}
	var chequedor int
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		require.EqualValues(t, chequedor, iter.VerActual())
		chequedor++
		iter.Siguiente()
	}
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
}

func TestItIntFuncionamiento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	contador := 0
	contador_ptr := &contador
	lista.Iterar(func(v int) bool {
		*contador_ptr += v
		return true
	})
	require.EqualValues(t, 11, contador)

}
func TestItIntCondicionDeCorteFalse(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(10)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	contador := 0
	contador_ptr := &contador
	lista.Iterar(func(v int) bool {
		if v == 3 {
			return false
		}
		*contador_ptr += v
		return true
	})
	require.EqualValues(t, 12, contador)
}
