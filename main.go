package main

import (
	"fmt"
	"goModDemo/area"
	"goModDemo/fibanocci"
	"goModDemo/squareroot"
)

func init() {
	fmt.Println("Main module 1st init")
}

func main() {
	fmt.Println("Main module main func")
	fmt.Println("Square root Value: ", squareroot.Calculate(49))
	fmt.Println("Area Value: ", area.Calculate(50.2, 69.5))
	fmt.Println("Fibonacci sequence: ", fibanocci.Sequence(10))
}
