package goroutines

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello goroutine")
}

func First() {
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("first goroutine")
}
