package main

import "fmt"

/** Tipos de datos

BASICOS
	- Numeros
	- Cadenas
	- Boolenas

CONJUNTOS (SET)
	- Arreglos
	- Estructura

REFERENCIAS
	- Punteros
	- Segmentos
	- Mapas
	- Funciones
	- Canales

INTERFACES
	-
 */

func main (){
	// Soportan enteros con signo
	//var integer8 int8		// Desde -128 a 127
	//var integer16 int16		// Desde -32768 a 32767
	var integer32 int32		// Desde -2147483648 a 2147483647
	var integer64 int64		// Desde -9223372036854775808 a 9223372036854775807

	// Soportan enteros sin signo
	//var uinteger8 uint8		// Desde 0 a 255
	//var uinteger16 uint16	// Desde 0 a 65535
	//var uinteger32 uint32	// Desde 0 a 4294967295
	//var uinteger64 uint64	// Desde 0 a 18446744073709551615

	// Alias
	//var enteroByte byte		// alias para uint8
	var enteroRune rune			// alias para int32

	// Tipos dependiente de la plataforma
	//var enteroUint uint			// 32 o 64
	//var enteroInt int			// 32 o 64
	//var enteroUintptr uintptr	// Entero sin signo lo sufic. grande para contener el valor de un puntero

	integer32 = 2361986
	integer64 = 3123812
	fmt.Println(integer32 + int32(integer64))
	fmt.Println(integer32 + enteroRune)

}
