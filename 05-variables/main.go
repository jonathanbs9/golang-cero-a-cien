package main

import(
	"fmt"
	"log"
)

func main(){

	var number int
	var name string
	var dni int
	var (
		jugador1 = "Mess1"
		jugador2 = "Ronald0"
	)
	number = 250
	mult := 10 * 23
	name = "This is a string value"
	isFlag := !true
	_nombre := "Jonathan"
	userName := "jonathanbs"
	// := Se utiliza en variables nuevas. Solamente va a aceptar string, en este caso.
	lastName := "keyboard"
	lastName, dni = "Schroeder", 9087

	fmt.Println(number,"\n", name, lastName, dni, mult, isFlag)
	log.Println(_nombre + " | " + userName)
	log.Println("Los candidatos para el balon de oro son => " + jugador1 + " | " + jugador2)


}

