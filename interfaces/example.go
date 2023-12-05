package interfaces

import "fmt"

type vowelFinder interface {
	FindVowels() []rune
}

type myString string

func (s myString) FindVowels() []rune {
	var vowels []rune
	for _, v := range s {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			vowels = append(vowels, v)
		}
	}
	return vowels
}

func Vowels() {
	var v vowelFinder = myString("anil karuturi")

	fmt.Printf("%c", v.FindVowels())
}
