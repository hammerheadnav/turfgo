package turfgo

import (
	"encoding/json"

	"github.com/kpawlik/geojson"
)

func EncodePoint(point *Point) *geojson.Point {
	var c geojson.Coordinate
	c[0] = geojson.CoordType(point.Lng)
	c[1] = geojson.CoordType(point.Lat)
	return geojson.NewPoint(c)
}

func DecodePoint(coord geojson.Coordinate) *Point {
	return &Point{float64(coord[1]), float64(coord[0])}
}

func EncodeMultiPointsIntoFeature(points []*Point) *geojson.Feature {
	var coordinates = make(geojson.Coordinates, len(points))
	for i := range points {
		var c geojson.Coordinate
		c[0] = geojson.CoordType(points[i].Lng)
		c[1] = geojson.CoordType(points[i].Lat)
		coordinates[i] = c
	}
	multiPoint := geojson.NewMultiPoint(coordinates)
	return geojson.NewFeature(multiPoint, nil, nil)
}

func EncodeMultiPointsIntoLineString(points []*Point) *geojson.Feature {
	var coordinates = make(geojson.Coordinates, len(points))
	for i := range points {
		var c geojson.Coordinate
		c[0] = geojson.CoordType(points[i].Lng)
		c[1] = geojson.CoordType(points[i].Lat)
		coordinates[i] = c
	}
	lineString := geojson.NewLineString(coordinates)
	return geojson.NewFeature(lineString, nil, nil)
}

func EncodeFeatureCollection(features []*geojson.Feature) *geojson.FeatureCollection {
	return geojson.NewFeatureCollection(features)
}

// DecodeLineStringFromFeatureJSON decode geojson feature type lineString into *LineString
func DecodeLineStringFromFeatureJSON(gj []byte) (*LineString, error) {
	var f *geojson.Feature
	json.Unmarshal(gj, &f)
	g, err := f.GetGeometry()
	if err != nil {
		return nil, err
	}
	ls, _ := g.(*geojson.LineString)
	points := []*Point{}
	for _, c := range ls.Coordinates {
		points = append(points, DecodePoint(c))
	}
	return NewLineString(points), nil
}
