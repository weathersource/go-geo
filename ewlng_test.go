package geo

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestEwLng(t *testing.T) {
	assert.Equal(t, float64(-180), EwLng(-180))
	assert.Equal(t, float64(-180), EwLng(180))
	assert.Equal(t, float64(0), EwLng(-360))
	assert.Equal(t, float64(0), EwLng(0))
	assert.Equal(t, float64(0), EwLng(360))
}
