package cola_test

import (
	"github.com/stretchr/testify/require"
	TDACola "tdas/cola"
	"testing"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestAlgunosElementos(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	// Encolo algunos elementos y veo que, al final, el primero sea el
	// primer elemento encolado (FIFO).
	cola.Encolar(1)
	cola.Encolar(2)
	cola.Encolar(3)
	require.EqualValues(t, 1, cola.VerPrimero())
	// Desencolo y veo que se vaya desencolando el elemento correcto.
	require.EqualValues(t, 1, cola.Desencolar())
	require.EqualValues(t, 2, cola.Desencolar())
	require.EqualValues(t, 3, cola.Desencolar())
	// Por ultimo, veo que la cola quede vacia.
	require.True(t, cola.EstaVacia())
}

func TestEncolarYDesencolarHastaVaciar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	// Encolo y desencolo elementos.
	cola.Encolar(1)
	cola.Encolar(2)
	cola.Encolar(3)
	cola.Desencolar()
	cola.Desencolar()
	cola.Desencolar()
	// Compruebo que la cola queda vacia.
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestPilaDeCadenas(t *testing.T) {
	colaStrings := TDACola.CrearColaEnlazada[string]()
	// Encolo distintos strings y veo que, al final, el primero sea el
	// primer string encolado (FIFO).
	colaStrings.Encolar("a")
	colaStrings.Encolar("b")
	colaStrings.Encolar("c")
	require.EqualValues(t, "a", colaStrings.VerPrimero())
	// Desencolo y veo que se vaya desencolando el string correcto.
	require.EqualValues(t, "a", colaStrings.Desencolar())
	require.EqualValues(t, "b", colaStrings.Desencolar())
	require.EqualValues(t, "c", colaStrings.Desencolar())
	// Por ultimo, veo que la cola quede vacia.
	require.True(t, colaStrings.EstaVacia())
}

func TestVolumen(t *testing.T) {
	cantElem := 10000
	cola := TDACola.CrearColaEnlazada[int]()
	// Encolo a una cola vacia 10000 elementos y compruebo que,
	// al final, el primero sea el correspondiente.
	for i := 1; i <= cantElem; i++ {
		cola.Encolar(i)
	}
	require.EqualValues(t, 1, cola.VerPrimero())
	// Voy desencolando los elementos y compruebo que el nuevo
	// primero es el esperado.
	for i := 1; i < cantElem; i++ {
		cola.Desencolar()
		require.EqualValues(t, i+1, cola.VerPrimero())
	}
	// Desencolo el elemento que queda y compruebo que la cola
	// quede vacia.
	cola.Desencolar()
	require.True(t, cola.EstaVacia())
}
