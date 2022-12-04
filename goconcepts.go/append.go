package main

import (
	"fmt"
)

// func main() {
// 	myslice1 := []int{1, 2, 3, 4, 5, 6}
// 	fmt.Printf("myslice1 = %v\n", myslice1)
// 	fmt.Printf("length = %d\n", len(myslice1))
// 	fmt.Printf("capacity = %d\n", cap(myslice1))

// 	myslice1 = append(myslice1, 20, 21)                                   // append the array item
// 	fmt.Printf("myslice1 = %v\n", myslice1)
// 	fmt.Printf("length = %d\n", len(myslice1))
// 	fmt.Printf("capacity = %d\n", cap(myslice1))
// }

//we can append one slice to another

func main() {
	myslice1 := []int{1, 2, 4}
	myslice2 := []int{7, 8}
	myslice3 := append(myslice1, myslice2...)

	fmt.Println("my new array is :", myslice3)
	// fmt.Println("length = %d\n", len(myslice3))
	// fmt.Println("capacity =%d\n", cap(myslice3))

}
