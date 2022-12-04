package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	// Arithmetic operators
	// result := 10+10 //
	// result := 10-10 //
	// result := 10/10 //
	// result := 10%10 //
	// result++		//
	// result--     //
	// Assigment Operators
	// (=,+=,-=,/=,%=)
	// Comparison Operators
	// (==,!=,<,>,<=,>=)
	//Logical operators
	// (&&,||,!)

	// Taking input from user:
	var (
		value1 int
		value2 int
	)
	fmt.Println("Enter input values:")
	fmt.Scanln(&value1, &value2)
	sum := value1 + value2
	fmt.Println(sum)
	comp := 50
	//Using Conditional statements
	if sum > comp {
		fmt.Printf("HI i'm greater then %d", comp)
	} else if sum < comp {
		fmt.Println("Hello i'm smaller")

	} else {
		fmt.Println("Equal")
	}

}
