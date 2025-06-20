package cola_prioridad_test

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"strings"
	TDAColaPrior "tdas/cola_prioridad"
	"testing"
)

const RANGO_RANDOM int = 100000

func TestColaPrioridadVacia(t *testing.T) {
	colaPrioridad := TDAColaPrior.CrearHeap[int](func(a, b int) int { return a - b })
	require.True(t, colaPrioridad.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.Desencolar() })
	require.EqualValues(t, 0, colaPrioridad.Cantidad())
}

func TestColaPrioridadSobreArregloVacio(t *testing.T) {
	colaPrioridad := TDAColaPrior.CrearHeapArr[int]([]int{}, func(a, b int) int { return a - b })
	require.True(t, colaPrioridad.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.Desencolar() })
	require.EqualValues(t, 0, colaPrioridad.Cantidad())
}

func TestCrearHeapArrVacio(t *testing.T) {
	arr := []int{}
	colaPrioridad := TDAColaPrior.CrearHeapArr[int](arr, func(a, b int) int { return a - b })
	ordenDeDesencolado := []int{25, 10, 4, 2}
	colaPrioridad.Encolar(4)
	colaPrioridad.Encolar(10)
	colaPrioridad.Encolar(25)
	colaPrioridad.Encolar(2)
	require.EqualValues(t, 4, colaPrioridad.Cantidad())
	require.EqualValues(t, 25, colaPrioridad.VerMax())
	// Voy desencolando y corroborando que se desencola el elemento correspondiente
	for i := 0; i < 4; i++ {
		require.EqualValues(t, ordenDeDesencolado[i], colaPrioridad.VerMax())
		require.EqualValues(t, ordenDeDesencolado[i], colaPrioridad.Desencolar())
	}
	require.True(t, colaPrioridad.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.Desencolar() })
}

func TestCrearHeapArr(t *testing.T) {
	arr := []int{2, 6, 1, 3, 4, 25, 10, 15}
	colaPrioridad := TDAColaPrior.CrearHeapArr[int](arr, func(a, b int) int { return a - b })
	ordenDeDesencolado := []int{25, 15, 10, 6, 4, 3, 2, 1}
	require.EqualValues(t, 8, colaPrioridad.Cantidad())
	require.EqualValues(t, 25, colaPrioridad.VerMax())
	// Voy desencolando y corroborando que se desencola el elemento correspondiente
	for i := 0; i < 8; i++ {
		require.EqualValues(t, ordenDeDesencolado[i], colaPrioridad.VerMax())
		require.EqualValues(t, ordenDeDesencolado[i], colaPrioridad.Desencolar())
	}
	require.True(t, colaPrioridad.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.Desencolar() })
}

func TestEncolarEnHeapCreadoConArreglo(t *testing.T) {
	arr := []int{2, 6, 1, 3, 4, 25, 10, 15}
	ordenDeDesencolado := []int{30, 25, 15, 10, 7, 6, 4, 3, 2, 1}
	colaPrioridad := TDAColaPrior.CrearHeapArr[int](arr, func(a, b int) int { return a - b })
	// Chequeo que si encolo algo que no es de mayor prioridad que lo que habia en el arreglo,
	// no se me cambio el maximo de mi cola.
	colaPrioridad.Encolar(7)
	require.EqualValues(t, 25, colaPrioridad.VerMax())
	// En cambio, si encolo algo de mayor prioridad, si se modifica.
	colaPrioridad.Encolar(30)
	require.EqualValues(t, 30, colaPrioridad.VerMax())
	// Voy desencolando y corroborando que se desencola el elemento correspondiente
	for i := 0; i < 10; i++ {
		require.EqualValues(t, ordenDeDesencolado[i], colaPrioridad.VerMax())
		require.EqualValues(t, ordenDeDesencolado[i], colaPrioridad.Desencolar())
	}
	require.True(t, colaPrioridad.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.Desencolar() })
}

func TestEncolarAlgunosElementos(t *testing.T) {
	// Corrobora que, luego de encolados varios elementos, el maximo es el que corresponde
	colaPrioridad := TDAColaPrior.CrearHeap[int](func(a, b int) int { return a - b })
	colaPrioridad.Encolar(6)
	colaPrioridad.Encolar(10)
	colaPrioridad.Encolar(2)
	colaPrioridad.Encolar(11)
	colaPrioridad.Encolar(4)

	require.False(t, colaPrioridad.EstaVacia())
	require.EqualValues(t, 11, colaPrioridad.VerMax())
	require.EqualValues(t, 5, colaPrioridad.Cantidad())
}

