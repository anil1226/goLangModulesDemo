package main

import (
	"fmt"
	"goModDemo/strings"
	"goModDemo/structs"
)

func init() {
	fmt.Println("Main module 1st init")
}

func hello(b ...int) {
	//fmt.Println("a :", a)
	fmt.Println("b:", b)
}

func main() {
	// fmt.Println("Main module main func")
	// fmt.Println("Square root Value: ", squareroot.Calculate(49))
	// fmt.Println("Area Value: ", area.Calculate(50.2, 69.5))
	// fmt.Println("Fibonacci sequence: ", fibanocci.Sequence(10))
	//fibanocci.Infinite()
	//slices.Testswitch()
	// slices.DemoSlice()
	//hello()
	// variadic.Find(1, 2, 3)

	//variadic.Find(4, 1, 2, 3, 4)

	// nums := []int{2, 6, 1, 2}

	// variadic.Find(6, nums...)

	//maps.Create()

	// strings.Run()

	// h := "hello"
	// fmt.Println(strings.Convert(([]rune(h))))

	fmt.Println(strings.IsPalindrome("madam"))
	fmt.Println(strings.IsPalindrome("anil"))

	p := structs.Person{
		Name: "Naveen",
		Age:  50,
		Address: structs.Address{
			City:  "hyd",
			State: "Te",
		},
	}

	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("City:", p.City)
	fmt.Println("State:", p.State)
}
