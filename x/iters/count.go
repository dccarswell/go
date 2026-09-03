package iters

// Seq

// Count fully drains seq and returns the number of elements it yielded.
func Count[T any](seq Seq[T]) int {
	count := 0
	for _ = range seq {
		count++
	}
	return count
}

// Count fully drains s and returns the number of elements it yielded. It is
// equivalent to Count(s).
func (s Seq[T]) Count() int {
	return Count(s)
}

// Seq2

// Count2 fully drains seq and returns the number of (key, value) pairs it
// yielded.
func Count2[K, V any](seq Seq2[K, V]) int {
	count := 0
	for _ = range seq {
		count++
	}
	return count
}

// Count fully drains s and returns the number of (key, value) pairs it
// yielded. It is equivalent to Count2(s).
func (s Seq2[K, V]) Count() int {
	return Count2(s)
}
