package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 2, 3, 4}
	//Length and capacity always same for go arrays
	fmt.Println("length:", len(arr1))
	fmt.Println("capacity:", cap(arr1))
	fmt.Println("array:", arr1)
	// 	for i:=0; i<len(arr1);i++{
	// 		fmt.Println(arr1[i])
	// 	}

	//Using array
	my_slice := arr1[0:2]
	fmt.Println("length:", len(my_slice))
	fmt.Println("capacity:", cap(my_slice))
	fmt.Println("array:", my_slice)

	//Normal
	my_slice2 := []int{5, 6, 7, 8}
	for i := 0; i < len(arr1); i++ {
		my_slice2 = append(my_slice2, arr1[i])
	}
	fmt.Println(my_slice2)

	//Using Make()
	my_slice3 := make([]int, 5)
	// fmt.Println(my_slice3)
	for i := 0; i < 5; i++ {
		my_slice3 = append(my_slice3, i)
		//changing elements
		my_slice3[i] = my_slice2[i]
	}
	// Access Elements of Slice
	fmt.Println(my_slice3)
	//changing elements
	my_slice3[3] = 100

	// copy()
	// copy(dest,src)
	my_slice4 := make([]int, 4)
	copy(my_slice4, arr1)
	fmt.Println(my_slice4)
}
