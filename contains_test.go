package geo

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	tests := []struct {
		poly      Poly
		pt        Pt
		contained bool
	}{
		{ // north of poly
			poly:      Poly{[]Pt{{1, 1}, {1, -1}, {-1, -1}, {1, 1}}},
			pt:        Pt{0, 2},
			contained: false,
		},
		{ // south of poly
			poly:      Poly{[]Pt{{1, 1}, {1, -1}, {-1, -1}, {1, 1}}},
			pt:        Pt{0, -2},
			contained: false,
		},
		{ // east of poly with positive slope segment
			poly:      Poly{[]Pt{{1, 1}, {1, -1}, {-1, -1}, {1, 1}}},
			pt:        Pt{2, 0},
			contained: false,
		},
		{ // west of poly with positive slope segment
			poly:      Poly{[]Pt{{1, 1}, {1, -1}, {-1, -1}, {1, 1}}},
			pt:        Pt{-2, 0},
			contained: false,
		},
		{ // east of poly with negative slope segment
			poly:      Poly{[]Pt{{-1, 1}, {1, -1}, {-1, -1}, {-1, 1}}},
			pt:        Pt{2, 0},
			contained: false,
		},
		{ // west of poly with negative slope segment
			poly:      Poly{[]Pt{{-1, 1}, {1, -1}, {-1, -1}, {-1, 1}}},
			pt:        Pt{-2, 0},
			contained: false,
		},
		{ // within poly bounding box, but outside poly
			poly:      Poly{[]Pt{{1, 1}, {1, -1}, {-1, -1}, {1, 1}}},
			pt:        Pt{-.5, 0},
			contained: false,
		},
		{ // inside poly
			poly:      Poly{[]Pt{{-1, 1}, {1, -1}, {-1, -1}, {-1, 1}}},
			pt:        Pt{-.5, 0},
			contained: true,
		},

		{ // on border line of simple poly (eastern boundary, raycast has 1 intersection)
			poly:      Poly{[]Pt{{1, 1}, {1, -1}, {-1, -1}, {1, 1}}},
			pt:        Pt{1, 0},
			contained: true,
		},
		{ // on border line of simple poly (western boundary, raycast has 2 intersections)
			poly:      Poly{[]Pt{{1, 1}, {1, -1}, {-1, -1}, {1, 1}}},
			pt:        Pt{0, 0},
			contained: false,
		},
		{ // on vertex line of simple poly (pt shifts out of poly)
			poly:      Poly{[]Pt{{1, 1}, {1, -1}, {-1, -1}, {1, 1}}},
			pt:        Pt{1, 1},
			contained: false,
		},
		{ // on vertex line of simple poly (pt shifts in to poly)
			poly:      Poly{[]Pt{{1, 1}, {1, -1}, {-1, -1}, {1, 1}}},
			pt:        Pt{1, -1},
			contained: true,
		},
		{ // in a hole
			poly: Poly{[]Pt{
				{2, 2}, {2, -2}, {-2, -2}, {-2, 2}, {2, 2},
				{1, 1}, {1, -1}, {-1, -1}, {-1, 1}, {1, 1},
			}},
			pt:        Pt{0, 0},
			contained: false,
		},
		{ // in a poly where raycast crosses segment between container and hole
			poly: Poly{[]Pt{
				{2, 2}, {2, -2}, {-2, -2}, {-2, 2}, {2, 2},
				{1, 1}, {1, -1}, {-1, -1}, {-1, 1}, {1, 1},
			}},
			pt:        Pt{0, 1.5},
			contained: true,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.contained, test.poly.Contains(test.pt))
	}
}
