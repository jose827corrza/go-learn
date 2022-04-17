package main

import "fmt"

type PC struct {
	ram   int
	brand string
	disk  int
}

func (myPc PC) String() string {
	return fmt.Sprintf("Tengo %d de RAM, de disco tengo %d GB y mi marca es %s", myPc.ram, myPc.disk, myPc.brand)
}

func main() {
	myPc := PC{ram: 16, brand: "Gygabyte", disk: 1000}
	fmt.Println(myPc)
}
