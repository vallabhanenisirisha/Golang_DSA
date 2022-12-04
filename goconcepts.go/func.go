package main //package deceleration

import (
	"fmt"
) //import statement
var name string = "sirisha"

func add(x int, y int) (result int) {
	result = x + y
	return result

}
func main() {
	name = "hi"
	fmt.Println("hello world!")
	fmt.Println(name)
	x := add(5, 5)
	fmt.Println(x)

}

// func myName(fname string) {
// 	fmt.Println("Hello", fname, "Refens")
// }

// func main() {
// 	myName("Sirisha")
// 	myName("lakshmi")
// 	myName("v")
// }
