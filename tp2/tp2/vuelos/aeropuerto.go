package vuelos

import TDADiccionario "tdas/diccionario"

type Aeropuerto interface {
	AgregarOModificarVuelo(nuevoVuelo *Vuelo)

	VerTablero(k int, modo, desde, hasta string) TDADiccionario.DiccionarioOrdenado[string, int]

	InfoVuelo(codigo int) (string, error)

	VerKVuelosPrioritarios(k int) ([]string, error)

	SiguienteVuelo(origen, destino, fecha string) string

	Borrar(desde, hasta string) []Vuelo
}
