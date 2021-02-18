package main

import(
	"fmt"
)

func main(){
	var number int
	number = 250
	var name string
	var dni int
	name = "This is a string value"
	mult := 10 * 23
	isFlag := !true

	// := Se utiliza en variables nuevas. Solamente va a aceptar string, en este caso.
	lastName := "keyboard"
	lastName, dni = "Schroeder", 9087
	fmt.Println(number,"\n", name, lastName, dni, mult, isFlag)

}
