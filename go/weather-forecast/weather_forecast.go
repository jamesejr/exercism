// Package weather forcasts the current weather conditions of various cities in Goblinocus.
package weather

// CurrentCondition stores the current weather condition (e.g. sunny, rainy, etc.).
var CurrentCondition string

// CurrentLocation stores the name of a city in Goblinocus (e.g. San Goblinsco).
var CurrentLocation string

// Forecast provides the current weather condition for a city in Goblinocus.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
