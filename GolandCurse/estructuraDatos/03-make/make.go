package main

import "fmt"

func main() {

	// para crear un slicen limpio o basio se utiliza la funcion make donde se determina el tipo de dato tamaño y capacidad
	numeros := make([]int, 3, 3)

	fmt.Println(numeros)

	numeros[0] = 100
	numeros[1] = 200
	numeros[2] = 300

	// para agregar valores se realiza con la funcion append

	numeros = append(numeros, 400)

	fmt.Println(numeros, len(numeros), cap(numeros))

}
