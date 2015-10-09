package turfgo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLineDiff(t *testing.T) {
	Convey("Given empty first line, should return empty array", t, func() {
		points1 := []*Point{}
		points2 := []*Point{&Point{1, 0}, &Point{1, 2}, &Point{2, 2}, &Point{4, 4}}
		lineString1 := NewLineString(points1)
		lineString2 := NewLineString(points2)
		diffs := LineDiff(lineString1, lineString2)
		So(len(diffs), ShouldEqual, 0)
	})

	Convey("Given empty second line, should return first line", t, func() {
		points1 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}}
		points2 := []*Point{}
		lineString1 := NewLineString(points1)
		lineString2 := NewLineString(points2)
		diffs := LineDiff(lineString1, lineString2)
		So(len(diffs), ShouldEqual, 1)
		So(diffs[0], ShouldResemble, lineString1)
	})

	// 0 0 0 0
	Convey("Given non intersecting line segments, should give full line", t, func() {
		points1 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}}
		points2 := []*Point{&Point{1, 0}, &Point{1, 2}, &Point{2, 2}, &Point{4, 4}}
		lineString1 := NewLineString(points1)
		lineString2 := NewLineString(points2)
		diffs := LineDiff(lineString1, lineString2)
		So(len(diffs), ShouldEqual, 1)
		So(diffs[0], ShouldResemble, lineString1)
	})

	// X X X X
	Convey("Given full intersection, should give no line", t, func() {
		points1 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}}
		points2 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}}
		lineString1 := NewLineString(points1)
		lineString2 := NewLineString(points2)
		diffs := LineDiff(lineString1, lineString2)
		So(len(diffs), ShouldEqual, 0)
	})

	// X 0 0 X
	Convey("Given line have common start and end point, should give full line", t, func() {
		points1 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}}
		points2 := []*Point{&Point{0, 0}, &Point{1, 2}, &Point{2, 2}, &Point{4, 5}}
		lineString1 := NewLineString(points1)
		lineString2 := NewLineString(points2)
		diffs := LineDiff(lineString1, lineString2)
		So(len(diffs), ShouldEqual, 1)
		So(diffs[0], ShouldResemble, lineString1)
	})

	// 0 X 0 0
	Convey("Given line cross each other only once, should give full line", t, func() {
		points1 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}}
		points2 := []*Point{&Point{0, 1}, &Point{1, 1}, &Point{2, 2}, &Point{4, 4}}
		lineString1 := NewLineString(points1)
		lineString2 := NewLineString(points2)
		diffs := LineDiff(lineString1, lineString2)
		So(len(diffs), ShouldEqual, 1)
		So(diffs[0], ShouldResemble, lineString1)
	})

	// 0 X 0 0 X 0
	Convey("Given line cross each other multiple times but not overlap, should give full line", t, func() {
		points1 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}, &Point{4, 4}, &Point{4, 9}}
		points2 := []*Point{&Point{0, 1}, &Point{1, 1}, &Point{2, 2}, &Point{4, 4}, &Point{8, 2}}
		lineString1 := NewLineString(points1)
		lineString2 := NewLineString(points2)
		diffs := LineDiff(lineString1, lineString2)
		So(len(diffs), ShouldEqual, 1)
		So(diffs[0], ShouldResemble, lineString1)
	})

	// 0 X X 0
	Convey("Given one overlap, should give correct result", t, func() {
		points1 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}}
		points2 := []*Point{&Point{0, 1}, &Point{1, 1}, &Point{2, 3}, &Point{4, 4}}
		lineString1 := NewLineString(points1)
		lineString2 := NewLineString(points2)
		diff1 := NewLineString([]*Point{&Point{0, 0}, &Point{1, 1}})
		diff2 := NewLineString([]*Point{&Point{2, 3}, &Point{4, 5}})
		diffs := LineDiff(lineString1, lineString2)
		So(len(diffs), ShouldEqual, 2)
		So(diffs[0], ShouldResemble, diff1)
		So(diffs[1], ShouldResemble, diff2)
	})

	// 0 X X X
	Convey("Given one overlap which proceed till end, should give correct result", t, func() {
		points1 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}}
		points2 := []*Point{&Point{0, 1}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}}
		lineString1 := NewLineString(points1)
		lineString2 := NewLineString(points2)
		diff1 := NewLineString([]*Point{&Point{0, 0}, &Point{1, 1}})
		diffs := LineDiff(lineString1, lineString2)
		So(len(diffs), ShouldEqual, 1)
		So(diffs[0], ShouldResemble, diff1)
	})

	// X X 0 0
	Convey("Given one overlap which start from beginning, should give correct result", t, func() {
		points1 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}}
		points2 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 4}, &Point{4, 6}}
		lineString1 := NewLineString(points1)
		lineString2 := NewLineString(points2)
		diff1 := NewLineString([]*Point{&Point{1, 1}, &Point{2, 3}, &Point{4, 5}})
		diffs := LineDiff(lineString1, lineString2)
		So(len(diffs), ShouldEqual, 1)
		So(diffs[0], ShouldResemble, diff1)
	})

	// 0 X X 0 0 X X X 0 0 0 0 0
	Convey("Given multiple overlap, should give correct result", t, func() {
		points1 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}, &Point{9, 5}, &Point{8, 5},
			&Point{4, 5}, &Point{11, 7}, &Point{9, 2}, &Point{4, 9}, &Point{12, 21}, &Point{12, 7}, &Point{21, 7}}
		points2 := []*Point{&Point{0, 1}, &Point{1, 1}, &Point{2, 3}, &Point{4, 4}, &Point{6, 7}, &Point{8, 5},
			&Point{4, 5}, &Point{11, 7}, &Point{13, 6}, &Point{12, 7}}
		lineString1 := NewLineString(points1)
		lineString2 := NewLineString(points2)
		diff1 := NewLineString([]*Point{&Point{0, 0}, &Point{1, 1}})
		diff2 := NewLineString([]*Point{&Point{2, 3}, &Point{4, 5}, &Point{9, 5}, &Point{8, 5}})
		diff3 := NewLineString([]*Point{&Point{11, 7}, &Point{9, 2}, &Point{4, 9}, &Point{12, 21}, &Point{12, 7}, &Point{21, 7}})
		diffs := LineDiff(lineString1, lineString2)
		So(len(diffs), ShouldEqual, 3)
		So(diffs[0], ShouldResemble, diff1)
		So(diffs[1], ShouldResemble, diff2)
		So(diffs[2], ShouldResemble, diff3)
	})

	//X X 0 % X X X 0 0 0 X X
	Convey("Given multiple overlap which start and end with line, should give correct result", t, func() {
		points1 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}, &Point{9, 5}, &Point{8, 5},
			&Point{4, 5}, &Point{11, 7}, &Point{9, 2}, &Point{4, 9}, &Point{12, 21}, &Point{12, 7}}
		points2 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{12, 3}, &Point{4, 4}, &Point{9, 5}, &Point{8, 5},
			&Point{4, 5}, &Point{21, 4}, &Point{13, 6}, &Point{12, 7}, &Point{12, 21}, &Point{12, 7}, &Point{23, 7}}
		lineString1 := NewLineString(points1)
		lineString2 := NewLineString(points2)
		diff1 := NewLineString([]*Point{&Point{1, 1}, &Point{2, 3}, &Point{4, 5}, &Point{9, 5}})
		diff2 := NewLineString([]*Point{&Point{4, 5}, &Point{11, 7}, &Point{9, 2}, &Point{4, 9}, &Point{12, 21}})
		diffs := LineDiff(lineString1, lineString2)
		So(len(diffs), ShouldEqual, 2)
		So(diffs[0], ShouldResemble, diff1)
		So(diffs[1], ShouldResemble, diff2)
	})
}

