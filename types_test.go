package turfgo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetPoints(t *testing.T) {
	Convey("For a given point, should return points array", t, func() {
		p := &Point{114.175329, 22.2524}
		So(p.getPoints(), ShouldResemble, []*Point{p})
	})

	Convey("For a given lineString, should return points array", t, func() {
		point1 := &Point{35.4691, -97.522259}
		point2 := &Point{35.463455, -97.502754}
		point3 := &Point{35.463245, -97.508269}
		points := []*Point{point1, point2, point3}
		lineString := NewLineString(points)
		So(lineString.getPoints(), ShouldResemble, points)
	})

	Convey("For a given multilineString, should return points array", t, func() {
		point1 := &Point{35.4691, -97.522259}
		point2 := &Point{35.463455, -97.502754}
		point3 := &Point{35.463245, -97.508269}
		point4 := &Point{22.7, -72.5}
		points1 := []*Point{point1, point2}
		points2 := []*Point{point3, point4}
		lineString1 := NewLineString(points1)
		lineString2 := NewLineString(points2)
		multiLineString := NewMultiLineString([]*LineString{lineString1, lineString2})
		So(multiLineString.getPoints(), ShouldResemble, append(points1, points2...))
	})

	Convey("For a given polygon, should return points array", t, func() {
		point1 := &Point{35.4691, -97.522259}
		point2 := &Point{35.463455, -97.502754}
		point3 := &Point{35.463245, -97.508269}
		point4 := &Point{22.7, -72.5}
		points1 := []*Point{point1, point2}
		points2 := []*Point{point3, point4}
		points3 := []*Point{point1, point3, point4}
		lineString1 := NewLineString(points1)
		lineString2 := NewLineString(points2)
		lineString3 := NewLineString(points3)
		polygon1 := NewPolygon([]*LineString{lineString1, lineString2})
		polygon2 := NewPolygon([]*LineString{lineString3})
		multiPolygon := NewMultiPolygon([]*Polygon{polygon1, polygon2})
		result := append(points1, points2...)
		result = append(result, points3...)
		So(multiPolygon.getPoints(), ShouldResemble, result)
	})

}
