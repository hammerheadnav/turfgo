package turfgo

// LineDiff take two lines and gives an array of lines by subracting second from first. Single coordinate overlaps are ignored.
// Line should not have duplicate values.
func LineDiff(firstLine *LineString, secondLine *LineString) []*LineString {
	diffSegments := []*LineString{}
	fPoints := firstLine.Points
	sPoints := secondLine.Points
	for i := 0; i < len(fPoints)-1; i++ {
		if !containLocationPair(sPoints, fPoints[i], fPoints[i+1]) {
			diffSegments = append(diffSegments, NewLineString([]*Point{fPoints[i], fPoints[i+1]}))
		}
	}
	return reduceDiffSegment(diffSegments)
}

// LineDiffPercentage take two lines and give the percentage of difference between first and second line with respect to first line.
// Single coordinate overlaps are ignored. Line should not have duplicate values.
func LineDiffPercentage(firstLine *LineString, secondLine *LineString) float64 {
	totalPoints := len(firstLine.Points)
	if totalPoints == 0 {
		return 0
	}

	diff := LineDiff(firstLine, secondLine)
	if len(diff) == 1 && totalPoints == len(diff[0].Points) {
		return 100
	}

	diffPoints := 0
	for _, line := range diff {
		isStartingSegment := isEqualLocation(firstLine.Points[0], line.Points[0])
		isEndingSegment := isEqualLocation(firstLine.Points[len(firstLine.Points)-1], line.Points[len(line.Points)-1])
		diffPoints += len(line.Points)

		if isStartingSegment || isEndingSegment {
			diffPoints--
		} else {
			diffPoints -= 2
		}
	}
	return (float64(diffPoints) / float64(totalPoints)) * 100
}

func reduceDiffSegment(segments []*LineString) []*LineString {
	if len(segments) == 0 {
		return segments
	}
	result := []*LineString{}
	previousSeg := segments[0]
	for i := 1; i < len(segments); i++ {
		currentSeg := segments[i]
		pLen := len(previousSeg.Points)
		previousSegLastPoint := previousSeg.Points[pLen-1]
		currentSegFirstPoint := currentSeg.Points[0]
		if isEqualLocation(previousSegLastPoint, currentSegFirstPoint) {
			mergedPoints := append(previousSeg.Points, currentSeg.Points[1])
			previousSeg = NewLineString(mergedPoints)
		} else {
			result = append(result, previousSeg)
			previousSeg = currentSeg
		}
	}
	result = append(result, previousSeg)
	return result
}

func containLocationPair(points []*Point, point1, point2 *Point) bool {
	for i := 0; i < len(points)-1; i++ {
		if isEqualLocation(point1, points[i]) && isEqualLocation(point2, points[i+1]) {
			return true
		}
	}
	return false
}
