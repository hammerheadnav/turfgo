package turfgo

const (
	infinity = 0x7FF0000000000000
)

//R is radius of earth
var R = map[string]float64{"mi": 3960,
	"km": 6373,
	"d":  57.2957795,
	"r":  1}

//Geometry is geoJson geometry
type Geometry interface {
	getPoints() []*Point
}

//Point geojson type
type Point struct {
	Lat float64
	Lng float64
}

func (p *Point) getPoints() []*Point {
	return []*Point{p}
}

//NewPoint creates a new point for given lat, lng
func NewPoint(lat float64, lon float64) *Point {
	return &Point{lat, lon}
}

//MultiPoint geojson type
type MultiPoint struct {
	Points []*Point
}

func (p *MultiPoint) getPoints() []*Point {
	return p.Points
}

//NewMultiPoint creates a new multiPoint for given points
func NewMultiPoint(points []*Point) *MultiPoint {
	return &MultiPoint{Points: points}
}

//LineString geojson type
type LineString struct {
	Points []*Point
}

func (p *LineString) getPoints() []*Point {
	return p.Points
}

//NewLineString creates a new lineString for given points
func NewLineString(points []*Point) *LineString {
	return &LineString{Points: points}
}

//MultiLineString geojson type
type MultiLineString struct {
	LineStrings []*LineString
}

func (p *MultiLineString) getPoints() []*Point {
	points := []*Point{}
	for _, lineString := range p.LineStrings {
		points = append(points, lineString.getPoints()...)
	}
	return points
}

//NewMultiLineString creates a new multiLineString for given lineStrings
func NewMultiLineString(lineStrings []*LineString) *MultiLineString {
	return &MultiLineString{LineStrings: lineStrings}
}

//Polygon geojson type
type Polygon struct {
	LineStrings []*LineString
}

func (p *Polygon) getPoints() []*Point {
	points := []*Point{}
	for _, lineString := range p.LineStrings {
		points = append(points, lineString.getPoints()...)
	}
	return points
}

//NewPolygon creates a new polygon for given lineStrings
func NewPolygon(lineStrings []*LineString) *Polygon {
	return &Polygon{LineStrings: lineStrings}
}

//MultiPolygon geojson type
type MultiPolygon struct {
	Polygons []*Polygon
}

func (p *MultiPolygon) getPoints() []*Point {
	points := []*Point{}
	for _, polygon := range p.Polygons {
		points = append(points, polygon.getPoints()...)
	}
	return points
}

//NewMultiPolygon creates a new multiPolygon for given polygons
func NewMultiPolygon(polygons []*Polygon) *MultiPolygon {
	return &MultiPolygon{Polygons: polygons}
}
