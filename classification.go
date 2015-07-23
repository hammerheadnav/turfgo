package turfgo

//Nearest takes a reference point and a set of points and returns the point from the set closest to the reference.
func Nearest(reference *Point, points []*Point) *Point {
	if len(points) == 0 {
		return nil
	}
	nearestPoint := points[0]
	distance, _ := Distance(reference, points[0], "mi")
	for i := 1; i < len(points); i++ {
		dist, _ := Distance(reference, points[i], "mi")
		if dist < distance {
			nearestPoint = points[i]
			distance = dist
		}
	}
	return nearestPoint
}
