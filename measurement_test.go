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
		point2 := &Point{39.123, -75.534}

		_, err := Distance(point1, point2, "invalidUnit")
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, fmt.Sprintf(unitError, "invalidUnit"))
	})

	Convey("Should return correct distance", t, func() {
		point1 := &Point{39.984, -75.343}
		point2 := &Point{39.123, -75.534}

		distMi, errM := Distance(point1, point2, "mi")
		So(errM, ShouldBeNil)
		So(distMi, ShouldEqual, 60.37218405837491)

		distKm, errK := Distance(point1, point2, "km")
		So(errK, ShouldBeNil)
		So(distKm, ShouldEqual, 97.15957803131901)

		distR, errR := Distance(point1, point2, "r")
		So(errR, ShouldBeNil)
		So(distR, ShouldEqual, 0.015245501024842149)

		distD, errD := Distance(point1, point2, "d")
		So(errD, ShouldBeNil)
		So(distD, ShouldEqual, 0.8735028650863799)
	})
}

func TestAlong(t *testing.T) {

	Convey("Given a wrong unit, should throw error", t, func() {
		point1 := &Point{39.984, -75.343}
		point2 := &Point{39.97074218352032, -75.4590397138299}
		lineString := NewLineString([]*Point{point1, point2})

		_, err := Along(lineString, 13, "invalidUnit")
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
		lineString := NewLineString([]*Point{point1, point2, point3, point4, point5, point6})
		expected := &Point{38.885335546214506, -77.02417351582903}

		p, err := Along(lineString, 1, "mi")
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
		lineString := NewLineString([]*Point{point1, point2, point3, point4, point5, point6})

		p, err := Along(lineString, 3, "mi")
		So(err, ShouldBeNil)
		So(p, ShouldResemble, point6)
	})
}

func TestExtent(t *testing.T) {

	Convey("Given different type of shapes, should return bounding box", t, func() {
		point := NewPoint(0.5, 102.0)

		lineStringPoints := []*Point{
			&Point{-10.0, 102.0},
			&Point{1.0, 103.0},
			&Point{0.0, 104.0},
			&Point{4.0, 130.0},
		}
		lineString := NewLineString(lineStringPoints)

		polygonOuterRing := NewLineString([]*Point{
			&Point{0.0, 20.0},
			&Point{0.0, 101.0},
		})

		polygonInnerRing := NewLineString([]*Point{
			&Point{1.0, 101.0},
			&Point{1.0, 100.0},
			&Point{0.0, 100.0},
		})

		polygonLinestrings := []*LineString{polygonOuterRing, polygonInnerRing}
		polygon := NewPolygon(polygonLinestrings)

		bBox := Extent(point, lineString, polygon)
		So(bBox[0], ShouldEqual, 20)
		So(bBox[1], ShouldEqual, -10)
		So(bBox[2], ShouldEqual, 130)
		So(bBox[3], ShouldEqual, 4)
	})
}

func TestCenter(t *testing.T) {

	Convey("Given an array of points, should return absolute center of points", t, func() {
		point1 := &Point{35.4691, -97.522259}
		point2 := &Point{35.463455, -97.502754}
		point3 := &Point{35.463245, -97.508269}
		point4 := &Point{35.465779, -97.516809}
		point5 := &Point{35.467072, -97.515372}
		lineString := NewLineString([]*Point{point1, point2, point3, point4})

		point := Center(lineString, point5)
		So(point.Lat, ShouldEqual, 35.4661725)
		So(point.Lng, ShouldEqual, -97.5125065)
	})
}

func TestOverlap(t *testing.T) {
	Convey("Should return false if boxes doesn't overlap", t, func() {
		// b2 above b1
		b1 := []float64{-0.2197265625, 19.31114335506464, 13.447265624999998, 28.304380682962783}
		b2 := []float64{3.4716796874999996, 32.24997445586331, 8.876953125, 35.88905007936091}
		b, err := Overlap(b1, b2)
		So(err, ShouldBeNil)
		So(b, ShouldBeFalse)

		// b2 on left of b1
		b1 = []float64{-0.2197265625, 19.31114335506464, 13.447265624999998, 28.304380682962783}
		b2 = []float64{-12.2197265625, 28.24997445586331, -2.876953125, 39.88905007936091}
		b, err = Overlap(b1, b2)
		So(err, ShouldBeNil)
		So(b, ShouldBeFalse)

		// b2 below b1
		b1 = []float64{-0.2197265625, 19.31114335506464, 13.447265624999998, 28.304380682962783}
		b2 = []float64{-12.2197265625, 2.24997445586331, 23.876953125, 15.88905007936091}
		b, err = Overlap(b1, b2)
		So(err, ShouldBeNil)
		So(b, ShouldBeFalse)

		// b2 on right of b1
		b1 = []float64{-0.2197265625, 19.31114335506464, 13.447265624999998, 28.304380682962783}
		b2 = []float64{15.2197265625, 18.24997445586331, 23.876953125, 29.88905007936091}
		b, err = Overlap(b1, b2)
		So(err, ShouldBeNil)
		So(b, ShouldBeFalse)
	})

	Convey("Should return true if boxes overlap", t, func() {
		// overlap where a point is inside either of the bbox
		b1 := []float64{-0.2197265625, 19.31114335506464, 13.447265624999998, 28.304380682962783}
		b2 := []float64{-2.4716796874999996, 15.24997445586331, 8.876953125, 21.88905007936091}
		b, err := Overlap(b1, b2)
		So(err, ShouldBeNil)
		So(b, ShouldBeTrue)

		// overlap where no point of either bbox reside inside any bbox
		b1 = []float64{-0.2197265625, 19.31114335506464, 13.447265624999998, 28.304380682962783}
		b2 = []float64{-4.2197265625, 21.24997445586331, 26.876953125, 24.88905007936091}
		b, err = Overlap(b1, b2)
		So(err, ShouldBeNil)
		So(b, ShouldBeTrue)
	})
}

func TestSurround(t *testing.T) {

	Convey("Given a point and bbox width, should return a bbox with given width and the point as its center", t, func() {
		point := &Point{13.04464000, 80.26688000}
		width := 500.0

		bBox := Surround(point, width)

		So(bBox[0], ShouldEqual, 80.2622657295878)
		So(bBox[1], ShouldEqual, 13.040144803113675)
		So(bBox[2], ShouldEqual, 80.2714942704122)
		So(bBox[3], ShouldEqual, 13.049135196886324)
	})

	Convey("Given a point and bbox width as zero, should return the same point as bbox", t, func() {
		point := &Point{35.4691, -97.522259}
		width := 0.0

		bBox := Surround(point, width)

		So(bBox[0], ShouldEqual, -97.522259)
		So(bBox[1], ShouldEqual, 35.4691)
		So(bBox[2], ShouldEqual, -97.522259)
		So(bBox[3], ShouldEqual, 35.4691)
	})
}
