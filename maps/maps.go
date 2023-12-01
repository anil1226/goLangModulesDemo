package maps

import "fmt"

func Create() {
	// bikes := map[string]string{}

	// bikes["make"] = "Honda"
	// bikes["model"] = "Shine"

	// if value, ok := bikes["make"]; ok {
	// 	fmt.Println(value)
	// }

	// fmt.Println(bikes)

	type employee struct {
		salary   int
		location string
	}

	emp1 := employee{
		salary:   10000,
		location: "India",
	}

	emp2 := employee{
		salary:   20000,
		location: "USA",
	}

	employeemap := map[string]employee{
		"Anil": emp1,
		"Raj":  emp2,
	}

	for key, value := range employeemap {
		fmt.Println(key, value.salary, value.location)
	}

	fmt.Println(employeemap)
}
