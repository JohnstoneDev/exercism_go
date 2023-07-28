/*
package that returns date after one gigasecond (one thousand million seconds)
*/
package gigasecond

// import path for the time package from the standard library
import (
	"time"
)

// AddGigasecond adds a gigasecond (1,000,000,000 seconds) to the input time 't' and
// returns the resulting time.
func AddGigasecond(t time.Time) time.Time {
	// Calculate a gigasecond
	giga := time.Duration(1e9) * time.Second
	// Add the gigasecond to the input
	end := t.Add(giga)
	return end
}
