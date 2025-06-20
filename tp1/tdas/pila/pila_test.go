package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "tdas/pila"
	"testing"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestAlgunosElementos(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	// Apilo algunos elementos y voy viendo que coincidan con el tope.
	pila.Apilar(1)
	require.EqualValues(t, 1, pila.VerTope())
	pila.Apilar(2)
	require.EqualValues(t, 2, pila.VerTope())
	pila.Apilar(3)
	require.EqualValues(t, 3, pila.VerTope())
	// Desapilo y veo que coincida con el ultimo dato apilado.
	require.EqualValues(t, 3, pila.Desapilar())
	require.EqualValues(t, 2, pila.Desapilar())
	require.EqualValues(t, 1, pila.Desapilar())
	// Por ultimo, veo que la pila quede vacia.
	require.True(t, pila.EstaVacia())
}

func TestApilarYDesapilarHastaVaciar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	// Apilo y desapilo elementos.
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	pila.Desapilar()
	pila.Desapilar()
	pila.Desapilar()
	// Compruebo que la pila queda vacia.
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestPilaDeCadenas(t *testing.T) {
	pilaStrings := TDAPila.CrearPilaDinamica[string]()
	// Apilo distintos strings y voy viendo que coinciden con el tope.
	pilaStrings.Apilar("a")
	require.EqualValues(t, "a", pilaStrings.VerTope())
	pilaStrings.Apilar("b")
	require.EqualValues(t, "b", pilaStrings.VerTope())
	pilaStrings.Apilar("c")
	require.EqualValues(t, "c", pilaStrings.VerTope())
	// Desapilo y veo que coincida con el ultimo dato apilado.
	require.EqualValues(t, "c", pilaStrings.Desapilar())
	require.EqualValues(t, "b", pilaStrings.Desapilar())
	require.EqualValues(t, "a", pilaStrings.Desapilar())
	// Por ultimo, veo que la pila quede vacia.
	require.True(t, pilaStrings.EstaVacia())
}

func TestVolumen(t *testing.T) {
	cantElem := 10000
	pila := TDAPila.CrearPilaDinamica[int]()
	// Apilo a una pila vacia 100000 elementos y compruebo que,
	// elemento a elemento, el tope sea el correspondiente.
	for i := 1; i <= cantElem; i++ {
		pila.Apilar(i)
		require.EqualValues(t, i, pila.VerTope())
	}
	// Voy desapilando los elementos y compruebo que el nuevo
	// tope es el esperado.
	for i := cantElem; i > 1; i-- {
		pila.Desapilar()
		require.EqualValues(t, i-1, pila.VerTope())
	}
	// Desapilo el elemento que queda y compruebo que la pila
	// deberia quedar vacia.
	pila.Desapilar()
	require.True(t, pila.EstaVacia())
}
