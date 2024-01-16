package loe

// Find search an element in a slice based on a predicate. It returns element and true if element was found.
func Find[T any](collection []T, predicate func(item T) (bool, error)) (T, bool, error) {
	for _, item := range collection {
		if isTrue, err := predicate(item); err != nil {
			return empty[T](), false, err
		} else if isTrue {
			return item, true, nil
		}
	}

	return empty[T](), false, nil
}

// FindIndexOf searches an element in a slice based on a predicate and returns the index and true.
// It returns -1 and false if the element is not found.
func FindIndexOf[T any](collection []T, predicate func(item T) (bool, error)) (T, int, bool, error) {
	for i, item := range collection {
		if isTrue, err := predicate(item); err != nil {
			return empty[T](), -1, false, err
		} else if isTrue {
			return item, i, true, nil
		}
	}

	return empty[T](), -1, false, nil
}

// FindLastIndexOf searches last element in a slice based on a predicate and returns the index and true.
// It returns -1 and false if the element is not found.
func FindLastIndexOf[T any](collection []T, predicate func(item T) (bool, error)) (T, int, bool, error) {
	length := len(collection)

	for i := length - 1; i >= 0; i-- {
		if isTrue, err := predicate(collection[i]); err != nil {
			return empty[T](), -1, false, err
		} else if isTrue {
			return collection[i], i, true, nil
		}
	}

	return empty[T](), -1, false, nil
}

// FindOrElse search an element in a slice based on a predicate. It returns the element if found or a given fallback value otherwise.
func FindOrElse[T any](collection []T, fallback T, predicate func(item T) (bool, error)) (T, error) {
	for _, item := range collection {
		if isTrue, err := predicate(item); err != nil {
			return empty[T](), err
		} else if isTrue {
			return item, nil
		}
	}

	return fallback, nil
}

// FindKeyBy returns the key of the first element predicate returns truthy for.
func FindKeyBy[K comparable, V any](object map[K]V, predicate func(key K, value V) (bool, error)) (K, bool, error) {
	for k, v := range object {
		if isTrue, err := predicate(k, v); err != nil {
			return empty[K](), false, err
		} else if isTrue {
			return k, true, nil
		}
	}

	return empty[K](), false, nil
}

// FindUniquesBy returns a slice with all the unique elements of the collection.
// The order of result values is determined by the order they occur in the array. It accepts `iteratee` which is
// invoked for each element in array to generate the criterion by which uniqueness is computed.
func FindUniquesBy[T any, U comparable](collection []T, iteratee func(item T) (U, error)) ([]T, error) {
	isDupl := make(map[U]bool, len(collection))

	for _, item := range collection {
		key, err := iteratee(item)

		if err != nil {
			return nil, err
		}

		duplicated, ok := isDupl[key]
		if !ok {
			isDupl[key] = false
		} else if !duplicated {
			isDupl[key] = true
		}
	}

	result := make([]T, 0, len(collection)-len(isDupl))

	for _, item := range collection {
		key, err := iteratee(item)

		if err != nil {
			return nil, err
		}

		if duplicated := isDupl[key]; !duplicated {
			result = append(result, item)
		}
	}

	return result, nil
}

// FindDuplicatesBy returns a slice with the first occurrence of each duplicated elements of the collection.
// The order of result values is determined by the order they occur in the array. It accepts `iteratee` which is
// invoked for each element in array to generate the criterion by which uniqueness is computed.
func FindDuplicatesBy[T any, U comparable](collection []T, iteratee func(item T) (U, error)) ([]T, error) {
	isDupl := make(map[U]bool, len(collection))

	for _, item := range collection {
		key, err := iteratee(item)

		if err != nil {
			return nil, err
		}

		duplicated, ok := isDupl[key]
		if !ok {
			isDupl[key] = false
		} else if !duplicated {
			isDupl[key] = true
		}
	}

	result := make([]T, 0, len(collection)-len(isDupl))

	for _, item := range collection {
		key, err := iteratee(item)

		if err != nil {
			return nil, err
		}

		if duplicated := isDupl[key]; duplicated {
			result = append(result, item)
			isDupl[key] = false
		}
	}

	return result, nil
}
