package main

import (
	"fmt"
)

func main() {

	//taking the normal variable
	var a int = 120

	//declaration for pointer
	var b *int

	//initialization pointer
	b = &a

	//displaying the result
	fmt.Println(b)
}
