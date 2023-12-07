package goroutines

import "fmt"

func Producer(d chan int) {
	for i := 0; i <= 10; i++ {
		d <- i
	}
	close(d)
}

func Third() {
	c := make(chan int)

	go Producer(c)

	for v := range c {

		fmt.Println(v)
	}
}
