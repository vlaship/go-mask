package mask

const (
	maskedPart = "*****"
	minLen     = 5
	cut        = 2
)

// String masks the middle part of a string.
func String(s string) string {
	if len(s) <= minLen {
		return maskedPart
	}
	return s[:cut] + maskedPart + s[len(s)-cut:]
}

// Strings masks the middle part of a string slice.
func Strings(s []string) []string {
	for i, v := range s {
		s[i] = String(v)
	}

	return s
}

// Runes masks the middle part of a rune slice.
func Runes(r []rune) []rune {
	return []rune(String(string(r)))
}

// Bytes masks the middle part of a byte slice.
func Bytes(b []byte) []byte {
	return []byte(String(string(b)))
}
