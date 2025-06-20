package comparaciones

import (
	"strconv"
	"strings"
)

func ConvertirFechaAEntero(fecha string) (int, int, int, int, int, int) {
	arrFecha := strings.Split(fecha, "T")
	añoMesDia, horaMinSeg := arrFecha[0], arrFecha[1]
	arrAñoMesDia := strings.Split(añoMesDia, "-")
	arrHoraMinSeg := strings.Split(horaMinSeg, ":")
	año, mes, dia := arrAñoMesDia[0], arrAñoMesDia[1], arrAñoMesDia[2]
	hora, min, seg := arrHoraMinSeg[0], arrHoraMinSeg[1], arrHoraMinSeg[2]
	añoInt, _ := strconv.Atoi(año)
	mesInt, _ := strconv.Atoi(mes)
	diaInt, _ := strconv.Atoi(dia)
	horaInt, _ := strconv.Atoi(hora)
	minInt, _ := strconv.Atoi(min)
	segInt, _ := strconv.Atoi(seg)
	return añoInt, mesInt, diaInt, horaInt, minInt, segInt
}

func CompararFechas(fecha1, fecha2 string) int {
	año1, mes1, dia1, hora1, min1, seg1 := ConvertirFechaAEntero(fecha1)
	año2, mes2, dia2, hora2, min2, seg2 := ConvertirFechaAEntero(fecha2)
	if año1 > año2 {
		return 1
	} else if año1 < año2 {
		return -1
	}
	if mes1 > mes2 {
		return 1
	} else if mes1 < mes2 {
		return -1
	}
	if dia1 > dia2 {
		return 1
	} else if dia1 < dia2 {
		return -1
	}
	if hora1 > hora2 {
		return 1
	} else if hora1 < hora2 {
		return -1
	}
	if min1 > min2 {
		return 1
	} else if min1 < min2 {
		return -1
	}
	if seg1 > seg2 {
		return 1
	} else if seg1 < seg2 {
		return -1
	}
	return 0
}

func CompararPrioridad(prior1, prior2 string) int {
	prior1Int, _ := strconv.Atoi(prior1)
	prior2Int, _ := strconv.Atoi(prior2)
	return prior1Int - prior2Int
}

func CompararFechasYCodigo(fechaCodigo1, fechaCodigo2 string) int {
	arrFechaCod1 := strings.Split(fechaCodigo1, " - ")
	arrFechaCod2 := strings.Split(fechaCodigo2, " - ")
	fecha1, codigo1 := arrFechaCod1[0], arrFechaCod1[1]
	fecha2, codigo2 := arrFechaCod2[0], arrFechaCod2[1]
	cmpFecha := CompararFechas(fecha1, fecha2)
	if cmpFecha == 0 {
		return strings.Compare(codigo1, codigo2)
	}
	return cmpFecha
}

func CompararPrioridadYCodigo(priorCodigo1, priorCodigo2 string) int {
	arrPriorCod1 := strings.Split(priorCodigo1, " - ")
	arrPriorCod2 := strings.Split(priorCodigo2, " - ")
	prior1, codigo1 := arrPriorCod1[0], arrPriorCod1[1]
	prior2, codigo2 := arrPriorCod2[0], arrPriorCod2[1]
	cmpPrior := CompararPrioridad(prior1, prior2)
	if cmpPrior == 0 {
		strCompareInv := func(a, b string) int { return strings.Compare(b, a) }
		return strCompareInv(codigo1, codigo2)
	}
	return cmpPrior
}
