package main

import (
	"fmt"
)

// func main() {

// 	var name string = "Siri"

// 	initial := "vallabhaneni"

// 	var student = "lakshmi"

// 	fmt.Println(name)
// 	fmt.Println(initial)
// 	fmt.Println(student)

// }

//declare multiple variables at single line

// func main() {
// 	var a, b, c int = 2, 3, 4
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// }

// If the type keyword is not specified, you can declare different types of variables in the same line:

// func main() {
// 	var a, b = "Hello", 6
// 	c, d := 7, 4.5

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)

// }

// Multiple variable declarations can also be grouped together into a block for greater readability:

func main() {
	var (
		a int
		b int    = 1
		c string = "Hello"
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
