package main

import "fmt"

type serVivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}

type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
}

type vegetal interface {
	ClasificacionVegetal() string
}

/** Género Humano */
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	vivo       bool
}
type mujer struct {
	hombre
}

func (h *hombre) respirar() { h.respirando = true }
func (h *hombre) comer()    { h.comiendo = true }
func (h *hombre) pensar()   { h.pensando = true }
func (h *hombre) sexo() string {
	if h.esHombre == true {
		return "Hombre"
	} else {
		return "Mujer"
	}
}
func (h *hombre) estVivo() bool { return h.vivo }
func HumanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.sexo())
}

/** Genero Animal */
type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

func (p *perro) respirar()         { p.respirando = true }
func (p *perro) comer()            { p.comiendo = true }
func (p *perro) EsCarnivoro() bool { return p.carnivoro }
func (p *perro) estVivo() bool     { return p.vivo }

func AnimalesRespirar(an animal) {
	an.respirar()
	fmt.Println("Soy un animal y estoy respírando")
}

func AnimalesCarnivoros(an animal) int {
	if an.EsCarnivoro() == true {
		return 1
	}
	return 0
}

/* SerVivo*/
func estoyVivo(v serVivo) bool {
	return v.estaVivo()
}
func main() {
	/**
	Pedro := new(hombre)
	Pedro.esHombre = true
	HumanosRespirando(Pedro)

	Maria := new(mujer)
	HumanosRespirando(Maria)
	*/
	totalCarviros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	Dogo.vivo = true
	AnimalesRespirar(Dogo)
	totalCarviros = +AnimalesCarnivoros(Dogo)
	fmt.Printf("Toral carnivoros %d \n", totalCarviros)

	//fmt.Printf("Estoy vivo = %t", estoyVivo(Dogo)) aqui sale error volver a vereficar
}
