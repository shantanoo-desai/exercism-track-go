package space

type Planet string

const (
	earth_year_secs   = 31557600
	mercury_year_secs = earth_year_secs * 0.2408467
	venus_year_secs   = earth_year_secs * 0.61519726
	mars_year_secs    = earth_year_secs * 1.8808158
	jupiter_year_secs = earth_year_secs * 11.862615
	saturn_year_secs  = earth_year_secs * 29.447498
	uranus_year_secs  = earth_year_secs * 84.016846
	neptune_year_secs = earth_year_secs * 164.79132
)

// Age: Calculate how old you are on a planet
func Age(seconds float64, planet Planet) float64 {

	switch planet {
	case "Earth":
		return seconds / earth_year_secs
	case "Mercury":
		return seconds / mercury_year_secs
	case "Venus":
		return seconds / venus_year_secs
	case "Mars":
		return seconds / mars_year_secs
	case "Jupiter":
		return seconds / jupiter_year_secs
	case "Saturn":
		return seconds / saturn_year_secs
	case "Uranus":
		return seconds / uranus_year_secs
	case "Neptune":
		return seconds / neptune_year_secs
	default:
		return seconds // In Case no Planet provided send the same seconds again
	}
}