func TestReduceDiffSegment(t *testing.T) {
	ls1 := NewLineString([]*Point{&Point{0, 0}, &Point{1, 1}})
	ls2 := NewLineString([]*Point{&Point{1, 1}, &Point{2, 3}})
	ls3 := NewLineString([]*Point{&Point{4, 5}, &Point{1, 1}})
	ls4 := NewLineString([]*Point{&Point{1, 1}, &Point{7, 8}})
	ls5 := NewLineString([]*Point{&Point{9, 7}, &Point{2, 4}})
	Convey("Given empty array, should return same", t, func() {
		emptySeg := []*LineString{}
		So(reduceDiffSegment(emptySeg), ShouldResemble, emptySeg)
	})

	Convey("Given one element array, should return same", t, func() {
		seg := []*LineString{ls1}
		So(reduceDiffSegment(seg), ShouldResemble, seg)
	})

	Convey("Given multiple segments ending on non continous, should stitch them", t, func() {
		seg := []*LineString{ls1, ls2, ls3, ls4, ls5}
		merged1 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}}
		merged2 := []*Point{&Point{4, 5}, &Point{1, 1}, &Point{7, 8}}
		merged3 := []*Point{&Point{9, 7}, &Point{2, 4}}
		exptectedResult := []*LineString{NewLineString(merged1), NewLineString(merged2), NewLineString(merged3)}
		So(reduceDiffSegment(seg), ShouldResemble, exptectedResult)
	})

	Convey("Given multiple segments ending on contionus, should stitch them", t, func() {
		seg := []*LineString{ls1, ls2, ls3, ls4}
		merged1 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}}
		merged2 := []*Point{&Point{4, 5}, &Point{1, 1}, &Point{7, 8}}
		exptectedResult := []*LineString{NewLineString(merged1), NewLineString(merged2)}
		So(reduceDiffSegment(seg), ShouldResemble, exptectedResult)
	})

	Convey("Given three continous segments, should stitch them", t, func() {
		ls3 = NewLineString([]*Point{&Point{2, 3}, &Point{4, 5}})
		seg := []*LineString{ls1, ls2, ls3}
		merged := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}}
		exptectedResult := []*LineString{NewLineString(merged)}
		So(reduceDiffSegment(seg), ShouldResemble, exptectedResult)
	})
}

