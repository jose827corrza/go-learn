package main

import (
	"Users/jose/go-learn/src/mypackage"
	"fmt"
)

func main() {
	fmt.Println("hola")
	var car mypackage.Car
	car.Brand = "Mazda"
	fmt.Println(car)
}
