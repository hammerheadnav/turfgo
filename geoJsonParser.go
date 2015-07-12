package turfgo

import "github.com/kpawlik/geojson"

func EncodePoint(point *Point)(*geojson.Point){
  var c geojson.Coordinate
  c[0] = geojson.CoordType(point.Lat)
  c[1] = geojson.CoordType(point.Lng)
  return geojson.NewPoint(c)
}

func EncodeMultiPointsIntoFeature(points []*Point)(*geojson.Feature){
  var coordinates = make(geojson.Coordinates, len(points))
  for i := range points{
    var c geojson.Coordinate
    c[0] = geojson.CoordType(points[i].Lat)
    c[1] = geojson.CoordType(points[i].Lng)
    coordinates[i] = c
  }
  multiPoint := geojson.NewMultiPoint(coordinates)
  return geojson.NewFeature(multiPoint, nil, nil)
}

func EncodeFeatureCollection(features []*geojson.Feature)(*geojson.FeatureCollection){
  return geojson.NewFeatureCollection(features)
}
