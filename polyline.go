package turfgo

import "math"

// Ported from https://github.com/mapbox/polyline/blob/master/src/polyline.js

func EncodePolyline(coordinates []*Point) string {
	if len(coordinates) == 0 {
		return ""
	}

	factor := math.Pow(10, 5)
	output := encode(coordinates[0].Lat, factor) + encode(coordinates[0].Lng, factor)

	for i := 1; i < len(coordinates); i++ {
		a := coordinates[i]
		b := coordinates[i-1]
		output += encode(a.Lat-b.Lat, factor)
		output += encode(a.Lng-b.Lng, factor)
	}

	return output
}

func encode(oldCoordinate float64, factor float64) string {
	coordinate := int(math.Floor(oldCoordinate*factor + 0.5))
	coordinate = coordinate << 1

	if coordinate < 0 {
		coordinate = ^coordinate
	}
	output := ""
	for coordinate >= 0x20 {
		runeC := string((0x20 | (coordinate & 0x1f)) + 63)
		output = output + runeC
		coordinate >>= 5
	}
	runeC := string(coordinate + 63)
	output = output + runeC
	return output
}
