package strings_test

import (
	"goModDemo/strings"
	"testing"
)

type TestCase struct {
	name     string
	value    string
	expected bool
	actual   bool
}

func TestIsPalindrome2(t *testing.T) {

	t.Run("testing for +ve and -ve testcases of palindrome", func(t *testing.T) {
		tc := []TestCase{
			{value: "madam", expected: true, name: "testing madam"},
			{value: "mister", expected: false, name: "testing mister"},
		}

		for _, test := range tc {
			t.Run(test.name, func(t *testing.T) {
				test.actual = strings.IsPalindrome2(test.value)
				if test.actual != test.expected {
					t.Fail()
				}
			})
		}

	})

}
