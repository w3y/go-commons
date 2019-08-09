package go_commons

import "math"

func RoundFloat(x, unit float64) float64 {
	return math.Ceil(x/unit) * unit
}