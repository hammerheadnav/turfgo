package turfgo

import (
	"fmt"
	"io/ioutil"
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
		_, _, _, err := PointOnLine(pointOutsideLine, lineString, "invalidUnit")
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, fmt.Sprintf(unitError, "invalidUnit"))
	})

	Convey("Given a point and a lineString, should calculate a point on line", t, func() {
		expected := &Point{38.88079693955019, -77.03031748983136}
		exptectedDistance := 0.4263205659547497
		result, distance, index, _ := PointOnLine(pointOutsideLine, lineString, "mi")
		So(result, ShouldResemble, expected)
		So(distance, ShouldEqual, exptectedDistance)
		So(index, ShouldEqual, 0)
	})

	Convey("Other tests copied from turfjs", t, func() {
		gj1, _ := ioutil.ReadFile("./testFiles/pointOnLine/line1.geojson")
		ls1, _ := DecodeLineStringFromFeatureJSON(gj1)
		p1 := &Point{22.254624939561698, -97.79617309570312}
		expected1 := &Point{22.245731173797786, -97.83538404669022}
		exptectedDistance1 := 2.5824950936100484
		result1, distance1, index1, _ := PointOnLine(p1, ls1, "mi")
		So(result1, ShouldResemble, expected1)
		So(distance1, ShouldAlmostEqual, exptectedDistance1)
		So(index1, ShouldEqual, 0)

		gj2, _ := ioutil.ReadFile("./testFiles/pointOnLine/route1.geojson")
		ls2, _ := DecodeLineStringFromFeatureJSON(gj2)
		p2 := &Point{37.60117623656667, -79.0850830078125}
		expected2 := &Point{37.578608, -79.049412}
		exptectedDistance2 := 2.4998919202861694
		result2, distance2, index2, _ := PointOnLine(p2, ls2, "mi")
		So(result2, ShouldResemble, expected2)
		So(distance2, ShouldAlmostEqual, exptectedDistance2)
		So(index2, ShouldEqual, 1449)

		gj3, _ := ioutil.ReadFile("./testFiles/pointOnLine/route2.geojson")
		ls3, _ := DecodeLineStringFromFeatureJSON(gj3)
		p3 := &Point{45.96021963947196, -112.60660171508789}
		expected3 := &Point{45.970203, -112.614288}
		exptectedDistance3 := 0.7825944108810942
		result3, distance3, index3, _ := PointOnLine(p3, ls3, "mi")
		So(result3, ShouldResemble, expected3)
		So(distance3, ShouldAlmostEqual, exptectedDistance3)
		So(index3, ShouldEqual, 3759)
	})
}
