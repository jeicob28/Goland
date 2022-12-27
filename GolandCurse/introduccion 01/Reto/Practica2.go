/*
Practica 02: Calcular cociente y el Residuo de dos números enteros
Enunciado: halar el cociente y el residuo (resto) de dos números enteros.

Análisis: para la solución de este problema, se requiere que el usuario ingrese dos números enteros y el sistema realice el cálculo respectivo para hallar el cociente y residuo.
*/
package main

import "fmt"

func cociente(a, b int) int {
	return a / b
}

func residuo(a, b int) int {
	return a % b
}

func main() {
	var a, b, cos, res int

	//Ingrso de datos:
	fmt.Println("ingrsar el primer numero N01= ")
	fmt.Scanln(&a)

	fmt.Println("ingrsar el segundo numero N02= ")
	fmt.Scanln(&b)

	//Calcula el cocietne  llamado de la funcion.
	cos = cociente(a, b)
	// Calcula el resudo llamando dicha funcion.
	res = residuo(a, b)

	fmt.Println("El valor del cociente es: ", cos)
	fmt.Println(" el valor del residuo es: ", res)
}
