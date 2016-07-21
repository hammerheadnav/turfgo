package turfgo

import (
	"math"

	"github.com/hammerheadnav/turfgo/turfgoMath"
)

func isEqualLocation(point1 *Point, point2 *Point) bool {
	return turfgoMath.IsEqualFloatPair(point1.Lat, point1.Lng, point2.Lat, point2.Lng, turfgoMath.TwelveDecimalPlaces)
}

func translate(point *Point, horizontalDisplacement float64, verticalDisplacement float64) *Point {
	latDisplacementRad := verticalDisplacement / R["m"]
	longDisplacementRad := horizontalDisplacement / (R["m"] * math.Cos(turfgoMath.DegreeToRad(point.Lat)))

	latDisplacement := turfgoMath.RadToDegree(latDisplacementRad)
	longDisplacement := turfgoMath.RadToDegree(longDisplacementRad)

	translatedPoint := NewPoint(point.Lat+latDisplacement, point.Lng+longDisplacement)

	return translatedPoint
}
