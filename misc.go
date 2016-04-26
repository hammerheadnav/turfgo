package turfgo

import (
	"errors"
	"math"

	"github.com/hammerheadnav/turfgo/turfgoMath"
)

const invalidBearing = -1234.0

// PointOnLine takes a Point and a LineString and calculates the closest Point on the LineString.
func PointOnLine(point *Point, lineString *LineString, units string) (*Point, float64, int, error) {
	closestPt := &Point{infinity, infinity}
	closestDistance := float64(infinity)
	index := -1

	coords := lineString.Points
	for i := 0; i < len(coords)-1; i++ {
		start := coords[i]
		stop := coords[i+1]
		startDist, err := Distance(point, start, units)
		if err != nil {
			return nil, -1, -1, err
		}
		stopDist, _ := Distance(point, stop, units)
		direction := Bearing(start, stop)
		height := math.Max(stopDist, startDist)

		perpendicularPt1, _ := Destination(point, height, direction+90, units)
		perpendicularPt2, _ := Destination(point, height, direction-90, units)
		intersect := lineIntersects(perpendicularPt1, perpendicularPt2, start, stop)
		intersectD := float64(infinity)
		if intersect != nil {
			intersectD, _ = Distance(point, intersect, units)
		}
		if startDist < closestDistance {
			closestPt = start
			closestDistance = startDist
			index = i
		}
		if stopDist < closestDistance {
			closestPt = stop
			closestDistance = stopDist
			index = i
		}
		if intersectD < closestDistance {
			closestPt = intersect
			closestDistance = intersectD
			index = i
		}
	}
	return closestPt, closestDistance, index, nil
}

// TriangularProjection calculate the projection of given point on the lineString, base angles for projection should be acute.
// If bearing should also be considerd, pass in a previous point also, otherwise it should be nil
func TriangularProjection(point *Point, previousPoint *Point, lineString *LineString, unit string) (*Point, float64, int, error) {
	bearing := invalidBearing
	if previousPoint != nil {
		bearing = Bearing(previousPoint, point)
	}
	for i := 0; i < len(lineString.Points)-1; i++ {
		start := lineString.Points[i]
		end := lineString.Points[i+1]
		if !isAnyBaseAngleObtuse(point, start, end) {
			bearingLs := Bearing(start, end)
			bearingDiff := bearing - bearingLs
			if bearing != invalidBearing && (bearingDiff < -45 || bearingDiff > +45) {
				continue
			}
			projection, distance, _, err := PointOnLine(point, NewLineString([]*Point{start, end}), unit)
			if err != nil {
				return nil, -1, -1, err
			}
			return projection, distance, i, nil
		}
	}
	return nil, -1, -1, errors.New("No Projection found")
}

func lineIntersects(line1Start *Point, line1End *Point, line2Start *Point, line2End *Point) *Point {
	// if the lines intersect, the result contains the x and y of the intersection (treating the lines as infinite) and booleans for whether line segment 1 or line segment 2 contain the point
	// denominator = ((line2EndY - line2StartY) * (line1EndX - line1StartX)) - ((line2EndX - line2StartX) * (line1EndY - line1StartY));
	denominator := ((line2End.Lng - line2Start.Lng) * (line1End.Lat - line1Start.Lat)) - ((line2End.Lat - line2Start.Lat) * (line1End.Lng - line1Start.Lng))
	if denominator == 0 {
		return nil
	}
	a := line1Start.Lng - line2Start.Lng
	b := line1Start.Lat - line2Start.Lat
	numerator1 := ((line2End.Lat - line2Start.Lat) * a) - ((line2End.Lng - line2Start.Lng) * b)
	numerator2 := ((line1End.Lat - line1Start.Lat) * a) - ((line1End.Lng - line1Start.Lng) * b)
	a = numerator1 / denominator
	b = numerator2 / denominator
	// if we cast these lines infinitely in both directions, they intersect here:
	lat := line1Start.Lat + (a * (line1End.Lat - line1Start.Lat))
	lng := line1Start.Lng + (a * (line1End.Lng - line1Start.Lng))
	onLine1 := false
	onLine2 := false

	// if line1 is a segment and line2 is infinite, they intersect if:
	if a > 0 && a < 1 {
		onLine1 = true
	}

	// if line2 is a segment and line1 is infinite, they intersect if:
	if b > 0 && b < 1 {
		onLine2 = true
	}

	// if line1 and line2 are segments, they intersect if both of the above are true
	if onLine1 && onLine2 {
		return &Point{lat, lng}
	}
	return nil
}

func isAnyBaseAngleObtuse(point *Point, start *Point, end *Point) bool {
	alpha, _ := Distance(point, start, "mi")
	beta, _ := Distance(point, end, "mi")
	gamma, _ := Distance(start, end, "mi")
	if gamma == 0 {
		return true
	}
	if turfgoMath.IsEqualFloat(alpha+beta, gamma, turfgoMath.TwelveDecimalPlaces) {
		return false
	}

	cosineA := ((alpha * alpha) + (gamma * gamma) - (beta * beta)) / (2 * alpha * gamma)
	cosineB := ((beta * beta) + (gamma * gamma) - (alpha * alpha)) / (2 * beta * gamma)
	if cosineA < 0 || cosineB < 0 {
		return true
	}
	return false
}
