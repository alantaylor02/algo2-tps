package votos

import "fmt"

type partidoImplementacion struct {
	nombre     string
	candidatos [CANT_VOTACION]string
	cant_votos [CANT_VOTACION]TipoVoto
}

type partidoEnBlanco struct {
	cant_votos [CANT_VOTACION]TipoVoto
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string, cant_votos [CANT_VOTACION]TipoVoto) Partido {
	return &partidoImplementacion{nombre, candidatos, cant_votos}
}

func CrearVotosEnBlanco() Partido {
	return &partidoEnBlanco{[CANT_VOTACION]TipoVoto{}}
}

func VotoEnSingularOPlural(voto TipoVoto) string {
	var palabra string
	if voto == 1 {
		palabra = "voto"
	} else {
		palabra = "votos"
	}
	return palabra
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.cant_votos[tipo]++
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	return fmt.Sprintf("%s - %s: %d %s", partido.nombre, partido.candidatos[tipo], partido.cant_votos[tipo], VotoEnSingularOPlural(partido.cant_votos[tipo]))
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.cant_votos[tipo]++
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	return fmt.Sprintf("Votos en Blanco: %d %s", blanco.cant_votos[tipo], VotoEnSingularOPlural(blanco.cant_votos[tipo]))
}
