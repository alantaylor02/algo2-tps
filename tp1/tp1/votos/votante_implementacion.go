package votos

import (
	"strconv"
	"tp1/errores"
)

type votanteImplementacion struct {
	dni               int
	voto              Voto
	estadosAnteriores []Voto
	yaVoto            bool
}

func CrearVotante(dni int) Votante {
	estadosAnteriores := make([]Voto, 0)
	estadosAnteriores = append(estadosAnteriores, Voto{})
	return &votanteImplementacion{dni, Voto{}, estadosAnteriores, false}
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, datoLeido string, cantidadPartidos int) error {
	if tipo == TIPO_VOTO_ERRONEO {
		return errores.ErrorTipoVoto{}
	}
	alternativa, errParseo := strconv.Atoi(datoLeido)

	if alternativa < LISTA_IMPUGNA || alternativa > cantidadPartidos || errParseo != nil {
		return errores.ErrorAlternativaInvalida{}
	}
	if votante.yaVoto {
		return errores.ErrorVotanteFraudulento{votante.LeerDNI()}
	}
	if alternativa == LISTA_IMPUGNA {
		votante.voto.Impugnado = true
	}
	votante.voto.VotoPorTipo[tipo] = alternativa
	votante.estadosAnteriores = append(votante.estadosAnteriores, votante.voto)
	return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	if votante.yaVoto {
		return errores.ErrorVotanteFraudulento{votante.LeerDNI()}
	}
	if len(votante.estadosAnteriores) == 1 {
		return errores.ErrorNoHayVotosAnteriores{}
	}
	votante.estadosAnteriores = votante.estadosAnteriores[:len(votante.estadosAnteriores)-1]
	ultimoEstado := votante.estadosAnteriores[len(votante.estadosAnteriores)-1]
	votante.voto = ultimoEstado
	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	if votante.yaVoto {
		return votante.voto, errores.ErrorVotanteFraudulento{votante.LeerDNI()}
	}
	votante.yaVoto = true
	return votante.voto, nil
}
