package maps

import "fmt"

type employee struct {
	salary   int
	location string
}

func Create() map[string]employee {
	// bikes := map[string]string{}

	// bike s["make"] = "Honda"
	// bikes["model"] = "Shine"

	// if value, ok := bikes["make"]; ok {
	// 	fmt.Println(value)
	// }

	// fmt.Println(bikes)

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

	//fmt.Println(employeemap)
	return employeemap
}
