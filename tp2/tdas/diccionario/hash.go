package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

const (
	TAM_INICIAL        int    = 10
	FACTOR_AUMENTO     int    = 3
	FACTOR_REDUCCION   int    = 1
	FACTOR_REDIMENSION int    = 2
	uint64Offset       uint64 = 0xcbf29ce484222325
	uint64Prime        uint64 = 0x00000100000001b3
)

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []TDALista.Lista[*parClaveValor[K, V]]
	tam      int
	cantidad int
}

type iteradorHash[K comparable, V any] struct {
	iterListaActual TDALista.IteradorLista[*parClaveValor[K, V]]
	posListaNoVacia int
	cantParesVistos int
	parActual       *parClaveValor[K, V]
	hash            *hashAbierto[K, V]
}

// URL: https://bit.ly/3LXjl21
func fvnHash(data []byte) (hash uint64) {
	hash = uint64Offset
	for _, b := range data {
		hash ^= uint64(b)
		hash *= uint64Prime
	}
	return
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	tabla := make([]TDALista.Lista[*parClaveValor[K, V]], TAM_INICIAL)
	for i := range tabla {
		tabla[i] = TDALista.CrearListaEnlazada[*parClaveValor[K, V]]()
	}
	return &hashAbierto[K, V]{tabla, TAM_INICIAL, 0}
}

func (h *hashAbierto[K, V]) redimensionarHash(nuevoTam int) {
	nuevaTabla := make([]TDALista.Lista[*parClaveValor[K, V]], nuevoTam)
	for i := range nuevaTabla {
		nuevaTabla[i] = TDALista.CrearListaEnlazada[*parClaveValor[K, V]]()
	}
	for i := range h.tabla {
		if !h.tabla[i].EstaVacia() {
			iterListaActual := h.tabla[i].Iterador()
			for iterListaActual.HaySiguiente() {
				parActual := iterListaActual.VerActual()
				bytesClave := convertirABytes[K](parActual.clave)
				nuevaPos := fvnHash(bytesClave) % uint64(nuevoTam)
				nuevaTabla[nuevaPos].InsertarUltimo(&parClaveValor[K, V]{parActual.clave, parActual.dato})
				iterListaActual.Siguiente()
			}
		}
	}
	h.tam = nuevoTam
	h.tabla = nuevaTabla
}

func (h hashAbierto[K, V]) obtenerPosicionTabla(clave K) uint64 {
	bytesClave := convertirABytes[K](clave)
	return fvnHash(bytesClave) % uint64(h.tam)
}

func (h *hashAbierto[K, V]) buscarClave(clave K, pos uint64, borrar bool) *parClaveValor[K, V] {
	iter := h.tabla[pos].Iterador()
	for iter.HaySiguiente() {
		claveValorActual := iter.VerActual()
		if claveValorActual.clave == clave {
			if borrar {
				iter.Borrar()
			}
			return claveValorActual
		}
		iter.Siguiente()
	}
	return nil
}

func (h *hashAbierto[K, V]) Guardar(clave K, dato V) {
	pos := h.obtenerPosicionTabla(clave)
	claveValorEncontrado := h.buscarClave(clave, pos, false)
	if claveValorEncontrado == nil {
		h.tabla[pos].InsertarUltimo(&parClaveValor[K, V]{clave, dato})
		h.cantidad++
	} else {
		claveValorEncontrado.dato = dato
	}
	if h.Cantidad()/h.tam > FACTOR_AUMENTO {
		h.redimensionarHash(h.tam * FACTOR_REDIMENSION)
	}
}

func (h hashAbierto[K, V]) Pertenece(clave K) bool {
	pos := h.obtenerPosicionTabla(clave)
	claveValorEncontrado := h.buscarClave(clave, pos, false)
	if claveValorEncontrado == nil {
		return false
	}
	return true
}

func (h hashAbierto[K, V]) Obtener(clave K) V {
	pos := h.obtenerPosicionTabla(clave)
	claveValorEncontrado := h.buscarClave(clave, pos, false)
	if claveValorEncontrado == nil {
		panic("La clave no pertenece al diccionario")
	}
	return claveValorEncontrado.dato
}

func (h *hashAbierto[K, V]) Borrar(clave K) V {
	pos := h.obtenerPosicionTabla(clave)
	claveValorBorrado := h.buscarClave(clave, pos, true)
	if claveValorBorrado == nil {
		panic("La clave no pertenece al diccionario")
	}
	h.cantidad--

	if h.Cantidad()/h.tam < FACTOR_REDUCCION && h.tam > TAM_INICIAL {
		h.redimensionarHash(h.tam / FACTOR_REDIMENSION)
	}
	return claveValorBorrado.dato
}

func (h hashAbierto[K, V]) Cantidad() int {
	return h.cantidad
}

func (h *hashAbierto[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for i := range h.tabla {
		if !h.tabla[i].EstaVacia() {
			for iterL := h.tabla[i].Iterador(); iterL.HaySiguiente(); iterL.Siguiente() {
				claveValorActual := iterL.VerActual()
				if !visitar(claveValorActual.clave, claveValorActual.dato) {
					return
				}
			}
		}
	}
}

func avanzarPosListaNoVacia[K comparable, V any](posListaActual int, tabla []TDALista.Lista[*parClaveValor[K, V]]) int {
	for posListaActual < len(tabla) && tabla[posListaActual].EstaVacia() {
		posListaActual++
	}
	return posListaActual
}

func (h hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	var parActual *parClaveValor[K, V]
	var iterLista TDALista.IteradorLista[*parClaveValor[K, V]]
	var posListaNoVacia int
	var cantParesVistos int
	if h.cantidad == 0 {
		iterLista = nil
		parActual = nil
	} else {
		posListaNoVacia = avanzarPosListaNoVacia[K, V](0, h.tabla)
		iterLista = h.tabla[posListaNoVacia].Iterador()
		parActual = iterLista.VerActual()
		cantParesVistos++
	}
	return &iteradorHash[K, V]{iterLista, posListaNoVacia, cantParesVistos, parActual, &h}
}

func (iter *iteradorHash[K, V]) HaySiguiente() bool {
	return iter.parActual != nil
}

func (iter iteradorHash[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.parActual.clave, iter.parActual.dato
}

func (iter *iteradorHash[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	if iter.iterListaActual.HaySiguiente() {
		iter.iterListaActual.Siguiente()
		if !iter.iterListaActual.HaySiguiente() && iter.cantParesVistos < iter.hash.cantidad {
			sigPosTabla := iter.posListaNoVacia + 1
			sigListaNoVacia := avanzarPosListaNoVacia[K, V](sigPosTabla, iter.hash.tabla)
			iter.posListaNoVacia = sigListaNoVacia
			iter.iterListaActual = iter.hash.tabla[iter.posListaNoVacia].Iterador()
			iter.parActual = iter.iterListaActual.VerActual()
			iter.cantParesVistos++
		} else if iter.iterListaActual.HaySiguiente() {
			iter.parActual = iter.iterListaActual.VerActual()
			iter.cantParesVistos++
		} else {
			iter.parActual = nil
		}
	}
}
