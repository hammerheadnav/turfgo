package turfgo

import "github.com/twpayne/gopolyline/polyline"

//EncodePolyline encodes given coordinates into a polyline
func EncodePolyline(coordinates []*Point) string {
	var flatC []float64
	for i := 0; i < len(coordinates); i++ {
		flatC = append(flatC, coordinates[i].Lat)
		flatC = append(flatC, coordinates[i].Lng)
	}

	return polyline.Encode(flatC, 2)
}

//DecodePolyline decodes given polyline and return coordinates
func DecodePolyline(line string) ([]*Point, error) {
	flatC, err := polyline.Decode(line, 2)
	if err != nil {
		return nil, err
	}
	var coordinates []*Point
	for i := 0; i < len(flatC)/2; i++ {
		point := &Point{Lat: flatC[2*i], Lng: flatC[2*i+1]}
		coordinates = append(coordinates, point)
	}
	return coordinates, nil
}
