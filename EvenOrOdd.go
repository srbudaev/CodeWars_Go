package main

import "fmt"

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	var result string
	result = EvenOrOdd(number)
	fmt.Print(result)
}
