package main

import (
	"fmt"
)

func main() {
	//slicen son smilares a los array pero funcionana diferente, ya que estos si son mutables
	// y podemos agregarle mas longitud y elementos a un slicen ya defeinidop

	numeros := []int{1, 2, 3}

	fmt.Println(numeros, len(numeros))

	// par agregar elementos aun slicen se utiliza la funcion append
	numeros = append(numeros, 4)
	numeros = append(numeros, 5)
	fmt.Println(numeros, len(numeros))

	//subslicen
	subslicen := numeros[:2]
	numeros[0] = 100
	fmt.Println(subslicen)
	fmt.Println(numeros)

	// esto quuiere decir que los slicen son refenciados de array padres pore so cuando se modifica el arrreglo padre los slince son modificados

	// caracterusticas de slicen.

	//Creo esta nota para valiar cambio en git remoto

	// las caracteristicas de un slicen
	// Puntero = para saber donde inicia y termina
	// logitud =
	//Capacidad =

	meses := []string{"Enero", "Febrero", "Marzo"}

	fmt.Printf(" len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)

	meses = append(meses, "Abril")

	fmt.Printf(" len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)

}
