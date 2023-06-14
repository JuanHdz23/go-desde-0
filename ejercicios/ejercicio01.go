package ejercicios

import (
	// "fmt"
	"strconv"
)

func Publica(valor string) (int, string) {
	// conversion, error := strconv.Atoi(valor)

	// if error != nil {
	// 	fmt.Println("Error during conversion")
	// }

	conversion, _ := strconv.Atoi(valor)
	if conversion > 100 {
		// fmt.Println("Es mayor a 100")
		return conversion, "Es mayor a 100"
	} else {
		// fmt.Println("Es menor a 100")
		return conversion, "Es menor a 100"
	}
}
