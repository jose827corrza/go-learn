package intermediate

import "fmt"

//Struct son las "clases " de Go
type Employee struct {
	id   int
	name string
}

//Estas se llaman receiver functions, lo que vendrian siendo los metodos de la clase
func (e *Employee) setId(id int) {
	e.id = id
}

func (e *Employee) setName(name string) {
	e.name = name
}

func main() {
	e := Employee{}
	e.setId(7)
	e.setName("Jose")
	fmt.Println(e)
}
