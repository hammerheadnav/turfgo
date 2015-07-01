package randomRouteGenerator

import(
  "math/rand"
  "math"
  "time"
  "fmt"
  "errors"
  "github.com/shashanktomar/turfgo"
)

func init(){
  rand.Seed(time.Now().UnixNano())
}

type waypointGenerator interface{
    createWaypoints(location *turfgo.Point, distance float64, unit string) ([]*turfgo.Point, error)
}

type circleWaypointGenerator struct{}
type rectangleWaypointGenerator struct{}
type triangleWaypointGenerator struct{}

func (c circleWaypointGenerator) createWaypoints(location *turfgo.Point,
                      distance float64, unit string) ([]*turfgo.Point, error){
  result := make([]*turfgo.Point, 8)
  result[0] = location

  randomBearing := rand.Float64() * 360
  radius := distance/(2 * math.Pi)
  center, err := turfgo.Destination(location, radius, randomBearing, unit)

  if err != nil{
    return nil, err
  }

  for i := 1; i < 8; i++{
    nextBearing := (float64(180) + randomBearing) + float64(45 * i)
    waypoint,_ := turfgo.Destination(center, radius, nextBearing, unit)
    result[i] = waypoint
  }

  return result, nil
}

func (c triangleWaypointGenerator) createWaypoints(location *turfgo.Point,
                      distance float64, unit string) ([]*turfgo.Point, error){
  result := make([]*turfgo.Point, 3)
  result[0] = location

  randomBearing := rand.Float64() * 360
  secondPointOnTriangle, err := turfgo.Destination(location, distance/3, randomBearing, unit)
  if err != nil{
    return nil, err
  }
  result[1] = secondPointOnTriangle

  thirdPointOnTriangle, err := turfgo.Destination(location, distance/3, randomBearing + 60, unit)
  result[2] = thirdPointOnTriangle

  return result, nil
}


func (c rectangleWaypointGenerator) createWaypoints(location *turfgo.Point,
                        distance float64, unit string) ([]*turfgo.Point, error){
  result := make([]*turfgo.Point, 4)
  result[0] = location
  sideOfRec := distance/4


  randomBearing := rand.Float64() * 360
  secondPointOnRectangle, err := turfgo.Destination(location, sideOfRec, randomBearing, unit)
  if err != nil{
    return nil, err
  }
  result[1] = secondPointOnRectangle

  diagonal := math.Sqrt(sideOfRec*sideOfRec + sideOfRec*sideOfRec)
  thirdPointOnRectangle, err := turfgo.Destination(location, diagonal, randomBearing+45, unit)
  result[2] = thirdPointOnRectangle

  fourthPointOnRectangle, err := turfgo.Destination(location, sideOfRec, randomBearing+90, unit)
  result[3] = fourthPointOnRectangle

  return result, nil
}

var generatorMap = map[int]waypointGenerator{0: circleWaypointGenerator{},
  1:rectangleWaypointGenerator{},
  2:triangleWaypointGenerator{}}

func RandomWaypoints(location *turfgo.Point, distance float64, unit string) ([]*turfgo.Point, error){
  generator := generatorMap[rand.Intn(3)]
  return generator.createWaypoints(location, distance, unit)
}

func RandomWaypointsWithGivenStrategy(location *turfgo.Point, distance float64,
                        unit string, strategy string) ([]*turfgo.Point, error){
  var generator waypointGenerator = circleWaypointGenerator{}
  switch strategy {
    case "circle":
      generator = circleWaypointGenerator{}
    case "rectangle":
      generator = rectangleWaypointGenerator{}
    case "triangle":
      generator = triangleWaypointGenerator{}
    default:
      return nil, errors.New(fmt.Sprintf("%s is not a valid strategy. Allowed strategies are circle, rectangle and triangle", strategy))
  }

  return generator.createWaypoints(location, distance, unit)
}
