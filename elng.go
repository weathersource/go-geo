package geo

// ELng converts lng to a full east longitude value [0, 360)
func ELng(lng float64) float64 {
	for 360 <= lng {
		lng -= 360
	}
	for lng < 0 {
		lng += 360
	}
	return lng
}
