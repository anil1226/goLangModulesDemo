package squareroot

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("Square root module 1st init")

}

func init() {
	fmt.Println("Square root module 2nd init")

}

func Calculate(i int) int {
	return int(math.Sqrt(float64(i)))
}
