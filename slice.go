package loe

// Filter iterates over elements of collection, returning an array of all elements predicate returns truthy for.
func Filter[V any](collection []V, predicate func(item V, index int) (bool, error)) ([]V, error) {
	result := make([]V, 0, len(collection))

	for i, item := range collection {
		if isTrue, err := predicate(item, i); err != nil {
			return nil, err
		} else if isTrue {
			result = append(result, item)
		}
	}

	return result, nil
}

// Map manipulates a slice and transforms it to a slice of another type.
func Map[T any, R any](collection []T, iteratee func(item T, index int) (R, error)) ([]R, error) {
	result := make([]R, len(collection))

	for i, item := range collection {
		item, err := iteratee(item, i)

		if err != nil {
			return nil, err
		}

		result[i] = item
	}

	return result, nil
}

// FilterMap returns a slice which obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
func FilterMap[T any, R any](collection []T, callback func(item T, index int) (R, bool, error)) ([]R, error) {
	result := []R{}

	for i, item := range collection {
		if r, ok, err := callback(item, i); err != nil {
			return nil, err
		} else if ok {
			result = append(result, r)
		}
	}

	return result, nil
}

// FlatMap manipulates a slice and transforms and flattens it to a slice of another type.
// The transform function can either return a slice or a `nil`, and in the `nil` case
// no value is added to the final slice.
func FlatMap[T any, R any](collection []T, iteratee func(item T, index int) ([]R, error)) ([]R, error) {
	result := make([]R, 0, len(collection))

	for i, item := range collection {
		items, err := iteratee(item, i)

		if err != nil {
			return nil, err
		}

		result = append(result, items...)
	}

	return result, nil
}

// Reduce reduces collection to a value which is the accumulated result of running each element in collection
// through accumulator, where each successive invocation is supplied the return value of the previous.
func Reduce[T any, R any](collection []T, accumulator func(agg R, item T, index int) (R, error), initial R) (R, error) {
	for i, item := range collection {
		var err error
		initial, err = accumulator(initial, item, i)

		if err != nil {
			return empty[R](), err
		}
	}

	return initial, nil
}

// ReduceRight helper is like Reduce except that it iterates over elements of collection from right to left.
func ReduceRight[T any, R any](collection []T, accumulator func(agg R, item T, index int) (R, error), initial R) (R, error) {
	for i := len(collection) - 1; i >= 0; i-- {
		var err error
		initial, err = accumulator(initial, collection[i], i)

		if err != nil {
			return empty[R](), err
		}
	}

	return initial, nil
}

// ForEach iterates over elements of collection and invokes iteratee for each element.
func ForEach[T any](collection []T, iteratee func(item T, index int) error) error {
	for i, item := range collection {
		if err := iteratee(item, i); err != nil {
			return err
		}
	}

	return nil
}

