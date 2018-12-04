package geo

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		lat1, lng1, lat2, lng2, dist float64
	}{
		{
			lat1: 90,
			lng1: 0,
			lat2: -90,
			lng2: 0,
			dist: 12437.1849,
		},
		{
			lat1: 0,
			lng1: 0,
			lat2: 0,
			lng2: 180,
			dist: 12437.1849,
		},
		{
			lat1: 0,
			lng1: -180,
			lat2: 0,
			lng2: 180,
			dist: 0,
		},
		{
			lat1: 0,
			lng1: 0,
			lat2: 0,
			lng2: 360,
			dist: 0,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.dist, Distance(test.lat1, test.lng1, test.lat2, test.lng2))
	}
}
