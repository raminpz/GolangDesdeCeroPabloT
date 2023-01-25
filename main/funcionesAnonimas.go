package main

import "fmt"

// Funcion anónima en una variable, en este caso Calculo
var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println("Suma 5 +7 = %d\n", Calculo(5, 7))

	// Restamos, usamos el mismo metodo redefinido
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Println("Restamos 6-4 = %d\n", Calculo(6, 4))

	// Dividimos, usamos el mismo metodo redefinido
	Calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}
	fmt.Println("Dividimos 100/10 = %d\n", Calculo(100, 10))

	Operaciones()

	// Closures
	tablaDel := 2
	MiTabla := Tabla(tablaDel)
	for i := -1; i < 11; i++ {
		fmt.Println(MiTabla())
	}
}

// Otro tipo de hacer funcion
func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}
	fmt.Println(resultado())
}

// Closure
func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}
