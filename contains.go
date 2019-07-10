package geo

import (
	"math"
)

// Pt is a representation of a 2D point. float32 is used to save 50% of the
// space - as this can add up with large polygons.
type Pt struct {
	Lat, Lng float32
}
type Pt64 struct {
	Lat, Lng float64
}

// Poly is a polygon represented by series of points. The lines of a polygon
// must be non-intersecting. Every simple polygon must begin and end with the
// same point. If a point is repeated, it is assumed to complete the polygon.
// A complex polygon with holesand/or islands is represented by sequential
// points representing complete simple polygons.
type Poly struct {
	Pts [](Pt)
}

// Contains performs a point-in-polygon test.
func (p Poly) Contains(point Pt) bool {

	start := len(p.Pts) - 1
	end := 0

	contains := false

	// If the first point does not match the last point, the polygon contains at
	// least one hole. If there are no holes, we can use simpler and faster logic.
	if p.Pts[start].Lat == p.Pts[end].Lat && p.Pts[start].Lng == p.Pts[end].Lng {
		for i := 1; i < len(p.Pts); i++ {
			if p.intersectsWithRaycast(point, p.Pts[i-1], p.Pts[i]) {
				contains = !contains
			}
		}
	} else {
		pt := p.Pts[0]
		ic := 1
		for i := 1; i < len(p.Pts); i++ {

			// if true, we are transitioning to a hole, skip
			if i != ic && pt.Lat == p.Pts[i-1].Lat && pt.Lng == p.Pts[i-1].Lng {
				pt = p.Pts[i]
				ic = i + 1
			} else {
				if p.intersectsWithRaycast(point, p.Pts[i-1], p.Pts[i]) {
					contains = !contains
				}
			}
		}
	}

	return contains
}

// Using the raycast algorithm, this returns whether or not the passed in point
// Intersects with the edge drawn by the passed in start and end points.
// Original implementation: http://rosettacode.org/wiki/Ray-casting_algorithm#Go
func (p Poly) intersectsWithRaycast(point Pt, start Pt, end Pt) bool {

	// The ray is cast in the east direction

	// The raycast algorithm requires the start point have a lesser or equal Lng
	// value
	if start.Lng > end.Lng {
		// Switch the points if otherwise.
		start, end = end, start
	}

	// Avoid the edge case where a ray intersects with a vertex.
	if point.Lng == start.Lng || point.Lng == end.Lng {

		newLng := math.Nextafter32(point.Lng, float32(math.Inf(1)))

		// If we collide with the other vertex, we need to increase precision to
		// get intermediate point on the line segment.
		// HACK: clone the rest of the logic with float64 instead of float32
		if newLng == start.Lng || newLng == end.Lng {
			// convert sources to float64 points
			point64 := Pt64{float64(point.Lat), float64(point.Lng)}
			start64 := Pt64{float64(start.Lat), float64(start.Lng)}
			end64 := Pt64{float64(end.Lat), float64(end.Lng)}

			newLng64 := math.Nextafter(point64.Lng, math.Inf(1))
			point64 = Pt64{point64.Lat, newLng64}

			// If the point64 is north or south of both start64 and end64, it
			// does not intersect.
			if point64.Lng < start64.Lng || end64.Lng < point64.Lng {
				return false
			}

			// If the point64 is east of both start64 and end64, it does not
			// intersect. If the point64 is west of both start64 and end64, it
			// does intersect.
			if start64.Lat < end64.Lat {
				if end64.Lat < point64.Lat {
					return false
				}
				if point64.Lat < start64.Lat {
					return true
				}
			} else if end64.Lat < start64.Lat {
				if start64.Lat < point64.Lat {
					return false
				}
				if point64.Lat < end64.Lat {
					return true
				}
			} else { // start64.Lat == end64.Lat
				if start64.Lat < point64.Lat {
					return false
				}
				if point64.Lat <= start64.Lat {
					return true
				}
			}

			// if we get here, the point64 is within the box formed by the Lat
			// and Lng values of start64 and end64. If the slope formed by
			// start64 and end64 is
			ptSlope := (point64.Lng - start64.Lng) / (point64.Lat - start64.Lat)
			segSlope := (end64.Lng - start64.Lng) / (end64.Lat - start64.Lat)

			return segSlope <= ptSlope
		}
		point = Pt{point.Lat, newLng}
	}

	// If the point is north or south of both start and end, it does not
	// intersect.
	if point.Lng < start.Lng || end.Lng < point.Lng {
		return false
	}

	// If the point is east of both start and end, it does not intersect. If the
	// point is west of both start and end, it does intersect.
	if start.Lat < end.Lat {
		if end.Lat < point.Lat {
			return false
		}
		if point.Lat < start.Lat {
			return true
		}
	} else if end.Lat < start.Lat {
		if start.Lat < point.Lat {
			return false
		}
		if point.Lat < end.Lat {
			return true
		}
	} else { // start.Lat == end.Lat
		if start.Lat < point.Lat {
			return false
		}
		if point.Lat <= start.Lat {
			return true
		}
	}

	// if we get here, the point is within the box formed by the Lat and Lng values
	// of start and end. If the slope formed by start and end is
	ptSlope := (point.Lng - start.Lng) / (point.Lat - start.Lat)
	segSlope := (end.Lng - start.Lng) / (end.Lat - start.Lat)

	return segSlope <= ptSlope

}
