package strings

import (
	"fmt"
)

func Run() {

	s := "Anil₹"

	fmt.Println(s)
	r := []rune(s)

	for i := 0; i < len(r); i++ {
		//fmt.Printf("%c", s[i])
		//fmt.Printf("%x", s[i])
		fmt.Println(i, " :", s[i])
	}

}

func Convert(str []rune) string {
	str[0] = 'a'
	return string(str)
}

func IsPalindrome(str string) bool {
	revstr := ""

	for i := len(str) - 1; i >= 0; i-- {
		revstr += string(str[i])
	}
	return revstr == str

}
