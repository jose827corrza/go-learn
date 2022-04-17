package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

func (fEmployee FullTimeEmployee) getMessage() string {
	return "FullTime employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary employee"
}

//Para usar interface, y la herencia de modo Go es asi...
type PrintInfo interface {
	getMessage() string
}

//Y ahora la funcion que va a recibir la interfaz que combino los "metodos" de cada clase...
func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	tEmployee := TemporaryEmployee{}
	fEmployee := FullTimeEmployee{}
	getMessage(tEmployee)
	getMessage(fEmployee)
}
