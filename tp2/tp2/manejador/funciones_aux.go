package manejador

import (
	"fmt"
	"strconv"
	"tp2/errores"
	"tp2/vuelos"
)

func EjecutarAgregarArchivo(inputUsuario []string, aeropuerto *vuelos.Aeropuerto) error {
	if len(inputUsuario) != 2 {
		return errores.ErrorEnComando{"agregar_archivo"}
	}
	errAgregar := LeerArchivoVuelos(inputUsuario[1], aeropuerto)
	if errAgregar != nil {
		return errAgregar
	}
	return nil
}

func EjecutarVerTablero(inputUsuario []string, aeropuerto *vuelos.Aeropuerto) error {
	if len(inputUsuario) != 5 {
		return errores.ErrorEnComando{"ver_tablero"}
	}
	k, _ := strconv.Atoi(inputUsuario[1])
	dicKVuelosOrdenados := (*aeropuerto).VerTablero(k, inputUsuario[2], inputUsuario[3], inputUsuario[4])
	for iterDic := dicKVuelosOrdenados.Iterador(); iterDic.HaySiguiente(); iterDic.Siguiente() {
		strFechaCodigo, _ := iterDic.VerActual()
		fmt.Println(strFechaCodigo)
	}
	return nil
}

func EjecutarInfoVuelo(inputUsuario []string, aeropuerto *vuelos.Aeropuerto) error {
	if len(inputUsuario) != 2 {
		return errores.ErrorEnComando{"info_vuelo"}
	}
	codigo, _ := strconv.Atoi(inputUsuario[1])
	infoVuelo, errInfoVuelo := (*aeropuerto).InfoVuelo(codigo)
	if errInfoVuelo != nil {
		return errInfoVuelo
	}
	fmt.Println(infoVuelo)
	return nil
}

func EjecutarPriorVuelos(inputUsuario []string, aeropuerto *vuelos.Aeropuerto) error {
	if len(inputUsuario) != 2 {
		return errores.ErrorEnComando{"prioridad_vuelos"}
	}
	k, _ := strconv.Atoi(inputUsuario[1])
	kPrioritarios, errPrior := (*aeropuerto).VerKVuelosPrioritarios(k)
	if errPrior != nil {
		return errPrior
	}
	for i := 0; i < len(kPrioritarios); i++ {
		fmt.Println(kPrioritarios[i])
	}
	return nil
}

func EjecutarSigVuelo(inputUsuario []string, aeropuerto *vuelos.Aeropuerto) {
	infoSigVuelo := (*aeropuerto).SiguienteVuelo(inputUsuario[1], inputUsuario[2], inputUsuario[3])
	fmt.Println(infoSigVuelo)
}

func EjecutarBorrar(inputUsuario []string, aeropuerto *vuelos.Aeropuerto) error {
	if len(inputUsuario) != 3 {
		return errores.ErrorEnComando{"borrar"}
	}
	vuelosBorrados := (*aeropuerto).Borrar(inputUsuario[1], inputUsuario[2])
	for _, vueloBorrado := range vuelosBorrados {
		fmt.Println(vueloBorrado.MostrarInformacion())
	}
	return nil
}

func RealizarAccionSegunComando(inputUsuario []string, aeropuerto *vuelos.Aeropuerto) error {
	var errEjecucion error
	comando := inputUsuario[0]
	switch comando {
	case "agregar_archivo":
		errEjecucion = EjecutarAgregarArchivo(inputUsuario, aeropuerto)
	case "ver_tablero":
		errEjecucion = EjecutarVerTablero(inputUsuario, aeropuerto)
	case "info_vuelo":
		errEjecucion = EjecutarInfoVuelo(inputUsuario, aeropuerto)
	case "prioridad_vuelos":
		errEjecucion = EjecutarPriorVuelos(inputUsuario, aeropuerto)
	case "siguiente_vuelo":
		EjecutarSigVuelo(inputUsuario, aeropuerto)
	case "borrar":
		errEjecucion = EjecutarBorrar(inputUsuario, aeropuerto)
	}
	return errEjecucion
}
