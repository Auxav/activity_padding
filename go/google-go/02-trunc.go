package main

import "fmt"

func main() {
	var input float32

	fmt.Printf("Input a floating point number: ")
	fmt.Scan(&input)
	fmt.Printf("You gave %f, the truncation is %d\n", input, int(input))
}
