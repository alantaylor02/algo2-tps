package diccionario

import (
	TDAPila "tdas/pila"
)

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

type iteradorABB[K comparable, V any] struct {
	pilaAux *TDAPila.Pila[*nodoAbb[K, V]]
	abb     *abb[K, V]
	desde   *K
	hasta   *K
}

func crearNodo[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{nil, nil, clave, dato}
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{nil, 0, funcion_cmp}
}

func (abb *abb[K, V]) buscarNodo(nodo, padre *nodoAbb[K, V], clave K) (*nodoAbb[K, V], *nodoAbb[K, V], bool) {
	if nodo == nil {
		return nodo, padre, false
	}
	if abb.cmp(nodo.clave, clave) == 0 {
		return nodo, padre, true
	} else if abb.cmp(nodo.clave, clave) < 0 {
		return abb.buscarNodo(nodo.derecho, nodo, clave)
	}
	return abb.buscarNodo(nodo.izquierdo, nodo, clave)
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	nodo, padre, fueEncontrado := abb.buscarNodo(abb.raiz, nil, clave)
	if fueEncontrado {
		nodo.dato = dato
	} else {
		nuevoNodo := crearNodo[K, V](clave, dato)
		if padre == nil {
			abb.raiz = nuevoNodo
		} else if abb.cmp(padre.clave, clave) < 0 {
			padre.derecho = nuevoNodo
		} else {
			padre.izquierdo = nuevoNodo
		}
		abb.cantidad++
	}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	_, _, fueEncontrado := abb.buscarNodo(abb.raiz, nil, clave)
	return fueEncontrado
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodoBuscado, _, fueEncontrado := abb.buscarNodo(abb.raiz, nil, clave)
	if !fueEncontrado {
		panic("La clave no pertenece al diccionario")
	}
	return nodoBuscado.dato
}

func (abb *abb[K, V]) buscarMinimo(nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo.izquierdo == nil {
		return nodo
	}
	return abb.buscarMinimo(nodo.izquierdo)
}

func (abb *abb[K, V]) Borrar(clave K) V {
	nodoABorrar, padre, fueEncontrado := abb.buscarNodo(abb.raiz, nil, clave)
	if !fueEncontrado {
		panic("La clave no pertenece al diccionario")
	}
	datoBorrado := nodoABorrar.dato
	var reemplazo *nodoAbb[K, V]
	if nodoABorrar.derecho != nil && nodoABorrar.izquierdo != nil {
		reemplazo = abb.buscarMinimo(nodoABorrar.derecho)
		claveReemplazo := reemplazo.clave
		datoReemplazo := abb.Borrar(reemplazo.clave)
		nodoABorrar.clave, nodoABorrar.dato = claveReemplazo, datoReemplazo
	} else {
		if padre == nil {
			if nodoABorrar.derecho != nil {
				abb.raiz = nodoABorrar.derecho
			} else {
				abb.raiz = nodoABorrar.izquierdo
			}
		} else {
			if nodoABorrar.derecho != nil {
				reemplazo = nodoABorrar.derecho
			} else {
				reemplazo = nodoABorrar.izquierdo
			}
			if abb.cmp(padre.clave, nodoABorrar.clave) < 0 {
				padre.derecho = reemplazo
			} else {
				padre.izquierdo = reemplazo
			}
		}
		abb.cantidad--
	}
	return datoBorrado
}

func (abb abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) iterarRango(nodo *nodoAbb[K, V], desde *K, hasta *K, visitar func(clave K, dato V) bool, cmp func(K, K) int) bool {
	if nodo == nil {
		return true
	}
	if !abb.iterarRango(nodo.izquierdo, desde, hasta, visitar, cmp) {
		return false
	}
	if desde == nil && hasta == nil {
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
	} else if desde == nil {
		if cmp(*hasta, nodo.clave) >= 0 && !visitar(nodo.clave, nodo.dato) {
			return false
		}
	} else if hasta == nil {
		if cmp(*desde, nodo.clave) <= 0 && !visitar(nodo.clave, nodo.dato) {
			return false
		}
	} else {
		if cmp(*desde, nodo.clave) <= 0 && cmp(*hasta, nodo.clave) >= 0 && !visitar(nodo.clave, nodo.dato) {
			return false
		}
	}
	return abb.iterarRango(nodo.derecho, desde, hasta, visitar, cmp)
}

func (abb abb[K, V]) Iterar(visitar func(K, V) bool) {
	abb.iterarRango(abb.raiz, nil, nil, visitar, abb.cmp)
}

func (abb abb[K, V]) IterarRango(desde, hasta *K, visitar func(clave K, dato V) bool) {
	abb.iterarRango(abb.raiz, desde, hasta, visitar, abb.cmp)
}

func apilarHijosIzquierdos[K comparable, V any](nodo *nodoAbb[K, V], pila *TDAPila.Pila[*nodoAbb[K, V]], desde *K, cmp func(K, K) int) {
	if nodo == nil {
		return
	}
	if desde != nil && cmp(*desde, nodo.clave) > 0 {
		apilarHijosIzquierdos(nodo.derecho, pila, desde, cmp)
	} else {
		(*pila).Apilar(nodo)
		apilarHijosIzquierdos(nodo.izquierdo, pila, desde, cmp)
	}
}

func (abb abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (abb abb[K, V]) IteradorRango(desde, hasta *K) IterDiccionario[K, V] {
	pilaAux := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	apilarHijosIzquierdos(abb.raiz, &pilaAux, desde, abb.cmp)
	return &iteradorABB[K, V]{&pilaAux, &abb, desde, hasta}
}

func (iter iteradorABB[K, V]) HaySiguiente() bool {
	if (*iter.pilaAux).EstaVacia() {
		return false
	}
	if iter.hasta != nil && iter.abb.cmp((*iter.pilaAux).VerTope().clave, (*iter.hasta)) > 0 {
		return false
	}
	return true
}

func (iter iteradorABB[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return (*iter.pilaAux).VerTope().clave, (*iter.pilaAux).VerTope().dato
}

func (iter *iteradorABB[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodoDesapilado := (*iter.pilaAux).Desapilar()
	if nodoDesapilado.derecho != nil {
		(*iter.pilaAux).Apilar(nodoDesapilado.derecho)
		apilarHijosIzquierdos[K, V](nodoDesapilado.derecho.izquierdo, iter.pilaAux, iter.desde, iter.abb.cmp)
	}
}
