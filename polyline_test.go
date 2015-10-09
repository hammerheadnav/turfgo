package turfgo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPolylineEncoder(t *testing.T) {
	Convey("Should encode an array of locations as a polyline string", t, func() {
		coordinates := []*Point{{38.5, -120.2}, {40.7, -120.95}, {43.252, -126.453}}
		polyline := EncodePolyline(coordinates)
		So(polyline, ShouldEqual, "_p~iF~ps|U_ulLnnqC_mqNvxq`@")
	})

	Convey("Should give empty string for empty locations", t, func() {
		coordinates := []*Point{}
		polyline := EncodePolyline(coordinates)
		So(polyline, ShouldEqual, "")
	})
}

func TestPolylineDecoder(t *testing.T) {
	Convey("Should decode a string as an array of locations", t, func() {
		result := []*Point{{38.5, -120.2}, {40.7, -120.95}, {43.252, -126.453}}
		coordinates, err := DecodePolyline("_p~iF~ps|U_ulLnnqC_mqNvxq`@")
		So(err, ShouldBeNil)
		So(len(coordinates), ShouldEqual, 3)
		So(coordinates[0], ShouldResemble, result[0])
		So(coordinates[1], ShouldResemble, result[1])
		So(coordinates[2], ShouldResemble, result[2])
	})

	Convey("Should throw error if invalid string", t, func() {
		_, err := DecodePolyline("invalidPolyline")
		So(err, ShouldNotBeNil)
	})
}
