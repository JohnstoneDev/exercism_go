package meteorology

import f "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type

func (input TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[input]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type

func (ti Temperature) String() string {
	return f.Sprintf("%v %v", ti.degree, ti.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit

func (input SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[input]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed

func (su Speed) String() string {
	return f.Sprintf("%v %v", su.magnitude, su.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData

func (input MeteorologyData) String() string {
	return f.Sprintf("%v: %v, Wind %s at %v, %d%% Humidity", input.location, input.temperature, input.windDirection, input.windSpeed, input.humidity)
}
