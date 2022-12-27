/*
Practica 03: Calcular el Precio de Venta
Enunciado: dado el valor de venta de productos, hallar el IGV (18%) y el precio de venta.

Análisis: para la solución de este problema, se requiere que el usuario ingrese el valor de venta del producto y el sistema realice el cálculo respectivo para hallar el IGV y el precio de venta.
*/

package main

import "fmt"

func calcularIGV(valorVenta, igv float64) float64 {
	return igv * valorVenta
}
func calcularVentaIGV(valorVenta, igv float64) float64 {
	return valorVenta + igv
}

func main() {
	// Declaracioens de variables
	var valorVenta, igv, valorPublico float64
	var igvD = 0.18

	// Ingreso de dato
	fmt.Println("ingrse el valor de venta:")
	fmt.Scanln(&valorVenta)

	// operacioens
	igv = calcularIGV(valorVenta, igvD)

	valorPublico = calcularVentaIGV(valorVenta, igv)

	//Salida de datos
	fmt.Println("==========================")
	fmt.Println("Valor de Venta:", valorVenta)
	fmt.Println("IGV:", igv)
	fmt.Println("Precio de Venta:", valorPublico)
	fmt.Println("==========================")
}
