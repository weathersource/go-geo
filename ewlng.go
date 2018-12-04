package geo

// EwLng converts lng to a east/west longitude value [-180, 180)
func EwLng(lng float64) float64 {
	for 180 <= lng {
		lng -= 360
	}
	for lng < -180 {
		lng += 360
	}
	return lng
}
