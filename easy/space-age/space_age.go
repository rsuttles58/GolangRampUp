//Package space implements a function that returns an age in years based on seconds and a planet provided.
package space

//Planet is used to create a type for the 8 unique planets.
type Planet string

//Age calculates an age in years based on the number of seconds alive and the planet chosen.
func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case "Earth":
		return seconds / (31557600)
	case "Mercury":
		return seconds / (31557600 * .2408467)
	case "Venus":
		return seconds / (31557600 * .61519726)
	case "Mars":
		return seconds / (31557600 * 1.8808158)
	case "Jupiter":
		return seconds / (31557600 * 11.862615)
	case "Saturn":
		return seconds / (31557600 * 29.447498)
	case "Uranus":
		return seconds / (31557600 * 84.016846)
	case "Neptune":
		return seconds / (31557600 * 164.79132)
	default:
		return 0
	}

}
