package geo

import (
	"math"

	mathext "github.com/weathersource/go-mathext"
)

// Distance calculates distance in miles between two lat/lng values.
func Distance(lat1, lng1, lat2, lng2 float64) float64 {

	// Haversine Formula Optomized: https://stackoverflow.com/questions/27928#21623206
	d2r := 0.0174532925199 // PI/180
	di := 7917.7578304     // 2x mean radius of Earth converted to from km to miles (2*6371.2*0.621371)
	a := 0.5 - math.Cos((lat2-lat1)*d2r)/2 +
		math.Cos(lat1*d2r)*math.Cos(lat2*d2r)*(1-math.Cos((lng2-lng1)*d2r))/2
	dist := di * math.Asin(math.Sqrt(a))
	return mathext.Roundp(dist, 4)
}

// Distance calculates distance in miles between two lat/lng values.
func Distance32(lat1, lng1, lat2, lng2 float32) float32 {

	// Haversine Formula Optomized: https://stackoverflow.com/questions/27928#21623206
	d2r := 0.0174532925199 // PI/180
	di := 7917.7578304     // 2x mean radius of Earth converted to from km to miles (2*6371.2*0.621371)
	a := 0.5 - math.Cos(float64(lat2-lat1)*d2r)/2 +
		math.Cos(float64(lat1)*d2r)*math.Cos(float64(lat2)*d2r)*(1-math.Cos(float64(lng2-lng1)*d2r))/2
	dist := di * math.Asin(math.Sqrt(a))
	return mathext.Roundp32(float32(dist), 4)
}
