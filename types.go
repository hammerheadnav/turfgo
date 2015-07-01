package turfgo

var R = map[string]float64{"miles": 3960,
  "kilometers": 6373,
  "degrees": 57.2957795,
  "radians": 1}


type Point struct{
  latitude float64
  longitude float64
}

func NewPoint(lat float64, lon float64) *Point{
  return &Point{lat, lon}
}
