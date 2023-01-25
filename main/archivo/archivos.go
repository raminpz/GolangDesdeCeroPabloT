package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	leoArchivo()
	leoArchivo2()
	graboArchivo()
	graboArchivo2()
}

/* Dos formas de leer un archivo*/
// Funcion para leer archivo
func leoArchivo() {
	archivo, err := os.ReadFile("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Println(string(archivo))
	}
}

// Metodo para leer archivo, pero esta con bufio con scaner lo scanea lo lee y lo convierte a texto
func leoArchivo2() {
	archivo, err := os.Open("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan() {
			registro := scanner.Text()
			fmt.Printf("Linea > " + registro + "\n")
		}
	}
	archivo.Close()
}

/* 2 Formas de guardar un archivo*/
func graboArchivo() {
	archivo, err := os.Create("./nuevoArchivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
		return
	}
	fmt.Fprintln(archivo, "Esta es una linea nueva")
	archivo.Close()
}

func graboArchivo2() {
	fileName := "./nuevoArchivo.txt"
	if Append(fileName, "\nEsta es una segunda linea") == false {
		fmt.Println("Error en la segunda linea")
	}
}

func Append(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}
	return true
}
