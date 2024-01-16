package loe

// ContainsBy returns true if predicate function return true.
func ContainsBy[T any](collection []T, predicate func(item T) (bool, error)) (bool, error) {
	for _, item := range collection {
		if isTrue, err := predicate(item); err != nil {
			return false, err
		} else if isTrue {
			return true, nil
		}
	}

	return false, nil
}

// EveryBy returns true if the predicate returns true for all of the elements in the collection or if the collection is empty.
func EveryBy[T any](collection []T, predicate func(item T) (bool, error)) (bool, error) {
	for _, v := range collection {
		if isTrue, err := predicate(v); err != nil {
			return false, err
		} else if !isTrue {
			return false, nil
		}
	}

	return true, nil
}

// SomeBy returns true if the predicate returns true for any of the elements in the collection.
// If the collection is empty SomeBy returns false.
func SomeBy[T any](collection []T, predicate func(item T) (bool, error)) (bool, error) {
	for _, v := range collection {
		if isTrue, err := predicate(v); err != nil {
			return false, err
		} else if isTrue {
			return true, nil
		}
	}

	return false, nil
}

// NoneBy returns true if the predicate returns true for none of the elements in the collection or if the collection is empty.
func NoneBy[T any](collection []T, predicate func(item T) (bool, error)) (bool, error) {
	for _, v := range collection {
		if isTrue, err := predicate(v); err != nil {
			return false, err
		} else if isTrue {
			return false, nil
		}
	}

	return true, nil
}
