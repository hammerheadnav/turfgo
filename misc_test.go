package turfgo

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPointOnLine(t *testing.T) {
	pointOutsideLine := &Point{38.884017, -77.037076}
	point1 := &Point{38.878605, -77.031669}
	point2 := &Point{38.881946, -77.029609}
	point3 := &Point{38.884084, -77.020339}
	point4 := &Point{38.885821, -77.025661}
	point5 := &Point{38.889563, -77.021884}
	point6 := &Point{38.892368, -77.019824}
	lineString := NewLineString([]*Point{point1, point2, point3, point4, point5, point6})

	Convey("Given a wrong unit, should throw error", t, func() {
		_, _, err := PointOnLine(pointOutsideLine, lineString, "invalidUnit")
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, fmt.Sprintf(unitError, "invalidUnit"))
	})

	Convey("Given a point and a lineString, should calculate a point on line", t, func() {
		expected := &Point{38.88079693955019, -77.03031748983136}
		exptectedDistance := 0.4263205659547497
		result, distance, _ := PointOnLine(pointOutsideLine, lineString, "mi")
		So(result, ShouldResemble, expected)
		So(distance, ShouldEqual, exptectedDistance)
	})

}
