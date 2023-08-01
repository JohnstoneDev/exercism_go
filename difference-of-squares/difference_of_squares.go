package diffsquares

func SquareOfSum(n int) int {
	total := 0

	for i := 1; i <= n; i++ {
		total += i
	}

	total *= total

	return total
}

func SumOfSquares(n int) int {
	total := 0

	for i := 1; i <= n; i++ {
		total += (i * i)
	}

	return total
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
