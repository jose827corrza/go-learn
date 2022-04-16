package main

import "fmt"

func main() {
	a := 50
	//Un puntero es como el compilador apunta a un lugar de la memoria
	//al usar & nos da la dir del espacio de mem.
	b := &a
	fmt.Println(a)
	fmt.Println(b)
	//Es asterisco es el contrario de &
	fmt.Println(*b)
	myPc := pc{ram: 16, disk: 500, brand: "Apple"}
	fmt.Println(myPc)
}

type pc struct {
	ram   int
	disk  int
	brand string
}
