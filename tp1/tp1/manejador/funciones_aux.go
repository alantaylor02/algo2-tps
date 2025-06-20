package manejador

import (
	"fmt"
	"strconv"
	TDACola "tdas/cola"
	"tp1/errores"
	"tp1/votos"
)

const PARTIDO_EN_BLANCO = 1

func countingSort(arrVotantes []votos.Votante, criterio func(int) int) []votos.Votante {
	frecs := make([]int, 10)
	for _, votante := range arrVotantes {
		frecs[criterio(votante.LeerDNI())]++
	}
	acum := make([]int, 10)
	for i := 1; i < len(acum); i++ {
		acum[i] = acum[i-1] + frecs[i-1]
	}
	res := make([]votos.Votante, len(arrVotantes))
	for _, votante := range arrVotantes {
		res[acum[criterio(votante.LeerDNI())]] = votante
		acum[criterio(votante.LeerDNI())]++
	}
	return res
}

func RadixSort(arrVotantes []votos.Votante) []votos.Votante {
	arrVotantes = countingSort(arrVotantes, func(n int) int { return n % 10 })
	arrVotantes = countingSort(arrVotantes, func(n int) int { return n / 10 % 10 })
	arrVotantes = countingSort(arrVotantes, func(n int) int { return n / 100 % 10 })
	arrVotantes = countingSort(arrVotantes, func(n int) int { return n / 1000 % 10 })
	arrVotantes = countingSort(arrVotantes, func(n int) int { return n / 10000 % 10 })
	arrVotantes = countingSort(arrVotantes, func(n int) int { return n / 100000 % 10 })
	arrVotantes = countingSort(arrVotantes, func(n int) int { return n / 1000000 % 10 })
	arrVotantes = countingSort(arrVotantes, func(n int) int { return n / 10000000 % 10 })
	return arrVotantes
}

func busquedaBinaria(arr []votos.Votante, elem int) int {
	izq := 0
	der := len(arr) - 1
	for izq <= der {
		medio := (izq + der) / 2
		if arr[medio].LeerDNI() == elem {
			return medio
		} else if arr[medio].LeerDNI() < elem {
			izq = medio + 1
		} else {
			der = medio - 1
		}
	}
	return -1
}

func ValidarDNI(votantes []votos.Votante, dni string) (error, *votos.Votante) {
	dniInt, errConversion := strconv.Atoi(dni)
	if dniInt <= 0 || errConversion != nil {
		return errores.DNIError{}, nil
	}
	res := busquedaBinaria(votantes, dniInt)
	if res == -1 {
		return errores.DNIFueraPadron{}, nil
	}
	return nil, &(votantes[res])
}

func ValidarFilaVacia(fila TDACola.Cola[*votos.Votante]) error {
	if fila.EstaVacia() {
		return errores.FilaVacia{}
	}
	return nil
}

func ComprobarTipoVoto(cargo string) votos.TipoVoto {
	switch cargo {
	case "Presidente":
		return votos.PRESIDENTE
	case "Gobernador":
		return votos.GOBERNADOR
	case "Intendente":
		return votos.INTENDENTE
	}
	return votos.TIPO_VOTO_ERRONEO
}

func CargoAImprimir(i votos.TipoVoto) string {
	switch i {
	case 0:
		return "Presidente"
	case 1:
		return "Gobernador"
	}
	return "Intendente"
}

func ImprimirResultados(partidos []votos.Partido, cantVotosImpugnados votos.TipoVoto) {
	var i votos.TipoVoto = 0
	for i = 0; i < 3; i++ {
		fmt.Printf("%s:\n", CargoAImprimir(i))
		for j := range partidos {
			fmt.Println(partidos[j].ObtenerResultado(i))
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Votos Impugnados: %d %s\n", cantVotosImpugnados, votos.VotoEnSingularOPlural(cantVotosImpugnados))
}

func esErrorVotanteFraudulento(err error, votante votos.Votante) bool {
	return err == errores.ErrorVotanteFraudulento{votante.LeerDNI()}
}

func EjecutarIngreso(inputUsuario []string, votantes *[]votos.Votante, fila *TDACola.Cola[*votos.Votante]) error {
	errIngresar, ptrVotante := ValidarDNI(*votantes, inputUsuario[1])
	if errIngresar != nil {
		return errIngresar
	}
	(*fila).Encolar(ptrVotante)
	return nil
}

func EjecutarVoto(inputUsuario []string, votantes *[]votos.Votante, partidos []votos.Partido, fila *TDACola.Cola[*votos.Votante]) error {
	errFilaVacia := ValidarFilaVacia(*fila)
	if errFilaVacia != nil {
		return errFilaVacia
	}
	votanteActual := (*fila).VerPrimero()
	errVotar := (*votanteActual).Votar(ComprobarTipoVoto(inputUsuario[1]), inputUsuario[2], len(partidos)-PARTIDO_EN_BLANCO)
	if errVotar != nil {
		if esErrorVotanteFraudulento(errVotar, *votanteActual) {
			(*fila).Desencolar()
		}
		return errVotar
	}
	return nil
}

func EjecutarDeshacer(votantes *[]votos.Votante, fila *TDACola.Cola[*votos.Votante]) error {
	errFilaVacia := ValidarFilaVacia(*fila)
	if errFilaVacia != nil {
		return errFilaVacia
	}
	votanteActual := (*fila).VerPrimero()
	errDeshacer := (*votanteActual).Deshacer()
	if errDeshacer != nil {
		if esErrorVotanteFraudulento(errDeshacer, *votanteActual) {
			(*fila).Desencolar()
		}
		return errDeshacer
	}
	return nil
}

func EjecutarFinVoto(votantes *[]votos.Votante, partidos []votos.Partido, fila *TDACola.Cola[*votos.Votante], cantVotosImpugnados *votos.TipoVoto) error {
	errFilaVacia := ValidarFilaVacia(*fila)
	if errFilaVacia != nil {
		return errFilaVacia
	}
	votanteFinalizado := (*fila).Desencolar()
	voto, errFinalizacion := (*votanteFinalizado).FinVoto()
	if errFinalizacion != nil {
		return errFinalizacion
	}
	if !voto.Impugnado {
		partidos[voto.VotoPorTipo[votos.PRESIDENTE]].VotadoPara(votos.PRESIDENTE)
		partidos[voto.VotoPorTipo[votos.GOBERNADOR]].VotadoPara(votos.GOBERNADOR)
		partidos[voto.VotoPorTipo[votos.INTENDENTE]].VotadoPara(votos.INTENDENTE)
	} else {
		*cantVotosImpugnados++
	}
	return nil
}

func RealizarAccionSegunComando(inputUsuario []string, votantes *[]votos.Votante, partidos []votos.Partido, fila *TDACola.Cola[*votos.Votante], cantVotosImpugnados *votos.TipoVoto) error {
	var errEjecucion error
	comando := inputUsuario[0]
	switch comando {
	case "ingresar":
		errEjecucion = EjecutarIngreso(inputUsuario, votantes, fila)
	case "votar":
		errEjecucion = EjecutarVoto(inputUsuario, votantes, partidos, fila)
	case "deshacer":
		errEjecucion = EjecutarDeshacer(votantes, fila)
	case "fin-votar":
		errEjecucion = EjecutarFinVoto(votantes, partidos, fila, cantVotosImpugnados)
	}
	return errEjecucion
}
