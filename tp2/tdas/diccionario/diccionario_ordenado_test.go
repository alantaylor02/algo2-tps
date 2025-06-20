package diccionario_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strings"
	TDADiccOrdenado "tdas/diccionario"
	"testing"
)

const RANGO_RANDOM = 1000000000000000

var TAMS_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}

func TestDiccionarioVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccOrdenado.CrearABB[string, string](strings.Compare)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestDiccionarioClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un ABB vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	dic := TDADiccOrdenado.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := TDADiccOrdenado.CrearABB[int, string](func(a, b int) int { return a - b })
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}

func TestUnElemento(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccOrdenado.CrearABB[string, int](strings.Compare)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccionarioGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccOrdenado.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccOrdenado.CrearABB[string, string](strings.Compare)
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestReemplazoDato2(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos.")

	dic := TDADiccOrdenado.CrearABB[int, int](func(a, b int) int { return a - b })
	n := 500
	claves := make([]int, 500)
	for i := 0; i < n; i++ {
		claves[i] = rand.Intn(RANGO_RANDOM)
		dic.Guardar(claves[i], i)
	}
	for i := 0; i < n; i++ {
		dic.Guardar(claves[i], 2*i)
	}
	ok := true
	for i := 0; i < n && ok; i++ {
		ok = dic.Obtener(claves[i]) == 2*i
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestDiccionarioBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccOrdenado.CrearABB[string, string](strings.Compare)

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestReutlizacionDeBorrados(t *testing.T) {
	t.Log("Prueba de caja blanca: revisa que no haya problema reinsertando un elemento borrado")
	dic := TDADiccOrdenado.CrearABB[string, string](strings.Compare)
	clave := "hola"
	dic.Guardar(clave, "mundo!")
	dic.Borrar(clave)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(clave))
	dic.Guardar(clave, "mundooo!")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "mundooo!", dic.Obtener(clave))
}

func TestConClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccOrdenado.CrearABB[int, string](func(a, b int) int { return a - b })
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

type basico struct {
	a string
	b int
}

type avanzado struct {
	w int
	x basico
	y basico
	z string
}

func compararAvanzado(a, b avanzado) int {
	return a.w - b.w
}

func TestConClavesStructs(t *testing.T) {
	t.Log("Valida que tambien funcione con estructuras mas complejas")
	dic := TDADiccOrdenado.CrearABB[avanzado, int](compararAvanzado)

	a1 := avanzado{w: 4, z: "hola", x: basico{a: "mundo", b: 8}, y: basico{a: "!", b: 10}}
	a2 := avanzado{w: 10, z: "aloh", x: basico{a: "odnum", b: 14}, y: basico{a: "!", b: 5}}
	a3 := avanzado{w: 2, z: "hello", x: basico{a: "world", b: 8}, y: basico{a: "!", b: 4}}

	dic.Guardar(a1, 0)
	dic.Guardar(a2, 1)
	dic.Guardar(a3, 2)

	require.True(t, dic.Pertenece(a1))
	require.True(t, dic.Pertenece(a2))
	require.True(t, dic.Pertenece(a3))
	require.EqualValues(t, 0, dic.Obtener(a1))
	require.EqualValues(t, 1, dic.Obtener(a2))
	require.EqualValues(t, 2, dic.Obtener(a3))
	dic.Guardar(a1, 5)
	require.EqualValues(t, 5, dic.Obtener(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))
	require.EqualValues(t, 5, dic.Borrar(a1))
	require.False(t, dic.Pertenece(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))

}

func TestClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccOrdenado.CrearABB[string, string](strings.Compare)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccOrdenado.CrearABB[string, *int](strings.Compare)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func buscar(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestIteradorInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccOrdenado.CrearABB[string, *int](strings.Compare)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	dic.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, -1, buscar(cs[0], claves))
	require.NotEqualValues(t, -1, buscar(cs[1], claves))
	require.NotEqualValues(t, -1, buscar(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestIteradorInternoValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccOrdenado.CrearABB[string, int](strings.Compare)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestIteradorInternoValoresConBorrados(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccOrdenado.CrearABB[string, int](strings.Compare)
	dic.Guardar(clave0, 7)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	dic.Borrar(clave0)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func ejecutarPruebaVolumen(b *testing.B, n int) {
	dic := TDADiccOrdenado.CrearABB[int, int](func(a, b int) int { return a - b })

	claves := make([]int, n)
	valores := make([]int, n)

	// Inserta 'n' parejas en el ABB
	for i := 0; i < n; i++ {
		nroRandom := rand.Intn(RANGO_RANDOM)
		valores[i] = nroRandom
		claves[i] = nroRandom
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	// Verifica que devuelva los valores correctos
	ok := true
	for i := 0; i < n; i++ {
		ok = dic.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = dic.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	// Verifica que borre y devuelva los valores correctos
	for i := 0; i < n; i++ {
		ok = dic.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

func BenchmarkDiccionario(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumen(b, n)
			}
		})
	}
}

func TestIterarDiccionarioVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccOrdenado.CrearABB[string, int](strings.Compare)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDiccionarioIterar(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccOrdenado.CrearABB[string, string](strings.Compare)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(primero, claves))

	iter.Siguiente()
	segundo, segundo_valor := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.EqualValues(t, valores[buscar(segundo, claves)], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(tercero, claves))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccOrdenado.CrearABB[string, string](strings.Compare)
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()
	iter3 := dic.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1, buscar(primero, claves))
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.NotEqualValues(t, -1, buscar(tercero, claves))
}

func TestPruebaIterarTrasBorrados(t *testing.T) {
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"

	dic := TDADiccOrdenado.CrearABB[string, string](strings.Compare)
	dic.Guardar(clave1, "")
	dic.Guardar(clave2, "")
	dic.Guardar(clave3, "")
	dic.Borrar(clave1)
	dic.Borrar(clave2)
	dic.Borrar(clave3)
	iter := dic.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	dic.Guardar(clave1, "A")
	iter = dic.Iterador()

	require.True(t, iter.HaySiguiente())
	c1, v1 := iter.VerActual()
	require.EqualValues(t, clave1, c1)
	require.EqualValues(t, "A", v1)
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func ejecutarPruebasVolumenIterador(b *testing.B, n int) {
	dic := TDADiccOrdenado.CrearABB[string, *int](strings.Compare)

	claves := make([]string, n)
	valores := make([]int, n)

	// Inserta 'n' parejas en el ABB
	for i := 0; i < n; i++ {
		nroRandom := rand.Intn(RANGO_RANDOM)
		claves[i] = fmt.Sprintf("%08d", nroRandom)
		valores[i] = nroRandom
		dic.Guardar(claves[i], &valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	iter := dic.Iterador()
	require.True(b, iter.HaySiguiente())

	ok := true
	var i int
	var clave string
	var valor *int

	for i = 0; i < n; i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		clave = c1
		if clave == "" {
			ok = false
			break
		}
		valor = v1
		if valor == nil {
			ok = false
			break
		}
		*valor = n
		iter.Siguiente()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")

	ok = true
	for i = 0; i < n; i++ {
		if valores[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok, "No se cambiaron todos los elementos")
}

func BenchmarkIterador(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenIterador(b, n)
			}
		})
	}
}

func TestIteradorInternoConCorte(t *testing.T) {
	dic := TDADiccOrdenado.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(5, 1)
	dic.Guardar(4, 2)
	dic.Guardar(10, 3)

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false
	dic.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c == 5 {
			seguirEjecutando = false
			return false
		}
		return true
	})
	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia, "No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

func TestVolumenIteradorCorte(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDADiccOrdenado.CrearABB[int, int](func(a, b int) int { return a - b })
	n := 10000

	// Inserta 'n' parejas en el ABB
	for i := 0; i < n; i++ {
		dic.Guardar(rand.Intn(RANGO_RANDOM), i)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia, "No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

func TestIteradorIntRango(t *testing.T) {
	t.Log("Valida que el diccionario se recorre internamente en el orden dado por el desde/hasta")
	clave1 := 4
	clave2 := 2
	clave3 := 10
	clave4 := 7
	clave5 := 12

	dic := TDADiccOrdenado.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	sumatoria := 0
	ptrSumatoria := &sumatoria
	dic.IterarRango(&clave1, &clave3, func(_ int, dato int) bool {
		*ptrSumatoria += dato
		return true
	})

	require.EqualValues(t, 13, sumatoria)
}

func TestIteradorIntRangoConDesdeNil(t *testing.T) {
	t.Log("Valida que el diccionario se recorre internamente en el orden dado por el desde/hasta " +
		"comenzando desde el principio porque desde es nil")
	clave1 := 4
	clave2 := 2
	clave3 := 10
	clave4 := 7
	clave5 := 12

	dic := TDADiccOrdenado.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	sumatoria := 0
	ptrSumatoria := &sumatoria
	dic.IterarRango(nil, &clave3, func(_ int, dato int) bool {
		*ptrSumatoria += dato
		return true
	})

	require.EqualValues(t, 15, sumatoria)
}

func TestIteradorIntRangoConHastaNil(t *testing.T) {
	t.Log("Valida que el diccionario se recorre internamente en el orden dado por el desde/hasta " +
		"recorriendo hasta el final el diccionario porque hasta es nil")
	clave1 := 4
	clave2 := 2
	clave3 := 10
	clave4 := 7
	clave5 := 12

	dic := TDADiccOrdenado.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	sumatoria := 0
	ptrSumatoria := &sumatoria
	dic.IterarRango(&clave4, nil, func(_ int, dato int) bool {
		*ptrSumatoria += dato
		return true
	})

	require.EqualValues(t, 12, sumatoria)
}

func TestIteradorIntRangoConDesdeYHastaNil(t *testing.T) {
	t.Log("Valida que el diccionario se recorre internamente en el orden dado por el desde/hasta " +
		"recorriendo todo el diccionario sin restricciones porque desde y hasta son nil")
	clave1 := 4
	clave2 := 2
	clave3 := 10
	clave4 := 7
	clave5 := 12

	dic := TDADiccOrdenado.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	sumatoria := 0
	ptrSumatoria := &sumatoria
	dic.IterarRango(nil, nil, func(_ int, dato int) bool {
		*ptrSumatoria += dato
		return true
	})

	require.EqualValues(t, 20, sumatoria)
}

func TestIteradorIntRangoConCorte(t *testing.T) {
	t.Log("Valida que el diccionario se recorre internamente en el orden dado por el desde/hasta " +
		"y termina cuando la condicion de corte lo indica")
	clave1 := 4
	clave2 := 2
	clave3 := 10
	clave4 := 7
	clave5 := 12
	clave6 := 8

	dic := TDADiccOrdenado.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)
	dic.Guardar(clave6, 1)

	sumatoria := 0
	ptrSumatoria := &sumatoria
	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false
	dic.IterarRango(&clave1, &clave3, func(clave int, dato int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if clave%10 == 0 {
			seguirEjecutando = false
			return false
		}
		*ptrSumatoria += dato
		return true
	})
	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia, "No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
	require.EqualValues(t, 11, sumatoria)
}

func TestIteradorExtRango(t *testing.T) {
	t.Log("Guardamos 4 valores en un Diccionario, e iteramos en el orden dado por desde/hasta " +
		"validando que los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := 4
	clave2 := 2
	clave3 := 10
	clave4 := 5
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	valor4 := "beeh"
	dic := TDADiccOrdenado.CrearABB[int, string](func(a, b int) int { return a - b })
	dic.Guardar(clave1, valor1)
	dic.Guardar(clave2, valor2)
	dic.Guardar(clave3, valor3)
	dic.Guardar(clave4, valor4)
	iter := dic.IteradorRango(&clave2, &clave4)

	require.True(t, iter.HaySiguiente())
	_, primer_valor := iter.VerActual()
	require.EqualValues(t, valor2, primer_valor)

	iter.Siguiente()
	_, segundo_valor := iter.VerActual()
	require.EqualValues(t, valor1, segundo_valor)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	_, tercer_valor := iter.VerActual()
	require.EqualValues(t, valor4, tercer_valor)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorExtRangoConDesdeNil(t *testing.T) {
	t.Log("Guardamos 5 valores en un Diccionario, e iteramos desde el principio, porque desde es nil, hasta el valor dado a 'hasta' " +
		"validando que los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := 4
	clave2 := 2
	clave3 := 10
	clave4 := 5
	clave5 := 1
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	valor4 := "beeh"
	valor5 := "roar"
	dic := TDADiccOrdenado.CrearABB[int, string](func(a, b int) int { return a - b })
	dic.Guardar(clave1, valor1)
	dic.Guardar(clave2, valor2)
	dic.Guardar(clave3, valor3)
	dic.Guardar(clave4, valor4)
	dic.Guardar(clave5, valor5)
	iter := dic.IteradorRango(nil, &clave4)

	require.True(t, iter.HaySiguiente())
	_, primer_valor := iter.VerActual()
	require.EqualValues(t, valor5, primer_valor)

	iter.Siguiente()
	_, segundo_valor := iter.VerActual()
	require.EqualValues(t, valor2, segundo_valor)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	_, tercer_valor := iter.VerActual()
	require.EqualValues(t, valor1, tercer_valor)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	_, cuarto_valor := iter.VerActual()
	require.EqualValues(t, valor4, cuarto_valor)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorExtRangoConHastaNil(t *testing.T) {
	t.Log("Guardamos 5 valores en un Diccionario, e iteramos desde un valor dado hasta el final porque hasta es nil " +
		"validando que los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := 4
	clave2 := 2
	clave3 := 10
	clave4 := 5
	clave5 := 1
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	valor4 := "beeh"
	valor5 := "roar"
	dic := TDADiccOrdenado.CrearABB[int, string](func(a, b int) int { return a - b })
	dic.Guardar(clave1, valor1)
	dic.Guardar(clave2, valor2)
	dic.Guardar(clave3, valor3)
	dic.Guardar(clave4, valor4)
	dic.Guardar(clave5, valor5)
	iter := dic.IteradorRango(&clave1, nil)

	require.True(t, iter.HaySiguiente())
	_, primer_valor := iter.VerActual()
	require.EqualValues(t, valor1, primer_valor)

	iter.Siguiente()
	_, segundo_valor := iter.VerActual()
	require.EqualValues(t, valor4, segundo_valor)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	_, tercer_valor := iter.VerActual()
	require.EqualValues(t, valor3, tercer_valor)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorExtRangoConDesdeYHastaNil(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos todo el diccionario porque desde y hasta son nil " +
		"validando que los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := 4
	clave2 := 2
	clave3 := 10
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	dic := TDADiccOrdenado.CrearABB[int, string](func(a, b int) int { return a - b })
	dic.Guardar(clave1, valor1)
	dic.Guardar(clave2, valor2)
	dic.Guardar(clave3, valor3)

	iter := dic.IteradorRango(nil, nil)

	require.True(t, iter.HaySiguiente())
	_, primer_valor := iter.VerActual()
	require.EqualValues(t, valor2, primer_valor)

	iter.Siguiente()
	_, segundo_valor := iter.VerActual()
	require.EqualValues(t, valor1, segundo_valor)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	_, tercer_valor := iter.VerActual()
	require.EqualValues(t, valor3, tercer_valor)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}
