package utils

func DiffUint8(a uint8, b uint8) uint8 {
	if a > b {
		return a - b
	}
	return b - a
}

func MaxUint8(a uint8, b uint8) uint8 {
	if a > b {
		return a
	}
	return b
}

func MinUint8(a uint8, b uint8) uint8 {
	if a < b {
		return a
	}
	return b
}
