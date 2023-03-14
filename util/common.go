package util

func BoolToUint8(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

func Uint8ToBool(u uint8) bool {
	return u == 1
}
