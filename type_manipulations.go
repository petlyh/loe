package loe

// Empty returns an empty value.
func empty[T any]() T {
	var zero T
	return zero
}
