package geo

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
	// location "github.com/weathersource/protorepo/go/location"
)

func TestELng(t *testing.T) {
	assert.Equal(t, float64(180), ELng(-180))
	assert.Equal(t, float64(180), ELng(180))
	assert.Equal(t, float64(180), ELng(540))
	assert.Equal(t, float64(0), ELng(0))
	assert.Equal(t, float64(0), ELng(360))
}
