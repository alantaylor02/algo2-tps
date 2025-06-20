package vuelos

import (
	"fmt"
	"strconv"
	"strings"
	TDAColaPrior "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
	TDAPila "tdas/pila"
	"tp2/comparaciones"
	"tp2/errores"
)

const (
	CODIGO_AUX_DESDE = -1
	CODIGO_AUX_HASTA = 9
)

type aeropuertoImplementacion struct {
	vuelosPorCodigo   *TDADiccionario.Diccionario[int, *Vuelo]
	vuelosPorFecha    *TDADiccionario.DiccionarioOrdenado[Vuelo, string]
	vuelosPorConexion *TDADiccionario.Diccionario[string, *TDADiccionario.DiccionarioOrdenado[Vuelo, int]]
}

func CrearAeropuerto() Aeropuerto {
	vuelosPorCodigo := TDADiccionario.CrearHash[int, *Vuelo]()
	vuelosPorFecha := TDADiccionario.CrearABB[Vuelo, string](compararVuelosPorFechaYCodigo)
	vuelosPorConexion := TDADiccionario.CrearHash[string, *TDADiccionario.DiccionarioOrdenado[Vuelo, int]]()
	return &aeropuertoImplementacion{&vuelosPorCodigo, &vuelosPorFecha, &vuelosPorConexion}
}

func compararVuelosPorFechaYCodigo(vuelo1, vuelo2 Vuelo) int {
	cmpFecha := comparaciones.CompararFechas(vuelo1.VerFecha(), vuelo2.VerFecha())
	if cmpFecha == 0 {
		strCod1, strCod2 := strconv.Itoa(vuelo1.VerCodigo()), strconv.Itoa(vuelo2.VerCodigo())
		return strings.Compare(strCod1, strCod2)
	}
	return cmpFecha
}

func (a *aeropuertoImplementacion) AgregarOModificarVuelo(nuevoVuelo *Vuelo) {
	if (*a.vuelosPorCodigo).Pertenece((*nuevoVuelo).VerCodigo()) {
		vueloAnterior := (*a.vuelosPorCodigo).Obtener((*nuevoVuelo).VerCodigo())
		conexVueloAnt := (*vueloAnterior).VerConexion()
		dicFechasConexionAnt := (*a.vuelosPorConexion).Obtener(conexVueloAnt)
		(*a.vuelosPorFecha).Borrar(*vueloAnterior)
		(*dicFechasConexionAnt).Borrar(*vueloAnterior)
		(*a.vuelosPorCodigo).Guardar((*nuevoVuelo).VerCodigo(), nuevoVuelo)
		(*a.vuelosPorFecha).Guardar(*nuevoVuelo, fmt.Sprint((*nuevoVuelo).VerFecha(), " - ", (*nuevoVuelo).VerCodigo()))
	} else {
		(*a.vuelosPorCodigo).Guardar((*nuevoVuelo).VerCodigo(), nuevoVuelo)
		(*a.vuelosPorFecha).Guardar(*nuevoVuelo, fmt.Sprint((*nuevoVuelo).VerFecha(), " - ", (*nuevoVuelo).VerCodigo()))
	}
	if !(*a.vuelosPorConexion).Pertenece((*nuevoVuelo).VerConexion()) {
		dicFechasConexionNuevo := TDADiccionario.CrearABB[Vuelo, int](compararVuelosPorFechaYCodigo)
		dicFechasConexionNuevo.Guardar(*nuevoVuelo, 0)
		(*a.vuelosPorConexion).Guardar((*nuevoVuelo).VerConexion(), &dicFechasConexionNuevo)
	} else {
		dicFechasConexion := (*a.vuelosPorConexion).Obtener((*nuevoVuelo).VerConexion())
		(*dicFechasConexion).Guardar(*nuevoVuelo, 0)
	}
}

func (a aeropuertoImplementacion) obtenerKVuelosPorFecha(k int, modo, desde, hasta string, cmp func(string, string) int) TDADiccionario.DiccionarioOrdenado[string, int] {
	kVuelosOrdenados := TDADiccionario.CrearABB[string, int](cmp)
	vueloAuxDesdeIter, vueloAuxHastaIter := CrearVuelo(CODIGO_AUX_DESDE, desde, "", "", 0, "", "", "", "", ""), CrearVuelo(CODIGO_AUX_HASTA, hasta, "", "", 0, "", "", "", "", "")
	iterPorFecha := (*a.vuelosPorFecha).IteradorRango(vueloAuxDesdeIter, vueloAuxHastaIter)
	p := TDAPila.CrearPilaDinamica[string]()
	if modo == "asc" {
		for i := 0; i < k; i++ {
			if !iterPorFecha.HaySiguiente() {
				break
			}
			_, strFechaCodigo := iterPorFecha.VerActual()
			kVuelosOrdenados.Guardar(strFechaCodigo, 0)
			iterPorFecha.Siguiente()
		}
	} else {
		for iterPorFecha.HaySiguiente() {
			_, strFechaCodigo := iterPorFecha.VerActual()
			p.Apilar(strFechaCodigo)
			iterPorFecha.Siguiente()
		}
		for i := 0; i < k; i++ {
			if p.EstaVacia() {
				break
			}
			kVuelosOrdenados.Guardar(p.Desapilar(), 0)
		}
	}
	return kVuelosOrdenados
}

