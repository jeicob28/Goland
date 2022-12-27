/*
Practica 01: Suma de dos números enteros
Enunciado: dado dos números enteros, hallar la suma.
*/
package main

import "fmt"

func sumar(a1, b1 int) int {
	return a1 + b1
}

func main() {
	a1, b1 := 0, 0
	fmt.Print(" Ingrese el un numero a sumar: ")
	fmt.Scanln(&a1)
	fmt.Print(" Ingrese el un numero a sumar: ")
	fmt.Scanln(&b1)
	rsultado := sumar(a1, b1)
	fmt.Println("el resultado de la suma es de: ", rsultado)

}
