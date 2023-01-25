package main

import "fmt"

func main() {
	/**
	tabla[0] = 1
	tabla[5] = 15
	Arrelo estático
	tabla := [10]int{1, 23, 45, 3, 67, 8, 15, 90, 70, 100}

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}


	// Matriz
	matriz[3][5] = 1
	fmt.Println(matriz)
	*/
	// Slice, matriz no estatico
	matriz := []int{2, 5, 3, 10}
	fmt.Println(matriz)
	variante2()
	variante3()
	variante4()
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[3:] // porcion, muestra elementos que están despúes de la posicion 3.
	// porcion:=elementos[:4] indica desde el primero elemento hasta el 4
	// porcion:=elementos[2:4] desde el segundo elemento hasta el 4
	fmt.Println(porcion)
}

func variante3() {
	elementos := make([]int, 5, 20) // sino se pone el tamaño se define Slice, 5 es el tamaño del arreglo y 20 es el espacio en memoria
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
}

func variante4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i) // El append es para agregar datos al parametro, pero se usa cuando ya se queda sin espacio, como ultimo recurso, dado que gasta mucha memoria
	}
	fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums))
}
