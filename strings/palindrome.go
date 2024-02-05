package strings

func IsPalindrome2(s string) bool {

	for i, v := range s {
		if string(v) != string(s[len(s)-1-i]) {
			return false
		}
	}
	return true

	//a := make(map[int]string)
	//a = append(a, "hello")
	//a[0] = "Hello"
	// a[1] = "World"

	// //fmt.Println(a[0], a[1]) //=> Hello World
	// fmt.Println(a) // => [Hello World]

	//return true

	// slice := []int{2, 3, 4}

	// fmt.Println(slice[0:2])

	// st := "hello"
	// i := 2

	// sr := fmt.Sprint(i) //string(i)

	// fmt.Println(sr, st)
}
