// Go supports methods defined on struct types.
package main

import "fmt"

type rect struct {
	width, height int
}

// a pointer receiver
func (r *rect) area() int {
	return r.width * r.height
}

// a value receiver
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatically handles conversion between values and pointers for method calls.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
