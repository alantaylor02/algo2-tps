package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	TDACola "tdas/cola"
	"tp1/errores"
	"tp1/manejador"
	"tp1/votos"
)

func main() {
	var cantVotosImpugnados votos.TipoVoto
	rutas, errParametros := manejador.LeerParametrosIniciales()
	if errParametros != nil {
		fmt.Println(errParametros.Error())
		return
	}
	partidos, errPartidos := manejador.LeerArchivoListas(rutas[0])
	votantes, errPadrones := manejador.LeerArchivoPadrones(rutas[1])
	if errPadrones != nil || errPartidos != nil {
		fmt.Println(errores.ErrorLeerArchivo{}.Error())
		return
	}
	fila := TDACola.CrearColaEnlazada[*votos.Votante]()

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		inputUsuario := strings.Split(s.Text(), " ")

		errObtenido := manejador.RealizarAccionSegunComando(inputUsuario, &votantes, partidos, &fila, &cantVotosImpugnados)
		if errObtenido != nil {
			fmt.Println(errObtenido.Error())
		} else {
			fmt.Println("OK")
		}
	}
	if !fila.EstaVacia() {
		fmt.Println(errores.ErrorCiudadanosSinVotar{}.Error())
	}
	manejador.ImprimirResultados(partidos, cantVotosImpugnados)
}
