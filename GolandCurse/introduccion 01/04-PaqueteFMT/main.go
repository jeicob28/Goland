package main

import "fmt"

func main() {
	hola := "Hola"
	mundo := "Mundo"
	fmt.Print(hola)    // prints que no realiaza salto de linea
	fmt.Println(mundo) // prints que  realiaza salto de linea

	nombre := "jeicob"
	edad := 31

	fmt.Printf("hola, %s y su edad es %d \n", nombre, edad)

	//fmt.Printf es utilizado para imprimir formateando la salida, tambien se utilizan simbilo y letras %s (indica que es un valor stringd)
	// %d indica que es un valor entero
	// %v se indica cuando no se sabe el tipo de valor que se espera

	fmt.Printf("hola, %v y su edad es %d  \n", nombre, edad)

	mensaje := fmt.Sprintf("hola, %v y su edad es %d  \n", nombre, edad)
	// se utiliza la funcion fmt.sprintf para almacenar todo el mensaje en uan variable, este formatea todo los datos y lo alamcena en la variable
	fmt.Println(mensaje)

	fmt.Printf("nombre: %T\n", nombre)
	// %T Se utiliza el caracte especial para saber que tipo de dato es la variable

	fmt.Print(" Ingrese otro nombre: ")
	fmt.Scanln(&nombre)
	fmt.Println("el otro nombre: ", nombre)
	// De esta forma ingresamos datos desde l aconsola y lo asignamso a uan variable
}
