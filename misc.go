package turfgo

// PointOnLine takes a Point and a LineString and calculates the closest Point on the LineString.
func PointOnLine(point *Point, lineString *LineString, units string) (*Point, float64, error) {
	closestPt := &Point{infinity, infinity}
	closestDistance := float64(infinity)

	coords := lineString.Points
	for i := 0; i < len(coords)-1; i++ {
		start := coords[i]
		stop := coords[i+1]
		startDist, err := Distance(point, start, units)
		if err != nil {
			return nil, -1, err
		}
		stopDist, _ := Distance(point, stop, units)
		direction := Bearing(start, stop)
		perpendicularPt, _ := Destination(point, 1000, direction+90, units) // 1000 = gross
		intersect := lineIntersects(point, perpendicularPt, start, stop)
		if intersect == nil {
			perpendicularPt, _ = Destination(point, 1000, direction+90, units)
			intersect = lineIntersects(point, perpendicularPt, start, stop)
		}
		intersectD := float64(infinity)
		if intersect != nil {
			intersectD, _ = Distance(point, intersect, units)
		}
		if startDist < closestDistance {
			closestPt = start
			closestDistance = startDist
		}
		if stopDist < closestDistance {
			closestPt = stop
			closestDistance = stopDist
		}
		if intersectD < closestDistance {
			closestPt = intersect
			closestDistance = intersectD
		}
	}
	return closestPt, closestDistance, nil
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
