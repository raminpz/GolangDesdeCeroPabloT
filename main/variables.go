package main

import "fmt"

// Variables
var numero int
var texto string
var status bool

func main() {

	numero := 4
	texto := "Hola Mundo"
	status := true

	fmt.Println(numero, texto, status)

	numero2, numero3, numero4, texto := 1, 2, 3, "Hola Rami"
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(texto)

}
