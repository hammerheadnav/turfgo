package turfgo

import (
  "testing"
  "github.com/shashanktomar/turfgo/turfgoMath"
)

func TestBearing(t *testing.T){
  point1 := &Point{39.984, -75.343}
  point2 := &Point{39.123, -75.534}

  var expected1 float64 = -170.2330491349224
  result1 := Bearing(point1, point2)
  if !turfgoMath.IsEqualFloat(result1, expected1, turfgoMath.TwelveDecimalPlaces){
    t.Errorf("Expected: %g, Actual: %g", expected1, result1)
  }

  point3 := &Point{12.9715987, 77.59456269999998}
  point4 := &Point{13.22328378, 77.77448784}

  var expected2 float64 = 34.828578946361255
  result2 := Bearing(point3, point4)
  if !turfgoMath.IsEqualFloat(result2, expected2, turfgoMath.TwelveDecimalPlaces){
    t.Errorf("Expected: %g, Actual: %g", expected2, result2)
  }
}

func TestDestinationForError(t *testing.T){
  startingPoint := &Point{39.984, -75.343}
  if _, ok := Destination(startingPoint, 32, 120, "invalidUnit"); ok == nil{
    t.Errorf("Should through an error if unit is not valid")
  }
}

func TestDestinationMiles(t *testing.T){
  startingPoint := &Point{39.984, -75.343}
  expected := &Point{39.74662966576427, -75.81645928866797}

  result, ok := Destination(startingPoint, 30, -123, "miles")
  if ok != nil{
    t.Errorf("Error should be nil")
  }
  if !turfgoMath.IsEqualFloatPair(result.latitude, result.longitude, expected.latitude, expected.longitude, turfgoMath.TwelveDecimalPlaces){
    t.Errorf("Expected: %v, Actual: %v", expected, result)
  }
}

func TestDestinationKilometers(t *testing.T){
  startingPoint := &Point{39.984, -75.343}
  expected := &Point{40.01636403124377, -75.20865245149336}

  result, ok := Destination(startingPoint, 12, 72.5, "kilometers")
  if ok != nil{
    t.Errorf("Error should be nil")
  }
  if !turfgoMath.IsEqualFloatPair(result.latitude, result.longitude, expected.latitude, expected.longitude, turfgoMath.TwelveDecimalPlaces){
    t.Errorf("Expected: %v, Actual: %v", expected, result)
  }
}

func TestDestinationRadians(t *testing.T){
  startingPoint := &Point{39.984, -75.343}
  expected := &Point{67.3178236932749, -216.61938960828266}

  result, ok := Destination(startingPoint, 1.2, 345, "radians")
  if ok != nil{
    t.Errorf("Error should be nil")
  }
  if !turfgoMath.IsEqualFloatPair(result.latitude, result.longitude, expected.latitude, expected.longitude, turfgoMath.TwelveDecimalPlaces){
    t.Errorf("Expected: %v, Actual: %v", expected, result)
  }
}

func TestDestinationDegrees(t *testing.T){
  startingPoint := &Point{39.984, -75.343}
  expected := &Point{-13.74474598397336, -92.31513759524121}

  result, ok := Destination(startingPoint, 56, 200, "degrees")
  if ok != nil{
    t.Errorf("Error should be nil")
  }
  if !turfgoMath.IsEqualFloatPair(result.latitude, result.longitude, expected.latitude, expected.longitude, turfgoMath.TwelveDecimalPlaces){
    t.Errorf("Expected: %v, Actual: %v", expected, result)
  }
}

func TestDistanceError(t *testing.T)  {
  point1 := &Point{39.984, -75.343}
  point2 := &Point{39.97074218352032, -75.4590397138299}

  if _, ok := Distance(point1,point2, "invalidUnit"); ok == nil{
    t.Errorf("Should through an error if unit is not valid")
  }
}

func TestDistance(t *testing.T)  {
  point1 := &Point{39.984, -75.343}
  point2 := &Point{39.97074218352032, -75.4590397138299}
  expected := 9.999999999999373

  result, ok := Distance(point1, point2, "kilometers")
  if ok != nil{
    t.Errorf("Error should be nil")
  }

  if !turfgoMath.IsEqualFloat(result, expected, turfgoMath.TwelveDecimalPlaces){
    t.Errorf("Expected: %g, Actual: %g", expected, result)
  }
}
