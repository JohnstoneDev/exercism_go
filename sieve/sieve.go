package sieve

func isPrime(n int) bool {

	if n <= 3 {
		return true
	 }

	if n % 2 == 0 || n % 3 == 0 {
		 return false
	}

	i := 5
	for i * i <= n {
		if n % i == 0 || n % (i + 2) == 0 {
			return false
		}
		i += 6
	}

	return true
}


func Sieve(limit int) []int {
	var primes []int
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes
}
