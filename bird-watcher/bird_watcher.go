package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var totalBirds int
	for count := 0; count < len(birdsPerDay); count++ {
		totalBirds += birdsPerDay[count]
	}
	return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var start, end int
	var getWeek []int

	if week == 1 {
		start, end = 0, start + 7
	} else {
		end = (7 * week)
		start =  (end - 7) + 1
	}

	getWeek = birdsPerDay[start:end]

	return TotalBirdCount(getWeek)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}
	return birdsPerDay
}
