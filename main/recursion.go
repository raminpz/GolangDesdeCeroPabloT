package main

import "fmt"

func main() {
	exponencia(2)
}

// Recursion es uba funcion que se llama a sí mismo
func exponencia(numero int) {
	if numero > 10000 {
		return
	}
	fmt.Println(numero)

	exponencia(numero * 2) // Esta funcion se llama a sí mismo
}
