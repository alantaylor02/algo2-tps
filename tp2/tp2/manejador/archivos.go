package manejador

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"tp2/errores"
	"tp2/vuelos"
)

func LeerArchivoVuelos(ruta string, aeropuerto *vuelos.Aeropuerto) error {
	archivo, err := os.Open(ruta)

	if err != nil {
		return errores.ErrorEnComando{"agregar_archivo"}
	}
	defer archivo.Close()

	s := bufio.NewScanner(archivo)
	for s.Scan() {
		splitVueloLeido := strings.Split(s.Text(), ",")
		codigo, _ := strconv.Atoi(splitVueloLeido[0])
		prior, _ := strconv.Atoi(splitVueloLeido[5])
		fecha, origen, destino := splitVueloLeido[6], splitVueloLeido[2], splitVueloLeido[3]
		aerolinea, numeroCola, retrasoDespegue, tiempoVuelo, cancelado := splitVueloLeido[1], splitVueloLeido[4], splitVueloLeido[7], splitVueloLeido[8], splitVueloLeido[9]
		vueloLeido := vuelos.CrearVuelo(codigo, fecha, origen, destino, prior, aerolinea, numeroCola, retrasoDespegue, tiempoVuelo, cancelado)
		(*aeropuerto).AgregarOModificarVuelo(vueloLeido)
	}
	return nil
}
