package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Publica2() string {
	var numero int
	var err error
	var texto string

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese un número: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	texto += fmt.Sprintln("Tabla Númerica del ", numero, "\n")

	for i := 1; i <= 10; i++ {
		// fmt.Println(numero, "x", i, "=", numero*i)
		texto += fmt.Sprintln(numero, "x", i, "=", numero*i)
	}
	texto += fmt.Sprintln("\n")

	return texto
}
