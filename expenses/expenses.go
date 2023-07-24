package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
  var satisfy []Record
	for _, record := range in {
		if predicate(record) {
			satisfy= append(satisfy, record)
		}
	}
	return satisfy
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	filterFunc := func (r Record) bool {
		filter := p.From <= r.Day && r.Day <= p.To
		return filter
	}
	return filterFunc
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	filterFunc := func (r Record) bool{
		return r.Category == c
	}
	return filterFunc
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	filter, count := ByDaysPeriod(p), 0.0
	for _, record := range in {
		if filter(record) {
			count += record.Amount
		}
	}
	return count
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	expensesFromCat := Filter(in, ByCategory(c))
	if len(expensesFromCat) == 0 {
		return 0, errors.New("unknown category")
	}

	return TotalByPeriod(expensesFromCat, p), nil
}
