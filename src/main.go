package main

import (
	"go-basic/src/mypackage"

	"fmt"
)

func main() {
	fmt.Println("hola")
	var car mypackage.Car
	car.Brand = "Mazda"
	fmt.Println(car)
}
