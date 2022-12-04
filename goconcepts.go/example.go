package main

import (
	"fmt"
	"strconv"
)

func main() {
	name := "NTR"
	y := 31
	z := strconv.Itoa(y) // to convert int to string and connactinate to string we
	// we use strconv.Itoa
	fmt.Println(name + " " + z)

	fmt.Printf("result: %s", name+" "+z)

}
