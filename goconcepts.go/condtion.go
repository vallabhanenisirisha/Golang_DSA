package main

import "fmt"

//Using Conditional statements
func main() {
	if sum > comp {
		fmt.Printf("HI i'm greater then %d", comp)
	} else if sum < comp {
		fmt.Println("Hello i'm smaller")

	} else {
		fmt.Println("Equal")
	}

	// Program to find greatest of three numbers:
	var (
		a int
		b int
		c int
	)

	fmt.Println("Enter the inputs:")
	fmt.Scanln(&a, &b, &c)

	if a >= b && a >= c {
		fmt.Printf("%d is greatest number", a)
	} else if b >= a && b >= c {
		fmt.Printf("%d is greatest number", b)
	} else {
		fmt.Printf("%d is greatest number", c)
	}