func (a aeropuertoImplementacion) VerTablero(k int, modo, desde, hasta string) TDADiccionario.DiccionarioOrdenado[string, int] {
	var tablero TDADiccionario.DiccionarioOrdenado[string, int]
	if modo == "asc" {
		tablero = a.obtenerKVuelosPorFecha(k, modo, desde, hasta, comparaciones.CompararFechasYCodigo)
	} else {
		cmpAnterior := comparaciones.CompararFechasYCodigo
		tablero = a.obtenerKVuelosPorFecha(k, modo, desde, hasta, func(a, b string) int { return cmpAnterior(b, a) })
	}
	return tablero
}

func (a aeropuertoImplementacion) InfoVuelo(codigo int) (string, error) {
	if !(*a.vuelosPorCodigo).Pertenece(codigo) {
		return "", errores.ErrorEnComando{"info_vuelo"}
	}
	vuelo := (*a.vuelosPorCodigo).Obtener(codigo)
	return vuelo.MostrarInformacion(), nil
}

func (a aeropuertoImplementacion) vuelosPorPrioridad(cmp func(a, b string) int) TDAColaPrior.ColaPrioridad[string] {
	arrVuelos := make([]string, 0)
	iterVuelos := (*a.vuelosPorCodigo).Iterador()
	for iterVuelos.HaySiguiente() {
		codigo, vueloActual := iterVuelos.VerActual()
		arrVuelos = append(arrVuelos, fmt.Sprint((*vueloActual).VerPrioridad(), " - ", codigo))
		iterVuelos.Siguiente()
	}
	colaPrioridad := TDAColaPrior.CrearHeapArr(arrVuelos, comparaciones.CompararPrioridadYCodigo)
	return colaPrioridad
}

func (a aeropuertoImplementacion) VerKVuelosPrioritarios(k int) ([]string, error) {
	if k <= 0 {
		return []string{}, errores.ErrorEnComando{"prioridad_vuelos"}
	}
	colaVuelosPrioritarios := a.vuelosPorPrioridad(comparaciones.CompararPrioridadYCodigo)
	kPrioritarios := make([]string, 0)
	for i := 0; i < k; i++ {
		if colaVuelosPrioritarios.EstaVacia() {
			break
		}
		kPrioritarios = append(kPrioritarios, colaVuelosPrioritarios.Desencolar())
	}
	return kPrioritarios, nil
}

func (a aeropuertoImplementacion) SiguienteVuelo(origen, destino, fecha string) string {
	conexion := fmt.Sprint(origen, "-", destino)
	if !(*a.vuelosPorConexion).Pertenece(conexion) {
		return fmt.Sprint("No hay vuelo registrado desde ", origen, " hacia ", destino, " desde ", fecha)
	}
	fechasParaConexion := (*a.vuelosPorConexion).Obtener(conexion)
	vueloAuxIter := CrearVuelo(CODIGO_AUX_DESDE, fecha, "", "", 0, "", "", "", "", "")
	iterFechas := (*fechasParaConexion).IteradorRango(vueloAuxIter, nil)
	if iterFechas.HaySiguiente() {
		vueloEncontrado, _ := iterFechas.VerActual()
		return vueloEncontrado.MostrarInformacion()
	}
	return fmt.Sprint("No hay vuelo registrado desde ", origen, " hacia ", destino, " desde ", fecha)
}

func (a *aeropuertoImplementacion) Borrar(desde, hasta string) []Vuelo {
	var vuelosBorrados []Vuelo
	vueloAuxDesdeIter, vueloAuxHastaIter := CrearVuelo(CODIGO_AUX_DESDE, desde, "", "", 0, "", "", "", "", ""), CrearVuelo(CODIGO_AUX_HASTA, hasta, "", "", 0, "", "", "", "", "")
	iterPorFecha := (*a.vuelosPorFecha).IteradorRango(vueloAuxDesdeIter, vueloAuxHastaIter)
	for iterPorFecha.HaySiguiente() {
		vueloBorrado, _ := iterPorFecha.VerActual()
		vuelosBorrados = append(vuelosBorrados, vueloBorrado)
		iterPorFecha.Siguiente()
	}
	for _, vueloBorrado := range vuelosBorrados {
		(*a.vuelosPorFecha).Borrar(vueloBorrado)
		(*a.vuelosPorCodigo).Borrar(vueloBorrado.VerCodigo())
		dicFechasConexion := (*a.vuelosPorConexion).Obtener(vueloBorrado.VerConexion())
		(*dicFechasConexion).Borrar(vueloBorrado)
	}
	return vuelosBorrados
}
