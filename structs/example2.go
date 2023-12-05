package structs

import "fmt"

type Laptops struct {
	maker string
	model string
	Config
}

type Config struct {
	Processor string
	RAM       string
	HDD       string
}

func View() {
	laptop1 := Laptops{
		maker: "ASUS",
		model: "Zenbook",
		Config: Config{
			Processor: "Intel",
			RAM:       "16 GB",
			HDD:       "1 TB",
		},
	}
	laptop2 := &Laptops{
		maker: "Dell",
		model: "XPS",
	}

	fmt.Println(laptop1.Processor, (*laptop2).model)

	if laptop1 == *laptop2 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}
}
