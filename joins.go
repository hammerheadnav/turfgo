package turfgo

// Inside takes a Point and a Polygon or MultiPolygon and determines if the point resides
// inside the polygon. The polygon can be convex or concave. The function accounts for holes.
func Inside(point *Point, polygon PolygonI) bool {
	polygons := polygon.getPolygons()
	insidePoly := false
	for i := 0; i < len(polygons) && !insidePoly; i++ {
		// check if it is in the outer ring first
		if inRing(point, polygons[i].LineStrings[0]) {
			inHole := false
			// check for the point in any of the holes
			for k := 1; k < len(polygons[i].LineStrings) && !inHole; k++ {
				if inRing(point, polygons[i].LineStrings[k]) {
					inHole = true
				}
			}
			if !inHole {
				insidePoly = true
			}
		}
	}
	return insidePoly
}

// Within takes a set of points and a set of polygons and returns the points that fall within the polygons.
func Within(points []*Point, polygons []PolygonI) []*Point {
	result := []*Point{}
	for _, polygon := range polygons {
		for _, point := range points {
			if Inside(point, polygon) {
				result = append(result, point)
			}
		}
	}
	return result
}

func inRing(point *Point, ring *LineString) bool {
	isInside := false
	ringPoints := ring.getPoints()
	for i, j := 0, len(ringPoints)-1; i < len(ringPoints); j, i = i, i+1 {
		xi, yi := ringPoints[i].Lng, ringPoints[i].Lat
		xj, yj := ringPoints[j].Lng, ringPoints[j].Lat
		intersect := ((yi > point.Lat) != (yj > point.Lat)) &&
			(point.Lng < (xj-xi)*(point.Lat-yi)/(yj-yi)+xi)
		if intersect {
			isInside = !isInside
		}
	}
	return isInside
}
