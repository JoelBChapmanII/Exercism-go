// Package weather provides tools to report weather conditions.
package weather

// CurrentCondition details the current weather conditions.
var CurrentCondition string

// CurrentLocation details the location in question.
var CurrentLocation string

/*
Forecast returns a string detailing the location given and its current weather conditions
it takes in a city and weather conditions for the given city as strings.
*/
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
