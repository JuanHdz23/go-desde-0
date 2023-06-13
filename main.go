package main

import (
	"fmt"

	"github.com/juanhdz23/go-desde-0/variables"
)

func main() {
	// fmt.Println("Hola Mundo JC")
	// variables.MuestroEnteros()
	// variables.RestoVariables()
	estado, texto := variables.ConviertoTexto(15)
	fmt.Println(estado)
	fmt.Println(texto)
}
