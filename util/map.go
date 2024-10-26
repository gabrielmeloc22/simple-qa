package util

func Map[T, V any](iterable []T, fn func(T) V) []V {
	result := make([]V, len(iterable))

	for i, t := range iterable {
		result[i] = fn(t)
	}

	return result
}
