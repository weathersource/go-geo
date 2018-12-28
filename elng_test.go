package geo

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestELng(t *testing.T) {
	assert.Equal(t, float64(180), ELng(-180))
	assert.Equal(t, float64(180), ELng(180))
	assert.Equal(t, float64(180), ELng(540))
	assert.Equal(t, float64(0), ELng(0))
	assert.Equal(t, float64(0), ELng(360))
}

func TestELng32(t *testing.T) {
	assert.Equal(t, float32(180), ELng32(-180))
	assert.Equal(t, float32(180), ELng32(180))
	assert.Equal(t, float32(180), ELng32(540))
	assert.Equal(t, float32(0), ELng32(0))
	assert.Equal(t, float32(0), ELng32(360))
}
