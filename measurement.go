package turfgo

import (
  "math"
  "fmt"
  "errors"
  "github.com/shashanktomar/turfgo/turfgoMath"
)

func Bearing(point1, point2 *Point) float64 {
  lat1 := turfgoMath.DegreeToRad(point1.latitude)
  lat2 := turfgoMath.DegreeToRad(point2.latitude)
  lon1 := turfgoMath.DegreeToRad(point1.longitude)
  lon2 := turfgoMath.DegreeToRad(point2.longitude)
  a := math.Sin(lon2 - lon1) * math.Cos(lat2)
  b := math.Cos(lat1) * math.Sin(lat2) -
        math.Sin(lat1) * math.Cos(lat2) * math.Cos(lon2 - lon1)
  return turfgoMath.RadToDegree(math.Atan2(a, b))
}

func Destination(startingPoint *Point, distance float64, bearing float64, unit string) (*Point, error){
  radius, ok := R[unit]
  if !ok{
    return nil, errors.New(fmt.Sprintf("%s is not a valid unit. Allowed units are miles, kilometers, degrees and radians", unit))
  }

  lat := turfgoMath.DegreeToRad(startingPoint.latitude)
  lon := turfgoMath.DegreeToRad(startingPoint.longitude)
  bearingRad := turfgoMath.DegreeToRad(bearing)

  destLat := math.Asin(math.Sin(lat) * math.Cos(distance / radius) +
        math.Cos(lat) * math.Sin(distance / radius) * math.Cos(bearingRad));
  destLon := lon + math.Atan2(math.Sin(bearingRad) * math.Sin(distance / radius) * math.Cos(lat),
        math.Cos(distance / radius) - math.Sin(lat) * math.Sin(destLat));

  return &Point{turfgoMath.RadToDegree(destLat), turfgoMath.RadToDegree(destLon)}, nil
}

func Distance(point1 *Point, point2 *Point, unit string) (float64, error){
  radius, ok := R[unit]
  if !ok{
    return 0, errors.New(fmt.Sprintf("%s is not a valid unit. Allowed units are miles, kilometers, degrees and radians", unit))
  }

  dLat := turfgoMath.DegreeToRad(point2.latitude - point1.latitude);
  dLon := turfgoMath.DegreeToRad(point2.longitude - point1.longitude);
  latRad1 := turfgoMath.DegreeToRad(point1.latitude);
  latRad2 := turfgoMath.DegreeToRad(point2.latitude);
  a := math.Sin(dLat/2) * math.Sin(dLat/2) +
          math.Sin(dLon/2) * math.Sin(dLon/2) * math.Cos(latRad1) * math.Cos(latRad2);
  c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a));
  return radius*c, nil
}
