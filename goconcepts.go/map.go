package main

import (
	"fmt"
)

/*syntax
var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
b := map[KeyType]ValueType{key1:value1, key2:value2,...} */

// func main() {
// 	var a = map[string]string{"name": "siri", "initial": "v", "location": "gunutr"}

// 	fmt.Println(a)
// 	fmt.Printf("%v", a)
// }

// creating map using make() function

// func main() {
// 	var a = make(map[string]string) // The map is empty now
// 	a["brand"] = "ford"
// 	a["year"] = "1964"

// 	fmt.Println(a)
// }

func main() {
	a := map[string]string{}

	fmt.Println(a == nil)

}
