package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func Publica2() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese un número: ")
	if scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())
		if err != nil {
			numero, err = strconv.Atoi(scanner.Text())
		}
	}

	fmt.Println()
	fmt.Println("Tabla Númerica del ", numero)

	for i := 1; i <= 10; i++ {
		fmt.Println(numero, " x ", i, " = ", numero*i)
	}
}
