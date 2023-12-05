package interfaces

import "fmt"

type Square struct {
	Length int
	Width  int
}

func Convert(i interface{}) {
	val, ok := i.(Square)
	if ok {
		fmt.Println(val)
	} else {
		fmt.Printf("in else :%T", i)
	}
}

func Exam3() {
	i := Square{
		Length: 1,
		Width:  2,
	}
	_ = i
	var j interface{} = "stringjife"
	Convert(j)
}
