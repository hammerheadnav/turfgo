package turfgoMath

import (
  "testing"
)

func TestRadToDegree(t *testing.T){
  var expected float64 = 57.295
  result := RadToDegree(1)
  if !IsEqualFloat(result, expected, ThreeDecimalPlaces){
    t.Errorf("Expected: %g, Actual: %g", expected, result)
  }
}


func TestDegreeToRand(t *testing.T){
  var expected float64 = 0.017
  result := DegreeToRad(1)
  if !IsEqualFloat(result, expected, ThreeDecimalPlaces){
    t.Errorf("Expected: %g, Actual: %g", expected, result)
  }
}
