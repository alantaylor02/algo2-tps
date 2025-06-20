package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tp2/manejador"
	"tp2/vuelos"
)

func main() {
	aeropuerto := vuelos.CrearAeropuerto()
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		inputUsuario := strings.Split(s.Text(), " ")
		errObtenido := manejador.RealizarAccionSegunComando(inputUsuario, &aeropuerto)
		if errObtenido != nil {
			fmt.Fprintf(os.Stderr, "%s\n", errObtenido.Error())
		} else {
			fmt.Println("OK")
		}
	}
}
