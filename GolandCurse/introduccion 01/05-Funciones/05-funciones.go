package main

import "fmt"

func Saludar(nombre string) {
	fmt.Println("hola: ", nombre)
}
func sumar(a, b int) int {
	// se debe espesificar el tipo de datio que va a retornar  la funcion y se realiza depues de la declaracion delas variable y antes de la lalves de la funcion
	return a + b
}
func main() {
	Saludar("jeicob")
	r := sumar(10, 20)
	fmt.Println("el resultado de la funcion suma es", r)
}

