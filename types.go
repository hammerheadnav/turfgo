package turfgo

var R = map[string]float64{"mi": 3960,
  "km": 6373,
  "d": 57.2957795,
  "r": 1}

type Point struct{
  Lat float64
  Lng float64
}

func NewPoint(lat float64, lon float64) *Point{
  return &Point{lat, lon}
}
