package main


import (
	"fmt"
)

// function that checks whether a number is Prime
func isPrime(n int) bool {

	if n <= 3 {
		return true
	 }

	if n % 2 == 0 || n % 3 == 0 {
		 return false
	}

	/*
		The code uses a variant of the optimized prime-checking algorithm known as the 6k ± 1 optimization.
		This optimization takes advantage of the fact that all prime numbers greater than 6 can be written
		in the form of either 6k - 1 or 6k + 1, where k is a positive integer.
	*/

	i := 5
	for i * i <= n {
		if n % i == 0 || n % (i + 2) == 0 {
			return false
		}
		i += 6
	}

	return true
}

// returns a pair of prime numbers between m & n(inclusive)
// where m - n == g
func Gap(g, m, n int) []int {

	if g < 2 || m < 2 {
		return nil
	}

	var primes, satisfies, pairs []int

	// find prime numbers in the range m...n
	for i := m; i <= n; i ++ {
		primeTest := isPrime(i)

		if primeTest {
			primes = append(primes, i)
		}
	}

	// loop through the prime numbers & find those with the desired difference
	for i := 0; i < len(primes); i++ {
		for j := i + 1; j < len(primes); j++ {
			currentNum, nextNum := primes[i], primes[j]

			// look for numbers that satisfy n - m == g
			if (nextNum - currentNum) == g {
				satisfies = append(satisfies, currentNum, nextNum)
			}
		}
	}

	// return the first two
	pairs = append(pairs, satisfies[0], satisfies[1])

	return pairs
}


func main() {
	fmt.Println(Gap(4, 100, 110), "<= Expecting [103, 107]")
	fmt.Println(Gap(6,100,110), "<= Expecting [101, 107]")
	fmt.Println(Gap(2,300,400), "<= Expecting [311,313]")
}