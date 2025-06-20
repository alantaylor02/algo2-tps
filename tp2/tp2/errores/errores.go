package errores

import "fmt"

type ErrorEnComando struct {
	Comando string
}

func (e ErrorEnComando) Error() string {
	return fmt.Sprintf("Error en comando %s", e.Comando)
}
