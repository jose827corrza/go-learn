package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	//Esta le dice que lo ultimo que haga de esta funcion es que el wg este finalizado
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	//Segun entiendo, esto agrupa y empieza a organizar los dif procesos para goroutines
	var wg sync.WaitGroup

	fmt.Println("Hello")
	//Aca yo establezco que lo que sigue, o lo que usa "go" va a irse a otra goroutine aparte
	wg.Add(1)
	go say("world", &wg)
	//Aca la dice que espere a la otra goroutines iniciada
	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("Adiosito")

	time.Sleep(time.Second * 2)
}
