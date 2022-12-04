package main

import (
	"fmt"
)

func main() {
	//Normal Switch
	value := 0
	fmt.Scanln(&value)
	switch value {
	case 1:
		fmt.Println("i'm here")

	case 2:
		fmt.Println("i'm here")

	case 3:
		fmt.Println("i'm here")

	case 4:
		fmt.Println("i'm here")

	case 5:
		fmt.Println("i'm here")

	case 6:
		fmt.Println("i'm here")
	default:
		fmt.Println("Print me if value is not present in swith case")
	}

	//Multi-case
	day := 5

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}
}
