package manejador

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"tp1/errores"
	"tp1/votos"
)

func LeerParametrosIniciales() ([]string, error) {
	rutas := os.Args
	if len(rutas) < 3 {
		return nil, errores.ErrorParametros{}
	}
	return rutas[1:], nil
}

func LeerArchivoListas(ruta string) ([]votos.Partido, error) {
	archivo, err := os.Open(ruta)

	if err != nil {
		return nil, errores.ErrorLeerArchivo{}
	}
	defer archivo.Close()

	var partidosLeidos []votos.Partido
	partidoEnBlanco := votos.CrearVotosEnBlanco()
	partidosLeidos = append(partidosLeidos, partidoEnBlanco)

	s := bufio.NewScanner(archivo)
	for s.Scan() {
		splitListaLeida := strings.Split(s.Text(), ",")
		cantVotos := [votos.CANT_VOTACION]votos.TipoVoto{}
		candidatos := [votos.CANT_VOTACION]string{}
		candidatos[0], candidatos[1], candidatos[2] = splitListaLeida[1], splitListaLeida[2], splitListaLeida[3]
		partido := votos.CrearPartido(splitListaLeida[0], candidatos, cantVotos)
		partidosLeidos = append(partidosLeidos, partido)
	}
	return partidosLeidos, nil
}

func LeerArchivoPadrones(ruta string) ([]votos.Votante, error) {
	archivo, err := os.Open(ruta)

	if err != nil {
		return nil, errores.ErrorLeerArchivo{}
	}
	defer archivo.Close()

	padronesLeidos := make([]votos.Votante, 0)

	s := bufio.NewScanner(archivo)
	for s.Scan() {
		dniLeido := s.Text()
		if len(dniLeido) == 7 {
			dniLeido = "0" + dniLeido
		}
		padron, _ := strconv.Atoi(dniLeido)
		padronesLeidos = append(padronesLeidos, votos.CrearVotante(padron))
	}
	return RadixSort(padronesLeidos), nil
}
