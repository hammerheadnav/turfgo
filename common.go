package turfgo

import "github.com/hammerheadnav/turfgo/turfgoMath"

func isEqualLocation(point1 *Point, point2 *Point) bool {
	return turfgoMath.IsEqualFloatPair(point1.Lat, point1.Lng, point2.Lat, point2.Lng, turfgoMath.TwelveDecimalPlaces)
}
