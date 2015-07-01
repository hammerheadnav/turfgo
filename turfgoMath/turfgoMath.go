package turfgoMath

import "math"

const(
  ThreeDecimalPlaces float64 = .001
  TwelveDecimalPlaces float64 = .000000000001
)

func RadToDegree(rad float64)  float64{
  return rad * 180 / math.Pi
}

func DegreeToRad(degree float64)  float64{
  return degree * math.Pi / 180
}

func IsEqualFloat(first float64, second float64, epsilon float64) bool{
  return math.Abs(first - second) < epsilon
}

func IsEqualFloatPair(p1X float64, p1Y float64, p2X float64, p2Y float64, epsilon float64) bool{
  return math.Abs(p1X - p2X) < epsilon && math.Abs(p1Y - p2Y) < epsilon
}
