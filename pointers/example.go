package pointers

import "fmt"

func Demo() {
	a := 100
	//b := new(int)
	//*b = 200
	b := &a
	fmt.Println(a, b, *b)
	//*b++
	change(b)
	fmt.Println(a, b, *b)
	c := hello()
	fmt.Println(c, *c)

	d := []int{5, 20, 30}
	modify(&d)
	fmt.Println(d)
}

func change(val *int) {
	*val++
}

func hello() *int {
	i := 200
	return &i
}

func modify(arr *[]int) {
	(*arr)[0] = 10
	//arr[0] = 10
}
