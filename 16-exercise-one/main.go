package main

import (
	"fmt"
)

func main(){
	header:= `
====================================
| Contador de numeros impares      |
====================================
	`
	fmt.Println(header)
	// ingresa primer valor
	fmt.Println("Ingrese primer valor entero: ")
	var n1 int
	fmt.Scanln(&n1)
	// Ingresa segundo valor
	fmt.Println("Ingrese segundo valor entero: ")
	var n2 int
	fmt.Scanln(&n2)

	var counter int

	for  i:= n1; i<= n2; i++ {
		if i%2 != 0 {
			counter++
			fmt.Printf(" %d es un número impar \n", i)
		}
	}

	footer :=`
====================================
|	Resultado del contador	   |
====================================`
	fmt.Println(footer)
	fmt.Printf("Entre el %d y el %d hay %d numeros impares.\n", n1, n2, counter)
}
