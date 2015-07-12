package turfgo

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBearing(t *testing.T) {

	Convey("Given two points, should calculate bearing between them", t, func() {
		point1 := &Point{39.984, -75.343}
		point2 := &Point{39.123, -75.534}
		expected1 := -170.2330491349224
		bearing1 := Bearing(point1, point2)
		So(bearing1, ShouldEqual, expected1)

		point3 := &Point{12.9715987, 77.59456269999998}
		point4 := &Point{13.22328378, 77.77448784}
		expected2 := 34.828578946361255
		bearing2 := Bearing(point3, point4)
		So(bearing2, ShouldEqual, expected2)
	})

}

func TestDestination(t *testing.T) {

	Convey("Given a wrong unit, should throw error", t, func() {
		startingPoint := &Point{39.984, -75.343}
		_, err := Destination(startingPoint, 32, 120, "invalidUnit")
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, fmt.Sprintf(unitError, "invalidUnit"))
	})

	Convey("Should return correct destination", t, func() {

		Convey("Given miles unit", func() {
			startingPoint := &Point{39.984, -75.343}
			expected := &Point{39.74662966576427, -75.81645928866797}

			dest, err := Destination(startingPoint, 30, -123, "mi")
			So(err, ShouldBeNil)
			So(dest, ShouldResemble, expected)
		})

		Convey("Given km unit", func() {
			startingPoint := &Point{39.984, -75.343}
			expected := &Point{40.01636403124376, -75.20865245149336}

			dest, err := Destination(startingPoint, 12, 72.5, "km")
			So(err, ShouldBeNil)
			So(dest, ShouldResemble, expected)
		})

		Convey("Given radian unit", func() {
			startingPoint := &Point{39.984, -75.343}
			expected := &Point{67.3178236932749, -216.61938960828266}

			dest, err := Destination(startingPoint, 1.2, 345, "r")
			So(err, ShouldBeNil)
			So(dest, ShouldResemble, expected)
		})

		Convey("Given degree unit", func() {
			startingPoint := &Point{39.984, -75.343}
			expected := &Point{-13.744745983973347, -92.31513759524121}

			dest, err := Destination(startingPoint, 56, 200, "d")
			So(err, ShouldBeNil)
			So(dest, ShouldResemble, expected)
		})

	})

}

func TestDistance(t *testing.T) {

	Convey("Given a wrong unit, should throw error", t, func() {
		point1 := &Point{39.984, -75.343}
		point2 := &Point{39.97074218352032, -75.4590397138299}

		_, err := Distance(point1, point2, "invalidUnit")
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, fmt.Sprintf(unitError, "invalidUnit"))
	})

	Convey("Should return correct distance", t, func() {
		point1 := &Point{39.984, -75.343}
		point2 := &Point{39.97074218352032, -75.4590397138299}
		expected := 9.999999999999373

		dist, err := Distance(point1, point2, "km")
		So(err, ShouldBeNil)
		So(dist, ShouldEqual, expected)
	})
}

func TestAlong(t *testing.T) {

	Convey("Given a wrong unit, should throw error", t, func() {
		point1 := &Point{39.984, -75.343}
		point2 := &Point{39.97074218352032, -75.4590397138299}
		points := []*Point{point1, point2}

		_, err := Along(points, 13, "invalidUnit")
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, fmt.Sprintf(unitError, "invalidUnit"))
	})

	Convey("Should return a point along distance", t, func() {
		point1 := &Point{38.878605, -77.031669}
		point2 := &Point{38.881946, -77.029609}
		point3 := &Point{38.884084, -77.020339}
		point4 := &Point{38.885821, -77.025661}
		point5 := &Point{38.889563, -77.021884}
		point6 := &Point{38.892368, -77.019824}
		points := []*Point{point1, point2, point3, point4, point5, point6}
		expected := &Point{38.885335546214506, -77.02417351582903}

		p, err := Along(points, 1, "mi")
		So(err, ShouldBeNil)
		So(p, ShouldResemble, expected)
	})

	Convey("Should return end point if distance longer then linestring", t, func() {
		point1 := &Point{38.878605, -77.031669}
		point2 := &Point{38.881946, -77.029609}
		point3 := &Point{38.884084, -77.020339}
		point4 := &Point{38.885821, -77.025661}
		point5 := &Point{38.889563, -77.021884}
		point6 := &Point{38.892368, -77.019824}
		points := []*Point{point1, point2, point3, point4, point5, point6}

		p, err := Along(points, 3, "mi")
		So(err, ShouldBeNil)
		So(p, ShouldResemble, point6)
	})
}
