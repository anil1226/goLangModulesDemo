package goroutines

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func ApiCall() {
	c1 := make(chan error)
	c2 := make(chan error)

	go first(c1)
	go second(c2)

	if err := <-c1; err != nil {
		log.Println(err.Error())
		return
	}
	if err := <-c2; err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Println("calling 3rd")
}

func first(c chan error) {
	res, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
		c <- err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err.Error())
		c <- err
	}
	fmt.Println(string(body))
	fmt.Println("first")
	c <- nil

}
func second(c chan error) {
	res, err := http.Get("http://pokeapi1234.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
		c <- err
		return
	}
	fmt.Println("jknervnjk")
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err.Error())
		c <- err
		return
	}
	fmt.Println(string(body))
	fmt.Println("second")
	time.Sleep(3 * time.Second)
	c <- nil
}
