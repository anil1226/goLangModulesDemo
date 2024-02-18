package strings

func IsPalindrome2(s string) bool {

	for i, v := range s {
		if string(v) != string(s[len(s)-1-i]) {
			return false
		}
	}
	return true

}
