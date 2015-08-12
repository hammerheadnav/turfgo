package turfgo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInside(t *testing.T) {
	Convey("Given a simple polygon", t, func() {
		point1 := &Point{0, 0}
		point2 := &Point{0, 100}
		point3 := &Point{100, 100}
		point4 := &Point{100, 0}
		point5 := &Point{0, 0}
		lineString := NewLineString([]*Point{point1, point2, point3, point4, point5})
		polygon := NewPolygon([]*LineString{lineString})
		Convey("Should return true if point fall in polygon", func() {
			ptIn := NewPoint(50, 50)
			So(Inside(ptIn, polygon), ShouldBeTrue)
		})

		Convey("Should return false if point does not fall in polygon", func() {
			ptOut := NewPoint(140, 150)
			So(Inside(ptOut, polygon), ShouldBeFalse)
		})
	})

	Convey("Given a concave polygon", t, func() {
		point1 := &Point{0, 0}
		point2 := &Point{50, 50}
		point3 := &Point{0, 100}
		point4 := &Point{100, 100}
		point5 := &Point{100, 0}
		point6 := &Point{0, 0}
		lineString := NewLineString([]*Point{point1, point2, point3, point4, point5, point6})
		polygon := NewPolygon([]*LineString{lineString})
		Convey("Should return true if point fall in polygon", func() {
			ptIn := NewPoint(75, 75)
			So(Inside(ptIn, polygon), ShouldBeTrue)
		})

		Convey("Should return false if point does not fall in polygon", func() {
			ptOut := NewPoint(25, 50)
			So(Inside(ptOut, polygon), ShouldBeFalse)
		})
	})

	Convey("Given a polygon with hole", t, func() {
		point1 := &Point{36.23084281427824, -86.70478820800781}
		point2 := &Point{36.21062368007896, -86.73980712890625}
		point3 := &Point{36.173495506147, -86.71371459960938}
		point4 := &Point{36.17709826419592, -86.67526245117186}
		point5 := &Point{36.20910010895552, -86.67303085327148}
		point6 := &Point{36.230427405208005, -86.68041229248047}
		point7 := &Point{36.23084281427824, -86.70478820800781}
		lineStringOuterRing := NewLineString([]*Point{point1, point2, point3, point4, point5, point6, point7})

		point8 := &Point{36.217271643303604, -86.6934585571289}
		point9 := &Point{36.20771501855801, -86.71268463134766}
		point10 := &Point{36.19067640168397, -86.70238494873047}
		point11 := &Point{36.19691047217554, -86.68487548828125}
		point12 := &Point{36.20993115142727, -86.68264389038086}
		point13 := &Point{36.217271643303604, -86.6934585571289}
		lineStringInnerRing := NewLineString([]*Point{point8, point9, point10, point11, point12, point13})

		polygon := NewPolygon([]*LineString{lineStringOuterRing, lineStringInnerRing})
		Convey("Should return false if point fall in hole", func() {
			ptInHole := NewPoint(36.20373274711739, -86.69208526611328)
			So(Inside(ptInHole, polygon), ShouldBeFalse)
		})

		Convey("Should return true if point fall in polygon", func() {
			ptInPoly := NewPoint(36.20258997094334, -86.72229766845702)
			So(Inside(ptInPoly, polygon), ShouldBeTrue)
		})

		Convey("Should return false if point fall outside polygon", func() {
			ptOutPoly := NewPoint(36.18527313913089, -86.75079345703125)
			So(Inside(ptOutPoly, polygon), ShouldBeFalse)
		})
	})

	Convey("Given a multiPolygon with hole", t, func() {
		point1 := &Point{36.23084281427824, -86.70478820800781}
		point2 := &Point{36.21062368007896, -86.73980712890625}
		point3 := &Point{36.173495506147, -86.71371459960938}
		point4 := &Point{36.17709826419592, -86.67526245117186}
		point5 := &Point{36.20910010895552, -86.67303085327148}
		point6 := &Point{36.230427405208005, -86.68041229248047}
		point7 := &Point{36.23084281427824, -86.70478820800781}
		lineStringOuterRing := NewLineString([]*Point{point1, point2, point3, point4, point5, point6, point7})

		point8 := &Point{36.217271643303604, -86.6934585571289}
		point9 := &Point{36.20771501855801, -86.71268463134766}
		point10 := &Point{36.19067640168397, -86.70238494873047}
		point11 := &Point{36.19691047217554, -86.68487548828125}
		point12 := &Point{36.20993115142727, -86.68264389038086}
		point13 := &Point{36.217271643303604, -86.6934585571289}
		lineStringInnerRing := NewLineString([]*Point{point8, point9, point10, point11, point12, point13})
		polygon1 := NewPolygon([]*LineString{lineStringOuterRing, lineStringInnerRing})

		point14 := &Point{36.171278341935434, -86.76624298095703}
		point15 := &Point{36.2014818084173, -86.77362442016602}
		point16 := &Point{36.19607929145354, -86.74100875854492}
		point17 := &Point{36.170862616662134, -86.74238204956055}
		point18 := &Point{36.171278341935434, -86.76624298095703}
		lineString := NewLineString([]*Point{point14, point15, point16, point17, point18})
		polygon2 := NewPolygon([]*LineString{lineString})

		multiPolygon := NewMultiPolygon([]*Polygon{polygon1, polygon2})

		Convey("Should return false if point fall in hole", func() {
			ptInHole := NewPoint(36.20373274711739, -86.69208526611328)
			So(Inside(ptInHole, multiPolygon), ShouldBeFalse)
		})

		Convey("Should return true if point fall in polygon", func() {
			ptInPoly := NewPoint(36.20258997094334, -86.72229766845702)
			ptInPoly2 := NewPoint(36.18527313913089, -86.75079345703125)
			So(Inside(ptInPoly, multiPolygon), ShouldBeTrue)
			So(Inside(ptInPoly2, multiPolygon), ShouldBeTrue)
		})

		Convey("Should return false if point fall outside polygon", func() {
			ptOutPoly := NewPoint(36.23015046460186, -86.75302505493164)
			So(Inside(ptOutPoly, multiPolygon), ShouldBeFalse)
		})
	})
}

