package main

import "fmt"

func main() {
	var nombre1 string
	nombre1 = "jeicob "

	var nombre2 string = "josue "
	// := al asinar variales de esa foram go hace que al variable tome el tipo de dato qeu se esta asignando

	edad := 26

	fmt.Println(nombre1, nombre2, edad)

	//variables para tipos de datos

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)
	/* al no asignar valores se asssigna valores inicles el cual es cero, en go los valores pro defecto
	no son ni null ni undefined	todas la variable sya tiene un valro inicializado.*/

	const pi = 3.1416
	fmt.Println(pi)

}
