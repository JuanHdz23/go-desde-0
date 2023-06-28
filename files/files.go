package files

import (
	// "bufio"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/juanhdz23/go-desde-0/ejercicios"
)

var pathFileName = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.Publica2()
	archivo, err := os.Create(pathFileName)
	if err != nil {
		fmt.Println("Error al crear el archivo", err.Error())
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.Publica2()
	if !Append(pathFileName, texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(file string, texto string) bool {
	if _, err := os.Stat(pathFileName); os.IsNotExist(err) {
		arch, err := os.Create(pathFileName)
		if err != nil {
			fmt.Println("Error al crear el archivo", err.Error())
		}
		arch.Close()
	}
	arch, err := os.OpenFile(pathFileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append", err.Error())
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString", err.Error())
		return false
	}

	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := ioutil.ReadFile(pathFileName)
	if err != nil {
		fmt.Println("Error al leer el archivo", err.Error())
		return
	}

	fmt.Println(string(archivo))
}

func LeoArchivo2() {
	archivo, err := os.Open(pathFileName)
	if err != nil {
		fmt.Println("Error al leer el archivo", err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println(">", registro)
	}
	archivo.Close()
}
