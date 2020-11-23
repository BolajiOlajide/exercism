package space

import (
	"fmt"
	"strconv"
)

// Planet represents the different planet in the solar system
type Planet string

const earthYearsInSeconds float64 = 31557600

// Age calculate how much time it'd take in another planet
func Age(timeInSeconds float64, planet Planet) float64 {
	switch planet {
	case "Earth":
		return getAgeForPlanet(timeInSeconds, 1.0)
	case "Mercury":
		return getAgeForPlanet(timeInSeconds, 0.2408467)
	case "Venus":
		return getAgeForPlanet(timeInSeconds, 0.61519726)
	case "Mars":
		return getAgeForPlanet(timeInSeconds, 1.8808158)
	case "Jupiter":
		return getAgeForPlanet(timeInSeconds, 11.862615)
	case "Uranus":
		return getAgeForPlanet(timeInSeconds, 84.016846)
	case "Neptune":
		return getAgeForPlanet(timeInSeconds, 164.79132)
	case "Saturn":
		return getAgeForPlanet(timeInSeconds, 29.447498)
	default:
		return getAgeForPlanet(timeInSeconds, 1.0)
	}
}

func getAgeForPlanet(timeInSeconds, planetAgeInEarthYears float64) float64 {
	age := timeInSeconds / (earthYearsInSeconds * planetAgeInEarthYears)

	ageInStr := fmt.Sprintf("%.2f", (age*100)/100)
	ageInFloat, _ := strconv.ParseFloat(ageInStr, 64)

	return ageInFloat
}
