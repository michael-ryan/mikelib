package streams

// Map applies a function f to all members of xs, then returns a new slice containing the mapped values.
func Map[T, U any](xs []T, f func(T) U) []U {
	ys := make([]U, len(xs))

	for i, x := range xs {
		ys[i] = f(x)
	}

	return ys
}

// Filter applies a predicate to all members of xs, then returns a new slice containing only the members of xs such that predicate(x) = true.
func Filter[T any](xs []T, predicate func(T) bool) []T {
	ys := make([]T, 0)

	for _, x := range xs {
		if predicate(x) {
			ys = append(ys, x)
		}
	}

	return ys
}

// ToGenerator returns a channel that will produce all members of xs before closing.
// The element at xs[0] will be the first element sent through the channel.
//
// This function performs the inverse operation of [Collect].
func ToGenerator[T any](xs []T) <-chan T {
	generator := make(chan T)

	go func(gen chan<- T, xs []T) {
		for _, x := range xs {
			gen <- x
		}
		close(gen)
	}(generator, xs)

	return generator
}

// Collect returns a slice that contains all values produced by xs before closing.
// The first element produced by the channel will be stored in index 0 of the slice.
//
// This function performs the inverse operation of [ToGenerator].
//
// Note that if xs never closes, this function will block indefinitely.
func Collect[T any](xs <-chan T) []T {
	collected := make([]T, 0)

	for x := range xs {
		collected = append(collected, x)
	}

	return collected
}

// Fold collapses a slice into a single value, using a given folder function.
//
// The folder function should take two arguments: the accumulator and the current value.
func Fold[T, U any](xs []T, folder func(U, T) U, initialValue U) U {
	for _, x := range xs {
		initialValue = folder(initialValue, x)
	}
	return initialValue
}

// Foreach applies a function f to all members of xs.
func Foreach[T any](xs []T, f func(T)) {
	for _, x := range xs {
		f(x)
	}
}
