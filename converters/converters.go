package converters

func InchesToFeet (inches float64) float64 {
	return inches / 12.0
}

func FeetToInches (feet float64) float64 {
	return feet * 12.0
}

func MetersToYards (meters float64) float64 {
	return meters * 1.0936133
}

func YardsToMeters (yards float64) float64 {
	return yards * 0.9144
}
