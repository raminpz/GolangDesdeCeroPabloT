package main

import "fmt"

func main() {
	paises := make(map[string]string, 5)
	fmt.Println(paises)

	paises["Perú"] = "Ayacucho"
	paises["Argentina"] = "Buenos Aires"

	campeonate := map[string]int{
		"Barcelona":   39,
		"Real Madrid": 38,
		"Chivas":      37,
		"Boca Junior": 30}
	fmt.Println(campeonate)

	campeonate["River Plate"] = 25    // Con esto agregamos nuevo valos al campeonato array en si
	campeonate["Chivas"] = 25         // Con esto cambiamos el valor de Chivas
	delete(campeonate, "Real Madrid") // Con este método eliminamos
	fmt.Println(campeonate)

	// Recorre mapa
	for equipo, puntaje := range campeonate {
		fmt.Printf("El equipo %s, tiene un puntaje de : %d \n", equipo, puntaje)
	}
	// Consultar si un equipo existe
	puntaje, existe := campeonate["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe: %t \n", puntaje, existe)
}
