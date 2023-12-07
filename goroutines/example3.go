package goroutines

import (
	"fmt"
)

func EvenOrNot() {

	numbers := []int{1, 54, 73, 21, 59, 76, 62, 91}

	even := make(chan int)

	odd := make(chan int)

	for _, v := range numbers {

		go Find(even, odd, v)

	}

	for i := 0; i < len(numbers); i++ {
		select {

		case e := <-even:
			fmt.Println("even", e)

		case d := <-odd:
			fmt.Println("odd", d)

		}
	}

}

func Find(even chan int, odd chan int, v int) {
	//for _, v := range nums {
	//time.Sleep(5 * time.Second)
	if v%2 == 0 {
		even <- v
	} else {
		odd <- v
	}

	//}
	//close(even)
	//close(odd)
}
