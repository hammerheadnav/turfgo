package turfgo

import (
	"fmt"
	"math"

	"github.com/hammerheadnav/turfgo/turfgoMath"
)

var unitError = "%s is not a valid unit. Allowed units are mi(miles), km(kilometers), d(degrees) and r(radians)"

// Along takes a line and returns a point at a specified distance along the line.
func Along(lineString []*Point, distance float64, unit string) (*Point, error) {
	var travelled float64
	for i, point := range lineString {
		if distance >= travelled && i == len(lineString)-1 {
			return point, nil
		} else if travelled >= distance {
			overshot := distance - travelled
			if overshot == 0 {
				return point, nil
			}
			direction := Bearing(lineString[i], lineString[i-1]) - 180
			interpolated, err := Destination(lineString[i], overshot, direction, unit)
			if err != nil {
				return nil, err
			}
			return interpolated, nil

		} else {
			t, err := Distance(lineString[i], lineString[i+1], unit)
			if err != nil {
				return nil, err
			}
			travelled += t
		}
	}
	return nil, nil
}

// Bearing takes two points and finds the geographic bearing between them.
func Bearing(point1, point2 *Point) float64 {
	lat1 := turfgoMath.DegreeToRad(point1.Lat)
	lat2 := turfgoMath.DegreeToRad(point2.Lat)
	lon1 := turfgoMath.DegreeToRad(point1.Lng)
	lon2 := turfgoMath.DegreeToRad(point2.Lng)
	a := math.Sin(lon2-lon1) * math.Cos(lat2)
	b := math.Cos(lat1)*math.Sin(lat2) -
		math.Sin(lat1)*math.Cos(lat2)*math.Cos(lon2-lon1)
	return turfgoMath.RadToDegree(math.Atan2(a, b))
}

// Center takes an array of points and returns the absolute center point of all features.
func Center(points []*Point) *Point {
	bBox := Extent(points)
	lng := (bBox[0] + bBox[2]) / 2
	lat := (bBox[1] + bBox[3]) / 2
	return &Point{lat, lng}
}

// Destination takes a Point and calculates the location of a destination point
// given a distance in degrees, radians, miles, or kilometers; and bearing in
// degrees. This uses the Haversine formula to account for global curvature.
func Destination(startingPoint *Point, distance float64, bearing float64, unit string) (*Point, error) {
	radius, ok := R[unit]
	if !ok {
		return nil, fmt.Errorf(unitError, unit)
	}

	lat := turfgoMath.DegreeToRad(startingPoint.Lat)
	lon := turfgoMath.DegreeToRad(startingPoint.Lng)
	bearingRad := turfgoMath.DegreeToRad(bearing)

	destLat := math.Asin(math.Sin(lat)*math.Cos(distance/radius) +
		math.Cos(lat)*math.Sin(distance/radius)*math.Cos(bearingRad))
	destLon := lon + math.Atan2(math.Sin(bearingRad)*math.Sin(distance/radius)*math.Cos(lat),
		math.Cos(distance/radius)-math.Sin(lat)*math.Sin(destLat))

	return &Point{turfgoMath.RadToDegree(destLat), turfgoMath.RadToDegree(destLon)}, nil
}

// Distance calculates the distance between two points in degress, radians, miles, or
// kilometers. This uses the Haversine formula to account for global curvature.
func Distance(point1 *Point, point2 *Point, unit string) (float64, error) {
	radius, ok := R[unit]
	if !ok {
		return 0, fmt.Errorf(unitError, unit)
	}

	dLat := turfgoMath.DegreeToRad(point2.Lat - point1.Lat)
	dLon := turfgoMath.DegreeToRad(point2.Lng - point1.Lng)
	latRad1 := turfgoMath.DegreeToRad(point1.Lat)
	latRad2 := turfgoMath.DegreeToRad(point2.Lat)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(latRad1)*math.Cos(latRad2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return radius * c, nil
}

// Extent takes a set of points, calculates the extent of all input points, and returns a bounding box.
// Linestring and polygons can be represented as array of points.
// Returns []float64 the bounding box of input given as an array in WSEN order (west, south, east, north)
func Extent(shapes ...[]*Point) []float64 {
	extent := []float64{INFINITY, INFINITY, -INFINITY, -INFINITY}

	for _, points := range shapes {
		for _, point := range points {
			if extent[0] > point.Lng {
				extent[0] = point.Lng
			}
			if extent[1] > point.Lat {
				extent[1] = point.Lat
			}
			if extent[2] < point.Lng {
				extent[2] = point.Lng
			}
			if extent[3] < point.Lat {
				extent[3] = point.Lat
			}
		}
	}
	return extent
}
