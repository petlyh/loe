package loe

// MapKeys manipulates a map keys and transforms it to a map of another type.
func MapKeys[K comparable, V any, R comparable](in map[K]V, iteratee func(value V, key K) (R, error)) (map[R]V, error) {
	result := make(map[R]V, len(in))

	for k, v := range in {
		key, err := iteratee(v, k)

		if err != nil {
			return nil, err
		}

		result[key] = v
	}

	return result, nil
}

// MapValues manipulates a map values and transforms it to a map of another type.
func MapValues[K comparable, V any, R any](in map[K]V, iteratee func(value V, key K) (R, error)) (map[K]R, error) {
	result := make(map[K]R, len(in))

	for k, v := range in {
		value, err := iteratee(v, k)

		if err != nil {
			return nil, err
		}

		result[k] = value
	}

	return result, nil
}

// MapEntries manipulates a map entries and transforms it to a map of another type.
func MapEntries[K1 comparable, V1 any, K2 comparable, V2 any](in map[K1]V1, iteratee func(key K1, value V1) (K2, V2, error)) (map[K2]V2, error) {
	result := make(map[K2]V2, len(in))

	for k1, v1 := range in {
		k2, v2, err := iteratee(k1, v1)

		if err != nil {
			return nil, err
		}

		result[k2] = v2
	}

	return result, nil
}

// MapToSlice transforms a map into a slice based on specific iteratee
func MapToSlice[K comparable, V any, R any](in map[K]V, iteratee func(key K, value V) (R, error)) ([]R, error) {
	result := make([]R, 0, len(in))

	for k, v := range in {
		item, err := iteratee(k, v)

		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	return result, nil
}
