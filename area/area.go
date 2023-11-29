package area

import "fmt"

func init() {
	fmt.Println("Area Module 1st init")
}

func init() {
	fmt.Println("Area Module 2nd init")
}

func Calculate(length float32, width float32) float32 {
	return length * width
}
