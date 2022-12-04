package main

import (
	"fmt"
)

func main() { //reverse
	var name string = "siri"
	new_name := " "

	for _, i := range name {
		new_name = string(i) + new_name
	}
	fmt.Println(new_name)
}

// convert slice to string

// func main() {
// 	names := []string{"siri", "lakshmi"}

// 	new_name := strings.Join(names, " ")

// 	fmt.Println(new_name)
// }
// string to slice
// func main() {
// 	name := "siri"
// 	new_name := strings.Split(name, "")
// 	fmt.Println(new_name)
// 	fmt.Printf("%T", new_name)
// }
