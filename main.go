package main

import (
	"fmt"

	"github.com/juanhdz23/go-desde-0/ejercicios"
	// "runtime"
	// "github.com/juanhdz23/go-desde-0/variables"
)

func main() {
	// fmt.Println("Hola Mundo JC")
	// variables.MuestroEnteros()
	// variables.RestoVariables()
	// estado, texto := variables.ConviertoTexto(15)
	// fmt.Println(estado)
	// fmt.Println(texto)

	// os := runtime.GOOS
	// if os=="Linux." {
	// fmt.Println(runtime.GOOS)
	// if os := runtime.GOOS; os == "linux" || os == "OS X." {
	// 	fmt.Println("Esto no es Windows, es", os)
	// } else {
	// 	fmt.Println("Esto es Windows")
	// }

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Esto es Linux")
	// case "darwin":
	// 	fmt.Println("Esto es Darwin")
	// default:
	// 	fmt.Printf("%s \n", os)
	// }

	variable1, variable2 := ejercicios.Publica("150")
	fmt.Println(variable1)
	fmt.Println(variable2)
}