// Times invokes the iteratee n times, returning an array of the results of each invocation.
// The iteratee is invoked with index as argument.
func Times[T any](count int, iteratee func(index int) (T, error)) ([]T, error) {
	result := make([]T, count)

	for i := 0; i < count; i++ {
		var err error
		result[i], err = iteratee(i)

		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// UniqBy returns a duplicate-free version of an array, in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the array. It accepts `iteratee` which is
// invoked for each element in array to generate the criterion by which uniqueness is computed.
func UniqBy[T any, U comparable](collection []T, iteratee func(item T) (U, error)) ([]T, error) {
	result := make([]T, 0, len(collection))
	seen := make(map[U]struct{}, len(collection))

	for _, item := range collection {
		key, err := iteratee(item)

		if err != nil {
			return nil, err
		}

		if _, ok := seen[key]; ok {
			continue
		}

		seen[key] = struct{}{}
		result = append(result, item)
	}

	return result, nil
}

// GroupBy returns an object composed of keys generated from the results of running each element of collection through iteratee.
func GroupBy[T any, U comparable](collection []T, iteratee func(item T) (U, error)) (map[U][]T, error) {
	result := map[U][]T{}

	for _, item := range collection {
		key, err := iteratee(item)

		if err != nil {
			return nil, err
		}

		result[key] = append(result[key], item)
	}

	return result, nil
}

// PartitionBy returns an array of elements split into groups. The order of grouped values is
// determined by the order they occur in collection. The grouping is generated from the results
// of running each element of collection through iteratee.
func PartitionBy[T any, K comparable](collection []T, iteratee func(item T) (K, error)) ([][]T, error) {
	result := [][]T{}
	seen := map[K]int{}

	for _, item := range collection {
		key, err := iteratee(item)

		if err != nil {
			return nil, err
		}

		resultIndex, ok := seen[key]
		if !ok {
			resultIndex = len(result)
			seen[key] = resultIndex
			result = append(result, []T{})
		}

		result[resultIndex] = append(result[resultIndex], item)
	}

	return result, nil

	// unordered:
	// groups := GroupBy[T, K](collection, iteratee)
	// return Values[K, []T](groups)
}

// RepeatBy builds a slice with values returned by N calls of callback.
func RepeatBy[T any](count int, predicate func(index int) (T, error)) ([]T, error) {
	result := make([]T, 0, count)

	for i := 0; i < count; i++ {
		item, err := predicate(i)

		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	return result, nil
}

// KeyBy transforms a slice or an array of structs to a map based on a pivot callback.
func KeyBy[K comparable, V any](collection []V, iteratee func(item V) (K, error)) (map[K]V, error) {
	result := make(map[K]V, len(collection))

	for _, v := range collection {
		k, err := iteratee(v)

		if err != nil {
			return nil, err
		}

		result[k] = v
	}

	return result, nil
}

// Associate returns a map containing key-value pairs provided by transform function applied to elements of the given slice.
// If any of two pairs would have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original array.
func Associate[T any, K comparable, V any](collection []T, transform func(item T) (K, V, error)) (map[K]V, error) {
	result := make(map[K]V, len(collection))

	for _, t := range collection {
		k, v, err := transform(t)

		if err != nil {
			return nil, err
		}

		result[k] = v
	}

	return result, nil
}

// SliceToMap returns a map containing key-value pairs provided by transform function applied to elements of the given slice.
// If any of two pairs would have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original array.
// Alias of Associate().
func SliceToMap[T any, K comparable, V any](collection []T, transform func(item T) (K, V, error)) (map[K]V, error) {
	return Associate(collection, transform)
}

// DropWhile drops elements from the beginning of a slice or array while the predicate returns true.
func DropWhile[T any](collection []T, predicate func(item T) (bool, error)) ([]T, error) {
	i := 0
	for ; i < len(collection); i++ {
		if isTrue, err := predicate(collection[i]); err != nil {
			return nil, err
		} else if !isTrue {
			break
		}
	}

	result := make([]T, 0, len(collection)-i)
	return append(result, collection[i:]...), nil
}

// DropRightWhile drops elements from the end of a slice or array while the predicate returns true.
func DropRightWhile[T any](collection []T, predicate func(item T) (bool, error)) ([]T, error) {
	i := len(collection) - 1
	for ; i >= 0; i-- {
		if isTrue, err := predicate(collection[i]); err != nil {
			return nil, err
		} else if !isTrue {
			break
		}
	}

	result := make([]T, 0, i+1)
	return append(result, collection[:i+1]...), nil
}

// Reject is the opposite of Filter, this method returns the elements of collection that predicate does not return truthy for.
func Reject[V any](collection []V, predicate func(item V, index int) (bool, error)) ([]V, error) {
	result := []V{}

	for i, item := range collection {
		if isTrue, err := predicate(item, i); err != nil {
			return nil, err
		} else if !isTrue {
			result = append(result, item)
		}
	}

	return result, nil
}

// CountBy counts the number of elements in the collection for which predicate is true.
func CountBy[T any](collection []T, predicate func(item T) (bool, error)) (count int, err error) {
	for _, item := range collection {
		if isTrue, err := predicate(item); err != nil {
			return -1, err
		} else if isTrue {
			count++
		}
	}

	return count, nil
}

// CountValuesBy counts the number of each element return from mapper function.
// Is equivalent to chaining lo.Map and lo.CountValues.
func CountValuesBy[T any, U comparable](collection []T, mapper func(item T) (U, error)) (map[U]int, error) {
	result := make(map[U]int)

	for _, item := range collection {
		key, err := mapper(item)

		if err != nil {
			return nil, err
		}

		result[key]++
	}

	return result, nil
}
