package mypackage

import "fmt"

type Car struct {
	Brand string
	Year  int
}

type PC struct {
	Ram   int
	Disk  int
	Brand string
}

func SeyHello(name string) string {
	return "Hello " + name
}

func (machine *PC) MultiplicarRam() {
	machine.Ram = machine.Ram * 2
}

// Stringer
func (myPc PC) String() string {
	return fmt.Sprintf("Tengo %d GB de RAM, %d GB de Disco y es una %s", myPc.Ram, myPc.Disk, myPc.Brand)
}
