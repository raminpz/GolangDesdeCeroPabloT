package main

import (
	"Golang/main/poo/user"
	"fmt"
	"time"
)

type pepe struct {
	user.Usuario
}

func main() {
	u := new(pepe)
	u.AltaUsuario(1, "Ramiro", time.Now(), true)
	fmt.Println(u.Usuario)

}
