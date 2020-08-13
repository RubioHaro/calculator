package calculator

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// LeerEntrada Funcion para leer inputs
func LeerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

//Calculadora ...
type Calculadora struct {
	entrada   string
	operacion string
}

//Operar ...
func (Calculadora) Operar(entrada string, operador string) int {

	entradaLimpia := strings.Split(entrada, operador)
	operador1 := convertir(entradaLimpia[0])
	operador2 := convertir(entradaLimpia[1])

	switch operador {
	case "+":
		fmt.Println("Suma de los dos operadores matematicamente: ", operador1+operador2)
		return operador1 + operador2
	case "-":
		fmt.Println("Suma de los dos operadores matematicamente: ", operador1-operador2)
		return operador1 - operador2
	case "*":
		fmt.Println("Suma de los dos operadores matematicamente: ", operador1*operador2)
		return operador1 * operador2
	case "/":
		fmt.Println("Suma de los dos operadores matematicamente: ", operador1/operador2)
		return operador1 / operador2
	default:
		fmt.Println("Sin operador valido ", operador)
		return 0
	}

}

func convertir(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

var separadores = [4]string{"+", "-", "*", "/"}

// DetectarOperador ...
func DetectarOperador(operacion string) (string, error) {
	var separador = ""
	for i := 0; i < len(separadores); i++ {
		indx := strings.Index(operacion, separadores[i])
		if indx != -1 && separador == "" {
			separador = separadores[i]
		} else if indx != -1 && separador != "" {
			return "", errors.New("La operacion tiene mas de un operador")
		}
	}

	if separador == "" {
		return "", errors.New("No se ha encontrado operador valido")
	}

	return separador, nil
}
