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
			ref := &Point{28.965797, 41.010086}
			point1 := &Point{28.973865, 41.011122}
			point2 := &Point{28.948459, 41.024204}
			point3 := &Point{28.938674, 41.013324}
			points := []*Point{point1, point2, point3}
			So(Nearest(ref, points), ShouldResemble, &Point{28.973865, 41.011122})
		})

	})

}
