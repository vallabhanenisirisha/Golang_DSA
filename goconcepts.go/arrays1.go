package main

import (
	"fmt"
)

func main() {
	var arr = [20]int{}
	var even = []int{}
	var odd = []int{}
	even_sum := 0
	odd_sum := 0

	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		x := arr[i]
		if x%2 == 0 {
			even = append(even, x)
		} else {
			odd = append(odd, x)
		}
		fmt.Println(arr[i])
	}
	fmt.Println(even)
	fmt.Println(odd)

	for i := 0; i < len(even); i++ {
		even_sum += even[i]
		odd_sum += odd[i]
	}
	fmt.Println(even_sum)
	fmt.Println(odd_sum)

	if even_sum > odd_sum {
		fmt.Println("even numbers won the game")
	} else {
		fmt.Println("odd numbers won the game")
	}

}
