// Package weather contains related functions and variables specific to weather report.
package weather

// CurrentCondition represents the current weather status.
var CurrentCondition string

// CurrentLocation represents the location for which the weather forecast is being read.
var CurrentLocation string

// Forecast retruns a human readable string representing the weather forecast for a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
