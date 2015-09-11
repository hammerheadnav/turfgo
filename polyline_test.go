package turfgo

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPolylineEncoder(t *testing.T) {
	Convey("Should encode an array of locations as a polyline string", t, func() {
		coordinates := []*Point{{41.85703, -87.64069}, {41.80715, -87.62695}}
		polyline := EncodePolyline(coordinates)
		So(polyline, ShouldEqual, "men~Fhi|uOvvH{tA")
	})

	Convey("Should give empty string for empty locations", t, func() {
		coordinates := []*Point{}
		polyline := EncodePolyline(coordinates)
		So(polyline, ShouldEqual, "")
	})
}
