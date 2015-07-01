package randomRouteGenerator

import(
  "testing"
  "github.com/shashanktomar/turfgo"
)

func TestRandomWaypoints(t *testing.T){
  location := turfgo.NewPoint(39.984, -75.343)

  result, _ := RandomWaypointsGeoJson(location, 30, "kilometers")
  t.Log(result)
}

func TestGenRandomWaypointsError(t *testing.T){
  location := turfgo.NewPoint(39.984, -75.343)

  _, err := RandomWaypointsWithGivenStrategy(location, 30, "kilometers", "invalidStrategy")

  if err == nil{
    t.Errorf("Error should not be nil")
  }
}

func TestGenRandomWaypointsCircle(t *testing.T){
  location := turfgo.NewPoint(39.984, -75.343)

  distance := float64(30)
  result, err := RandomWaypointsWithGivenStrategy(location, distance, "kilometers", "circle")

  if err != nil{
    t.Errorf("Error should not be nil")
  }

  if len(result) != 8{
    t.Errorf("Circle should have 8 waypoints")
  }
}

func TestGenRandomWaypointsTriangle(t *testing.T){
  location := turfgo.NewPoint(39.984, -75.343)

  distance := float64(30)
  result, err := RandomWaypointsWithGivenStrategy(location, distance, "kilometers", "triangle")

  if err != nil{
    t.Errorf("Error should not be nil")
  }

  if len(result) != 3{
    t.Errorf("Circle should have 8 waypoints")
  }
}

func TestGenRandomWaypointsRectangle(t *testing.T){
  location := turfgo.NewPoint(39.984, -75.343)

  distance := float64(30)
  result, err := RandomWaypointsWithGivenStrategy(location, distance, "kilometers", "rectangle")

  if err != nil{
    t.Errorf("Error should not be nil")
  }

  if len(result) != 4{
    t.Errorf("Circle should have 8 waypoints")
  }
}