func TestWithin(t *testing.T) {
	Convey("Given a point and a polygon", t, func() {
		Convey("Should return points that fall in polygon", func() {
			point1 := &Point{0, 0}
			point2 := &Point{0, 100}
			point3 := &Point{100, 100}
			point4 := &Point{100, 0}
			point5 := &Point{0, 0}
			lineString := NewLineString([]*Point{point1, point2, point3, point4, point5})
			polygon := NewPolygon([]*LineString{lineString})
			pt := NewPoint(50, 50)
			points := []*Point{pt}
			result := Within(points, []PolygonI{polygon})
			So(result, ShouldResemble, points)
		})
	})

	Convey("Given multiple points and multiple polygons", t, func() {
		Convey("Should return points that fall in polygons", func() {
			lineString1 := NewLineString([]*Point{&Point{0, 0}, &Point{10, 0}, &Point{10, 10}, &Point{0, 10}, &Point{0, 0}})
			lineString2 := NewLineString([]*Point{&Point{10, 0}, &Point{20, 10}, &Point{20, 20}, &Point{20, 0}, &Point{10, 0}})
			polygon1 := NewPolygon([]*LineString{lineString1})
			polygon2 := NewPolygon([]*LineString{lineString2})
			point1 := NewPoint(1, 1)
			point2 := NewPoint(1, 3)
			point3 := NewPoint(14, 2)
			point4 := NewPoint(13, 1)
			point5 := NewPoint(19, 7)
			point6 := NewPoint(100, 7)
			points := []*Point{point1, point2, point3, point4, point5, point6}
			result := Within(points, []PolygonI{polygon1, polygon2})
			So(result, ShouldResemble, []*Point{point1, point2, point3, point4, point5})
		})
	})
}
