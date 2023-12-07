package goroutines

import (
	"fmt"
)

func hello1(d chan<- bool) {
	fmt.Println("hello goroutine")
	d <- true
}

func Second() {
	c := make(chan bool)
	go hello1(c)
	<-c
	//time.Sleep(1 * time.Second)
	fmt.Println("first goroutine")
}
