// Package weather that is used to describe a weather Forecast for a location.
package weather

 // CurrentCondition variable that stores the weather.
var CurrentCondition string

 // CurrentLocation variable that stores the location for a forecast.
var CurrentLocation string

// Forecast function takes city & condition parameters and returns a string showing the weather forecast for the location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
