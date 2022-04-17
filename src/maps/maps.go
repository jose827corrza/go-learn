package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["Jose"] = 1
	m["Nubia"] = 2

	fmt.Println(m)

	//Recorre el map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar el valor
	value, ok := m["Estanislavo"]
	fmt.Println(value, ok)
	value2, ok2 := m["Jose"]
	fmt.Println(value2, ok2)
}
