package slices

import (
	"fmt"
	"math/rand"
)

func Testswitch() {
infloop:
	for {
		switch i := rand.Intn(20); {
		case i%2 == 0:
			fmt.Println(i)
			fallthrough
		case i%2 != 0:
			fmt.Println(i)
			break infloop
		}

	}
}

func DemoSlice() {
	planets := [...]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	//fmt.Println("Planets in our Solar system:", planets)
	fmt.Println("Total:", len(planets))
	for _, v := range planets {
		//fmt.Printf("%v. %v \n", i+1, v)
		fmt.Printf("%v \n", v)
	}

	slicedp := planets[1:4]

	fmt.Println("Sliced Length:", len(slicedp))
	fmt.Println("Sliced Capacity:", cap(slicedp))
	fmt.Println("Sliced Array:", slicedp)

	slicecopy := make([]string, len(slicedp))

	copy(slicecopy, slicedp)

	slicecopy = append(slicecopy, "Pluto")

	fmt.Println("Copy Length:", len(slicecopy))
	fmt.Println("Copy Capacity:", cap(slicecopy))
	fmt.Println("Copy Array:", slicecopy)

	slicecopy = append(slicecopy[:1], slicecopy[2:]...)

	fmt.Println("removed Slice:", slicecopy)
}
