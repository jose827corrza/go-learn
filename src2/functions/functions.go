package main

import (
	"fmt"
	"time"
)

func main() {
	// x := 5

	// z := func() int {
	// 	return x * 2
	// }() //Los parentesis para que se ejecute
	// fmt.Println(z)

	//La gracia es hacer anonimas de algo que solo se realizra una vezz, no romper el patron DRY
	//Dont repeat yourself

	//USANDO GO ROUTINES
	//Crea un canal
	c := make(chan int)
	go func() {
		fmt.Println("Starting long process..")
		time.Sleep(5 * time.Second)
		fmt.Println("End...")
		c <- 1 //Enviar el canal 1
	}()
	//Bloquear el programa para que quede esperando
	<-c
}
