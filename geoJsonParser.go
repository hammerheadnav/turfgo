package turfgo

import "github.com/kpawlik/geojson"

func EncodePoint(point *Point)(*geojson.Point){
  var c geojson.Coordinate
  c[0] = geojson.CoordType(point.longitude)
  c[1] = geojson.CoordType(point.latitude)
  return geojson.NewPoint(c)
}

func EncodeMultiPointsIntoFeature(points []*Point)(*geojson.Feature){
  var coordinates = make(geojson.Coordinates, len(points))
  for i := range points{
    var c geojson.Coordinate
    c[0] = geojson.CoordType(points[i].longitude)
    c[1] = geojson.CoordType(points[i].latitude)
    coordinates[i] = c
  }
  multiPoint := geojson.NewMultiPoint(coordinates)
  return geojson.NewFeature(multiPoint, nil, nil)
}

func EncodeFeatureCollection(features []*geojson.Feature)(*geojson.FeatureCollection){
  return geojson.NewFeatureCollection(features)
}
