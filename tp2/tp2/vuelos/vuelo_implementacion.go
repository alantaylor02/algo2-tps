package vuelos

import "fmt"

type Vuelo struct {
	codigo          int
	fecha           string
	origen          string
	destino         string
	prioridad       int
	aerolinea       string
	numeroCola      string
	retrasoDespegue string
	tiempoVuelo     string
	cancelado       string
}

func CrearVuelo(codigo int, fecha, origen, destino string, prioridad int, aerolinea, numeroCola, retrasoDespegue, tiempoVuelo, cancelado string) *Vuelo {
	return &Vuelo{codigo, fecha, origen, destino, prioridad, aerolinea, numeroCola, retrasoDespegue, tiempoVuelo, cancelado}
}

func (v Vuelo) VerCodigo() int {
	return v.codigo
}

func (v Vuelo) VerConexion() string {
	return fmt.Sprint(v.origen, "-", v.destino)
}

func (v Vuelo) VerFecha() string {
	return v.fecha
}

func (v Vuelo) VerPrioridad() int {
	return v.prioridad
}

func ignorarCerosIzquierda(cadena string) string {
	var res string
	for i := 0; i < len(cadena); i++ {
		if string(cadena[i]) == "-" {
			res += "-"
		} else if string(cadena[i]) != "0" {
			res += cadena[i:]
			break
		}
	}
	if res == "" {
		return "0"
	}
	return res
}

func (v Vuelo) MostrarInformacion() string {
	return fmt.Sprint(v.codigo, " ", v.aerolinea, " ", v.origen, " ", v.destino, " ", v.numeroCola, " ", v.prioridad, " ", v.fecha, " ", ignorarCerosIzquierda(v.retrasoDespegue), " ", v.tiempoVuelo, " ", v.cancelado)
}