func TestDesencolarAlgunosElementos(t *testing.T) {
	// Corrobora que, luego de encolados varios elementos, el maximo es el que corresponde
	colaPrioridad := TDAColaPrior.CrearHeap[int](func(a, b int) int { return a - b })
	colaPrioridad.Encolar(6)
	colaPrioridad.Encolar(10)
	colaPrioridad.Encolar(2)
	colaPrioridad.Encolar(11)
	colaPrioridad.Encolar(4)
	require.EqualValues(t, 11, colaPrioridad.VerMax())
	// Voy desencolando y viendo que el elemento que sale es el mayor prioridad de la cola
	require.EqualValues(t, 11, colaPrioridad.Desencolar())
	require.EqualValues(t, 10, colaPrioridad.Desencolar())
	require.EqualValues(t, 6, colaPrioridad.Desencolar())
	require.EqualValues(t, 4, colaPrioridad.Desencolar())
	require.EqualValues(t, 2, colaPrioridad.Desencolar())

	require.True(t, colaPrioridad.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.Desencolar() })
}

func TestEncolarDesencolarConStrings(t *testing.T) {
	// Corrobora que, luego de encolados varios elementos, el maximo es el que corresponde
	colaPrioridad := TDAColaPrior.CrearHeap[string](strings.Compare)
	colaPrioridad.Encolar("Z")
	colaPrioridad.Encolar("X")
	colaPrioridad.Encolar("A")
	colaPrioridad.Encolar("H")
	colaPrioridad.Encolar("L")
	require.EqualValues(t, "Z", colaPrioridad.VerMax())
	require.EqualValues(t, 5, colaPrioridad.Cantidad())

	// Voy desencolando y viendo que el elemento que sale es el mayor prioridad de la cola
	require.EqualValues(t, "Z", colaPrioridad.Desencolar())
	require.EqualValues(t, "X", colaPrioridad.Desencolar())
	require.EqualValues(t, "L", colaPrioridad.Desencolar())
	require.EqualValues(t, "H", colaPrioridad.Desencolar())
	require.EqualValues(t, "A", colaPrioridad.Desencolar())

	require.True(t, colaPrioridad.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.Desencolar() })
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(arr1, arr2 []int) []int {
	i, j := 0, 0
	resultado := []int{}
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			resultado = append(resultado, arr1[i])
			i++
		} else {
			resultado = append(resultado, arr2[j])
			j++
		}
	}
	resultado = append(resultado, arr1[i:]...)
	resultado = append(resultado, arr2[j:]...)
	return resultado
}

func TestVolumenEncolarDesencolar(t *testing.T) {
	// Voy encolando numeros aleatorios en un rango especifico 100.000 veces
	colaPrioridad := TDAColaPrior.CrearHeap[int](func(a, b int) int { return a - b })
	j := 100000
	encolados := make([]int, j)
	for i := 0; i < j; i++ {
		numero := rand.Intn(RANGO_RANDOM)
		encolados[i] = numero
		colaPrioridad.Encolar(encolados[i])
	}
	// Ordeno el arreglo para saber el orden en el que se deberian desencolar los elementos
	ordenDeDesencolado := mergeSort(encolados)
	// Desencolo y voy chequeando que se desencola correctamente
	for k := j - 1; k >= 0; k-- {
		require.EqualValues(t, ordenDeDesencolado[k], colaPrioridad.VerMax())
		require.EqualValues(t, ordenDeDesencolado[k], colaPrioridad.Desencolar())
	}
	require.True(t, colaPrioridad.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.Desencolar() })
}

func TestHeapsortSobreArregloVacio(t *testing.T) {
	arr := []int{}
	TDAColaPrior.HeapSort[int](arr, func(a, b int) int { return a - b })
	// Luego de hacer HeapSort, no deberia hacer nada y el arreglo tendria que seguir vacio.
	require.EqualValues(t, arr, []int{})
}

func TestHeapsort(t *testing.T) {
	arr := []int{6, 2, 10, 25, 1, 6, 5, 7, 21}
	TDAColaPrior.HeapSort[int](arr, func(a, b int) int { return a - b })
	// Luego de hacer HeapSort, el arreglo deberia haberse ordenado ascendentemente.
	require.EqualValues(t, []int{1, 2, 5, 6, 6, 7, 10, 21, 25}, arr)
}

func TestHeapsortStrings(t *testing.T) {
	arr := []string{"B", "A", "C", "X", "H", "D", "A"}
	TDAColaPrior.HeapSort[string](arr, strings.Compare)
	// Luego de hacer HeapSort, el arreglo deberia haberse ordenado ascendentemente.
	require.EqualValues(t, []string{"A", "A", "B", "C", "D", "H", "X"}, arr)
}
