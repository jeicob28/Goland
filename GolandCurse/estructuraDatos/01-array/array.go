// Go al ser un leguaje de programacion tipado y de tipado estatico, al trabajar con arreglos se allmacenan cantidades de datos.
package main

import "fmt"

func main() {
	//los rray se define igual qeu las variable y se agrega corchete al finalizar y el tipo de datos que este almacenara al igual que la cantidad.
	// los array son inmutables ya que se definen el tamaño del arrglo
	var numeros [5]int
	fmt.Println(numeros)

	// una ofrm ade modificar los indices en loqe u se alamcenan los datos en go es escribir el arreglo y en corcheste el indice que se queire modificar.

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	// par aver la modificacioen realizadas se corre la linea de impresion en consola

	fmt.Println(numeros)
	// y para imprimir un solo valro de los elemento del array
	fmt.Println(numeros[2])

	// tambien podemos ddefinir array con sus valores definidos
	nombres := [3]string{"jeicob", "raquel", "josue"}
	fmt.Println(nombres)

	// si quiero cantidad de x elentos aun arreglo pero no se cuantos elemtos voy aagregar

	colores := [...]string{
		"rojo",
		"verde",
		"negro",
		"azul",
	}
	fmt.Println(colores, len(colores))

	//podemos definir los indices que tiene un array

	monedas := [...]string{
		0: "Dolar",
		2: "Peso Colombiano",
		3: "soles",
		5: "Euro",
	}
	fmt.Println(monedas, len(monedas))

	// podemos tener array dentro de otro array
	subMoneda := monedas[0:3]
	// en esta declaracion estoy tomando los datos par ael sudarreglo del arreglo prncipal
	//e imprimo solo el contenido desde lo sindices que se quiere tomar.
	fmt.Println(subMoneda)
	//tambien puedo tomar los datos de diferentes formas  par ael subarreglo
	subMoneda1 := monedas[3:] // de esta forma slo mostraremos el contenido desde el idice 3 incluido hasta el final
	fmt.Println(subMoneda1)

}
