package main

import "fmt"

func main() {
	//Arrays
	var array [4]int
	array[0] = 70
	array[1] = 30
	fmt.Println(array, len(array), cap(array))

	slice := []int{9, 8, 7, 6, 5, 4}
	fmt.Println(slice, len(slice), cap(slice))
	newSlice := []int{1, 2, 3}
	slice = append(slice, newSlice...) //Esos puntos significan que "descomprime y toma los valores dentro"
	fmt.Println(slice)
}
