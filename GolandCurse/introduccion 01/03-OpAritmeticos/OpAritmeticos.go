package main

import "fmt"

func main() {
	a := 20
	b := 10

	//suma
	result := a + b
	fmt.Println("El valor de la suma es: ", result)
	//resta
	result = a - b
	fmt.Println("El valor de la Resta es: ", result)

	//multiply
	result = a * b
	fmt.Println("El valor de la Multiplicacion  es: ", result)
	//divide
	result = a / b // ya que las variables esta toamdno valores enteros por eso muestra que la dicion entre a y b es de 2
	fmt.Println("El valor de la division es: ", result)

	// si se require hace runa division exacta, esto queire decir qeu lso valores que se debe asignar a las variable deben ser del mismo tipo es decir:

	var div float64 = 3.0 / 2.0
	fmt.Println("El valor de la division  es: ", div)

	//moduloDivision
	result = a % b
	fmt.Println("El valor deL Modulo es:", result)
}
