package main

import "fmt"

//Esta interface agrupa las clases que tienen comportamientos parecidos, para que sea mas legible todo al usarlo al mismo tiempo,
//que el codigo sepa en base a que tipo se esta manejando que funciones aplican para ese caso puntiualmente.
type figuras2D interface {
	area() float64
}

//Aca estan las clases que son diff, pero q comparten comportamientos
type Cuadrado struct {
	base float64
}
type Rectangulo struct {
	base   float64
	altura float64
}

//Son las funciones de cada tipo de clase
func (c Cuadrado) area() float64 {
	return c.base * c.base
}
func (r Rectangulo) area() float64 {
	return (r.base * r.altura) / 2
}

//Es una funcion que recibe como parametro la interface que recoge las clases que se parecen
func calcular(f figuras2D) {
	fmt.Println("El area de la figura es: ", f.area())
}

func main() {
	myCuadrado := Cuadrado{base: 5}
	myRectangulo := Rectangulo{base: 5, altura: 3}

	calcular(myCuadrado)
	calcular(myRectangulo)

	//Normalmente no se puede de diff tipos, pero usando interfaces se logra, esta es la forma de llevarlo a cabo
	slice := []interface{}{5, "Jose", 3.1416}
	fmt.Println(slice)
}
