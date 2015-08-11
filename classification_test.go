package turfgo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNearest(t *testing.T) {

	Convey("Given a reference point and a bunch of points", t, func() {

		Convey("Should return nil and 0 if no points", func() {
			ref := &Point{114.175329, 22.2524}
			So(Nearest(ref, []*Point{}), ShouldBeNil)
		})

		Convey("Should return nearest point", func() {
			ref := &Point{39.4, -75.4}
			point1 := &Point{39.284, -75.833}
			point2 := &Point{39.984, -75.6}
			point3 := &Point{39.125, -75.221}
			point4 := &Point{39.987, -75.358}
			point5 := &Point{39.27, -75.9221}
			point6 := &Point{39.123, -75.534}
			point7 := &Point{39.12, -75.21}
			point8 := &Point{39.33, -75.22}
			point9 := &Point{39.55, -75.44}
			point10 := &Point{39.66, -75.77}
			point11 := &Point{39.11, -75.44}
			point12 := &Point{39.92, -75.05}
			point13 := &Point{39.98, -75.88}
			point14 := &Point{39.55, -75.55}
			point15 := &Point{39.44, -75.33}
			point16 := &Point{39.24, -75.56}
			point17 := &Point{39.36, -75.56}
			points := []*Point{point1, point2, point3, point4, point5, point6, point7, point8, point9, point10,
				point11, point12, point13, point14, point15, point16, point17}
			So(Nearest(ref, points), ShouldResemble, &Point{39.44, -75.33})
		})

	})

}
