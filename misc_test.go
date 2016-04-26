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
		expected := &Point{38.881361463229524, -77.02996941477018}
		exptectedDistance := 0.4241146325840119
		result, distance, index, _ := PointOnLine(pointOutsideLine, lineString, "mi")
		So(result, ShouldResemble, expected)
		So(distance, ShouldEqual, exptectedDistance)
		So(index, ShouldEqual, 0)
	})

	Convey("Other tests copied from turfjs", t, func() {
		gj1, _ := ioutil.ReadFile("./testFiles/pointOnLine/line1.geojson")
		ls1, _ := DecodeLineStringFromFeatureJSON(gj1)
		p1 := &Point{22.254624939561698, -97.79617309570312}
		expected1 := &Point{22.247393614241208, -97.83572934173806}
		exptectedDistance1 := 2.5792333253307405
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

func TestIsTriangleObtuse(t *testing.T) {
	Convey("Given three sides of triangle", t, func() {
		Convey("Should return true if triangle is obtuse", func() {
			result := isAnyBaseAngleObtuse(NewPoint(0, 0), NewPoint(3, 0), NewPoint(5, 3))
			So(result, ShouldBeTrue)
		})

		Convey("Should return false if triangle is not obtuse triangle", func() {
			result := isAnyBaseAngleObtuse(NewPoint(0, 0), NewPoint(3, 0), NewPoint(2, 3))
			So(result, ShouldBeFalse)
		})
	})
}

func TestTriangularProjection(t *testing.T) {
	Convey("Should find the projection without bearing if previous point is nil and current point fall in middle", t, func() {
		a := NewPoint(40.88969576429507, -74.02225255966187)
		b := NewPoint(40.890660929502566, -74.02167320251465)
		c := NewPoint(40.891504423701505, -74.02114748954773)
		p := NewPoint(40.89128544066409, -74.02228474617003)
		expectedProj := NewPoint(40.89099374267134, -74.02146577465254)
		lineString := NewLineString([]*Point{a, b, c})
		proj, dist, index, err := TriangularProjection(p, nil, lineString, "mi")
		So(err, ShouldBeNil)
		So(proj, ShouldResemble, expectedProj)
		So(dist, ShouldAlmostEqual, 0.047301111679236334)
		So(index, ShouldEqual, 1)
	})

	Convey("Should pass if previous point is given and bearing is within 45 degree of either side", t, func() {
		a := NewPoint(40.88969576429507, -74.02225255966187)
		b := NewPoint(40.890660929502566, -74.02167320251465)
		c := NewPoint(40.891504423701505, -74.02114748954773)
		p := NewPoint(40.89128544066409, -74.02228474617003)
		p1 := NewPoint(40.89057171314116, -74.02269244194031)
		expectedProj := NewPoint(40.89099374267134, -74.02146577465254)
		lineString := NewLineString([]*Point{a, b, c})
		proj, dist, index, err := TriangularProjection(p, p1, lineString, "mi")
		So(err, ShouldBeNil)
		So(proj, ShouldResemble, expectedProj)
		So(dist, ShouldAlmostEqual, 0.047301111679236334)
		So(index, ShouldEqual, 1)
	})

	Convey("Should pass if previous point is given and bearing is outside 45 degree of either side", t, func() {
		a := NewPoint(40.88969576429507, -74.02225255966187)
		b := NewPoint(40.890660929502566, -74.02167320251465)
		c := NewPoint(40.891504423701505, -74.02114748954773)
		p := NewPoint(40.89128544066409, -74.02228474617003)
		p1 := NewPoint(40.891155672596184, -74.02328252792358)
		lineString := NewLineString([]*Point{a, b, c})
		proj, dist, index, err := TriangularProjection(p, p1, lineString, "mi")
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "No Projection found")
		So(proj, ShouldBeNil)
		So(dist, ShouldEqual, -1)
		So(index, ShouldEqual, -1)
	})

	Convey("Should fail if no projection on line", t, func() {
		a := NewPoint(40.88969576429507, -74.02225255966187)
		b := NewPoint(40.890660929502566, -74.02167320251465)
		c := NewPoint(40.891504423701505, -74.02114748954773)
		p := NewPoint(40.892404679685235, -74.02269244194031)
		lineString := NewLineString([]*Point{a, b, c})
		proj, dist, index, err := TriangularProjection(p, nil, lineString, "mi")
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "No Projection found")
		So(proj, ShouldBeNil)
		So(dist, ShouldEqual, -1)
		So(index, ShouldEqual, -1)
	})

	Convey("Should fail if not enough points on linestring", t, func() {
		a := NewPoint(40.88969576429507, -74.02225255966187)
		p := NewPoint(40.892404679685235, -74.02269244194031)
		lineString := NewLineString([]*Point{a})
		proj, dist, index, err := TriangularProjection(p, nil, lineString, "mi")
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "No Projection found")
		So(proj, ShouldBeNil)
		So(dist, ShouldEqual, -1)
		So(index, ShouldEqual, -1)
	})
}
