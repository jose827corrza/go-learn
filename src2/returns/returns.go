package main

import (
	"fmt"
)

func sum(a, b int) int {
	return a + b
}

//Pero si no se sabe cuantos valores tendremos de argumentos se hace asi, variaticos se llama esto
/*
Funcion variadicas
*/
func sumPro(values ...int) int {
	total := 0

	for _, num := range values {
		total += num
	}
	return total
}

func nameList(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

//RETORNOS CON NOMBRE
func getValues(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	//Funciones variadicas

	fmt.Println(sum(1, 2))
	fmt.Println(sumPro(1, 2, 3, 4, 5))
	fmt.Println(getValues(2))
}
