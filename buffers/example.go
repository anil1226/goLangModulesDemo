package buffers

import "fmt"

func PrimeNumber(count int) string {
	primes := ""
	for i := 2; i <= count; i++ {
		isPrime := true

		for j := 2; j <= 9; j++ {
			if i != j {

				if i%j == 0 {
					isPrime = false
				}
			}

		}

		if isPrime {
			primes += fmt.Sprint(i) + ", "
		}
	}
	return primes

}

//2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97
