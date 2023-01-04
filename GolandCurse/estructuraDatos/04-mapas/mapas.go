package main

import "fmt"

func main() {

	// los mapas son una coleccion de datos, son elemtos desordenados se define pro claver y valor

	dias := make(map[int]string)

	fmt.Println(dias)

	// Para agregar datos a los mapas
	// en los mapa no ponemos los indices si no que colocan las claves y puedn se orndenada o desordenadas

	dias[1] = "Domingo"
	dias[2] = "Lunes"
	dias[3] = "Martes"
	dias[4] = "Miercoles"
	dias[5] = "Jueves"
	dias[6] = "Viernes"
	dias[7] = "Sabado"
	fmt.Println(dias)

	// para ralizar alguna modificacion en los mapas se accede mediante las claves

	dias[7] = "jeicob"
	fmt.Println(dias)

	// en los mapas tambien podemo eliminar datos cn la  funcion delete

	delete(dias, 7)
	fmt.Println(dias)

	// tambien a los mapas podemos asignarle s tamaño con la funcion len

	fmt.Println(dias, len(dias))

	// los mapas pueden llegar a convertirce una estructura de datos mas compleja

}
