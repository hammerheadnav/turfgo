package turfgo

import (
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDecodeLineString(t *testing.T) {
	Convey("Given geoJson file, should return lineString", t, func() {
		j, _ := ioutil.ReadFile("./testFiles/pointOnLine/line1.geojson")
		ls, err := DecodeLineStringFromFeatureJSON(j)

		points := ls.Points

		So(err, ShouldBeNil)
		So(len(points), ShouldEqual, 3)
		So(points[0], ShouldResemble, &Point{22.466878364528448, -97.88131713867188})
		So(points[1], ShouldResemble, &Point{22.175960091218524, -97.82089233398438})
		So(points[2], ShouldResemble, &Point{21.8704201873689, -97.6190185546875})
	})

}
