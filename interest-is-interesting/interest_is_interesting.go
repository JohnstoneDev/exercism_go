package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var rate float32;

	if balance < 0 {
		rate = 3.213
	}

	if balance >= 0 && balance < 1000 {
		rate = 0.5
	}

	if balance >= 1000 && balance < 5000 {
		rate = 1.621
	}

	if balance >= 5000 {
		rate = 2.475
	}

	return rate
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return Interest(balance) + balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var count int

	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		count++
	}
	return count
}