func TestContainLocationPair(t *testing.T) {
	points := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}, &Point{1, 1}}
	Convey("Given two points, should return true if points are present after each other", t, func() {
		found := containLocationPair(points, &Point{1, 1}, &Point{2, 3})
		So(found, ShouldBeTrue)
	})

	Convey("Given two points, should return false if points are not present after each other", t, func() {
		found := containLocationPair(points, &Point{1, 1}, &Point{4, 5})
		So(found, ShouldBeFalse)
	})

	Convey("Given two points, should return false if point is on edge", t, func() {
		found := containLocationPair(points, &Point{1, 1}, &Point{0, 0})
		So(found, ShouldBeFalse)
	})
}

func TestLineDiffPercentage(t *testing.T) {
	Convey("Given empty first line, should return 0", t, func() {
		points1 := []*Point{}
		points2 := []*Point{&Point{1, 0}, &Point{1, 2}, &Point{2, 2}, &Point{4, 4}}
		lineString1 := NewLineString(points1)
		lineString2 := NewLineString(points2)
		p := LineDiffPercentage(lineString1, lineString2)
		So(p, ShouldEqual, 0)
	})

	Convey("Given empty second line, should return 100 percent", t, func() {
		points1 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}}
		points2 := []*Point{}
		lineString1 := NewLineString(points1)
		lineString2 := NewLineString(points2)
		p := LineDiffPercentage(lineString1, lineString2)
		So(p, ShouldEqual, 100)
	})

	// X X X X
	Convey("Given full intersection, should give 0 percent", t, func() {
		points1 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}}
		points2 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}}
		lineString1 := NewLineString(points1)
		lineString2 := NewLineString(points2)
		p := LineDiffPercentage(lineString1, lineString2)
		So(p, ShouldEqual, 0)
	})

	// 0 0 0 0
	Convey("Given non intersecting line segments, should give 100 percentage", t, func() {
		points1 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}}
		points2 := []*Point{&Point{1, 0}, &Point{1, 2}, &Point{2, 2}, &Point{4, 4}}
		lineString1 := NewLineString(points1)
		lineString2 := NewLineString(points2)
		p := LineDiffPercentage(lineString1, lineString2)
		So(p, ShouldEqual, 100)
	})

	// 0 X X 0
	Convey("Given one overlap, should give correct result", t, func() {
		points1 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}}
		points2 := []*Point{&Point{0, 1}, &Point{1, 1}, &Point{2, 3}, &Point{4, 4}}
		lineString1 := NewLineString(points1)
		lineString2 := NewLineString(points2)
		p := LineDiffPercentage(lineString1, lineString2)
		So(p, ShouldEqual, 50)
	})

	// 0 X X X
	Convey("Given one overlap which proceed till end, should give correct result", t, func() {
		points1 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}}
		points2 := []*Point{&Point{0, 1}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}}
		lineString1 := NewLineString(points1)
		lineString2 := NewLineString(points2)
		p := LineDiffPercentage(lineString1, lineString2)
		So(p, ShouldEqual, 25)
	})

	// X X 0 0
	Convey("Given one overlap which start from beginning, should give correct result", t, func() {
		points1 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}}
		points2 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 4}, &Point{4, 6}}
		lineString1 := NewLineString(points1)
		lineString2 := NewLineString(points2)
		p := LineDiffPercentage(lineString1, lineString2)
		So(p, ShouldEqual, 50)
	})

	// 0 X X 0 0 X X X 0 0 0 0 0
	Convey("Given multiple overlap, should give correct result", t, func() {
		points1 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}, &Point{9, 5}, &Point{8, 5},
			&Point{4, 5}, &Point{11, 7}, &Point{9, 2}, &Point{4, 9}, &Point{12, 21}, &Point{12, 7}, &Point{21, 7}}
		points2 := []*Point{&Point{0, 1}, &Point{1, 1}, &Point{2, 3}, &Point{4, 4}, &Point{6, 7}, &Point{8, 5},
			&Point{4, 5}, &Point{11, 7}, &Point{13, 6}, &Point{12, 7}}
		lineString1 := NewLineString(points1)
		lineString2 := NewLineString(points2)
		p := LineDiffPercentage(lineString1, lineString2)
		So(p, ShouldAlmostEqual, float64(8)/float64(13)*100)
	})

	//X X 0 % X X X 0 0 0 X X
	Convey("Given multiple overlap which start and end with line, should give correct result", t, func() {
		points1 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{2, 3}, &Point{4, 5}, &Point{9, 5}, &Point{8, 5},
			&Point{4, 5}, &Point{11, 7}, &Point{9, 2}, &Point{4, 9}, &Point{12, 21}, &Point{12, 7}}
		points2 := []*Point{&Point{0, 0}, &Point{1, 1}, &Point{12, 3}, &Point{4, 4}, &Point{9, 5}, &Point{8, 5},
			&Point{4, 5}, &Point{21, 4}, &Point{13, 6}, &Point{12, 7}, &Point{12, 21}, &Point{12, 7}, &Point{23, 7}}
		lineString1 := NewLineString(points1)
		lineString2 := NewLineString(points2)
		p := LineDiffPercentage(lineString1, lineString2)
		So(p, ShouldAlmostEqual, float64(5)/float64(12)*100)
	})

}
