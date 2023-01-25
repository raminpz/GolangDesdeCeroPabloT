package main

import "fmt"

func main() {
	/**
	fmt.Println(uno(5))
	numero, estado := dos(1)
	fmt.Println(numero)
	fmt.Println(estado)

	*/

	fmt.Println(calculo(5, 10))
	fmt.Println(calculo(2, 30, 34, 4))
	fmt.Println(calculo(10, 10, 343, 5, 1, 0))
	fmt.Println(calculo(9))

}

func uno(numero int) int {

	return numero * 2
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}

}

func calculo(numero ...int) int {
	total := 0
	for item, num := range numero {
		total = total + num
		fmt.Printf("\n Item %d \n", item)
	}
	return total
}
