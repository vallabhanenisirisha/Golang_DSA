package main

import "fmt"

// func main() {
// 	names := []string{}

// 	var value int
// 	fmt.Println("Enter your value: ")
// 	fmt.Scanln(&value)

// 	for i := 0; i < value; i++ {
// 		var new_name string
// 		fmt.Println("Enter your names: ")
// 		fmt.Scanln(&new_name)

// 		names = append(names, new_name)
// 	}
// 	fmt.Println(names)
// }

// user input from user

// func main() {
// 	fmt.Print("enter your name: ")
// 	var name string
// 	fmt.Scanln(&name)
// 	fmt.Println(name)
// }

// taking input elements form array

func main() {
	fmt.Println("Enter size of your array: ")
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		// fmt.Printf("Enter %dth element: ", i)
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("Your array is: ", arr)
}
