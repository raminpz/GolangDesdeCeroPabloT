package main

import "fmt"

func main() {
	/**
	numero := 0
	for {
		fmt.Println("Continuo")
		fmt.Println("Ingrese el numero secreto")
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}

	}


	var i = 0

	for i < 10 {
		fmt.Printf("\n Valor de i: %d, i")
		if i == 5 {
			fmt.Printf("Multiplicamos por 2 \n")
			i = i * 2
			continue
		}
		fmt.Printf("      Pasó por aquí \n")
		i++

	}
	*/

	var i int = 0
RUTINA:
	for i < 10 {
		if i == 4 {
			i = i + 2
			fmt.Println("Voy a RUTINA sumando 2 a i")
			goto RUTINA
		}
		fmt.Printf("Valor de i: %d\n", i)
		i++
	}

}
