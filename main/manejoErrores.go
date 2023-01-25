package main

import (
	"log"
)

func main() {
	/**
	archivo := "prueba.txt"
	f, err := os.Open(archivo)

	defer f.Close()

	if err != nil {
		fmt.Println("Error abriendo el archivo")
		os.Exit(1)
	}
	*/
	ejemploPanic()
}

func ejemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrío un error que generó Pánic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontró el valor de 1")
	}
}
