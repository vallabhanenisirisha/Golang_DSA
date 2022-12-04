package main

import (
	"fmt"
)

// func main() {
// 	var arr1 = [3]int{2, 3, 4}            //we can declare arry using length
// 	arr2 := [2]int{5, 6}

// 	fmt.Println(arr1)
// 	fmt.Println(arr2)
// }

// func main() {
// 	var arr1 = [...]int{3, 4, 5}                     //we can define array with inferred length
// 	arr2 := [...]int{6, 7, 8}

// 	fmt.Println(arr1)
// 	fmt.Println(arr2)
// }

// func main() {
// 	var cars = [3]string{"Tata", "Audi", "Ola"} 		//declares an array of strings

// 	fmt.Println(cars)
// 	fmt.Println(cars[0])                       			//accessing the array elements

// 	cars[2] = "Mahindra" 								//change the array element

// 	fmt.Println(cars)
// }

// func main() {                                              //initaialize an array
// 	arr1 := [3]int{}										   //not initialize
// 	arr2 := [4]int{1, 2}                                       //partially intilize
// 	arr3 := [5]int{1, 2, 3, 4, 5}							   //fully intialize
// 	fmt.Println(arr1)
// 	fmt.Println(arr2)
// 	fmt.Println(arr3)
// }

func main() { // we can intialize only specific elements of array
	arr1 := [5]int{1: 10, 2: 4}
	fmt.Println(arr1)

	fmt.Println(len(arr1)) //we can find length of array
}
