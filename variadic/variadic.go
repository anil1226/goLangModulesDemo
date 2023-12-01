package variadic

import "fmt"

func Find(num int, nums ...int) {
	fmt.Printf("Type of nums is %T \n", nums)
	found := false

	for i, v := range nums {
		if v == num {
			fmt.Println("Found the number ", num, "at index ", i)
			found = true
		}
	}

	if !found {
		fmt.Println(num, " number not found in ", nums)
	}

}
