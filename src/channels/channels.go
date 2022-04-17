package main

import "fmt"

func say(text string, c chan<- string) { //Canal de entrada c chan<-
	c <- text //Si fuera canal de salida c <-chan
	//text = <- c
}

func message(text string, c chan<- string) {
	c <- text
}

func main() {
	c := make(chan string, 1)
	fmt.Println("Hello")
	go say("Jose", c)

	fmt.Println(<-c)

	x := make(chan string, 2)
	x <- "message1"
	x <- "message2"
	//Es muuuy buena practica cerrar los canales para que no se queden esperando datos
	close(x)

	fmt.Println(len(x), cap(x))

	for message := range x {
		fmt.Println(message)
	}
	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensajito1", email1)
	go message("mensajito2", email2)
	//Tener presente la cantidad de canales y su informacion
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Recibido mensaje del canal 1", m1)
		case m2 := <-email2:
			fmt.Println("Recibido mensaje del canal 2", m2)
		}
	}
	fmt.Printf("2,%T", 2)
}
