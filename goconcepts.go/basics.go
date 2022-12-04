// Syntax::

// Package declaration
package main

// Import packages
import (
	"fmt"
)

// Functions
func main() {
	// Data types and varables:
	// data types -> {int,float32,string,bool}

	// 1 -->
	// Declaring the variables:::
	// Using var keyword::
	// var variablename type
	var firstname string = "vyshu"
	// fmt.Println(firstname)

	// 2
	// Declaring variables using :=
	secondname := "Hello"
	// fmt.Println(secondname)
	fmt.Println(secondname + " " + firstname)

	// 3
	// Multiple variable declaration
	var a, b, c int = 1, 2, 3
	// grouping together
	var (
		d int    = 1
		e string = "cat"
	)
	fmt.Println(a, b, c, d, e)

	// 4
	// Using const keyword::
	const number int = 20 // const keyword is used to give fixed value to any variable
	// number =30 // we can't assign new value to the variable declare with const keyword
	fmt.Println(number)
	// grouping Together
	const (
		name  = "Luffy"
		power = 100000
	)
	fmt.Println(name, power)

	// 5
	// Types of print() Functions
	fmt.Print(a, firstname)                                // next print statement is coming in same line and no space b/w the variable like println().
	fmt.Println(b, secondname)                             // Run to understant the difference.
	fmt.Printf("hi i am %v my age is %v \n", name, number) // like python format function formats the variables in given order.
	// %v to get value in variable, %T to get type of the variable.
	// %b	Base 2
	// %d	Base 10
	// %s	Prints the value as plain string
	// %q	Prints the value as a double-quoted string
	//Try this values with printf().

	//6
	// Arrays in go
	// Declaring array with var keyword
	//Syntax
	// var array_name = [length]datatype{values} //length is given
	// var array_name = [...]datatype{values}	//length will be given while compiling
	//With :=
	//array_name := [length]datatype{values}
	//array_name := [...]datatype{values}
	var arr1 = [3]int{1, 2, 3}
	arr2 := [...]int{5, 6, 7}
	fmt.Println(arr1, arr2)
	//Access Elements of an array
	x := arr1[0]
	fmt.Println(x)
	fmt.Println(arr2[0])
	// changing Elements in Array
	arr1[2] = 100
	fmt.Println(arr1)
	//Array Initialization
	arr3 := [5]int{}              //not initialized
	arr4 := [5]int{1, 2}          //partially initialized
	arr5 := [5]int{1, 2, 3, 4, 5} //fully initialized
	fmt.Println(arr3, arr4, arr5)
	// Initialize Only Specific Elements
	arr6 := [5]int{1: 10, 2: 40}
	fmt.Println(arr6)
	// Find the Length of an Array
	// we use len() to find length
	fmt.Println("length:", len(arr1))

	// 6
	// In Go, there are several ways to create a slice:
	// Using the []datatype{values} format
	// Create a slice from an array
	// Using the make() function
	// syntaxs:
	// slice_name := []datatype{values}
	names := []string{"anya", "yori", "loid"}
	fmt.Println(names)
	// It has 2 functions len(),cap()
	fmt.Println(len(names), cap(names))
	//Using arrays
	// slice_name := array[start:end] //slice made from array
	new_slice := arr5[1:4]
	fmt.Println(new_slice)
	//Using make()
	// slice_name := make([]type, length, capacity)
	my_slice := make([]int, 5)
	fmt.Println(my_slice)

}
