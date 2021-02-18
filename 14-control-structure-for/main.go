package main

import (
	"fmt"
)

func main() {
	/*  Estructuras de Control: For (único) => No existe while, do while
	 	Imprimir los numeros del 1 al 5
		fmt.Println("1")
	 	fmt.Println("2")
	 	fmt.Println("3")
	 	fmt.Println("4")
	 	fmt.Println("5")
	*/


	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++

	}

	for j := 1; j <= 5; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("555")
		fmt.Println("444")
		fmt.Println("333")
		fmt.Println("222")
		fmt.Println("111")
		break
	}

}
