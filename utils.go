package config

// Default returns the first non-zero value.
// If all values are zero, return the zero value.
//
//	Default("", "foo") // "foo"
//	Default("bar", "foo") // "bar"
//	Default("", "", "foo") // "foo"
func Default[T comparable](values ...T) T {
	var zero T
	for _, value := range values {
		if value != zero {
			return value
		}
	}
	return zero
}
