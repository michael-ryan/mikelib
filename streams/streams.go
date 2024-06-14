package streams

func Map[T, U any](xs []T, f func(T) U) []U {
	ys := make([]U, len(xs))

	for i, x := range xs {
		ys[i] = f(x)
	}

	return ys
}

func Filter[T any](xs []T, predicate func(T) bool) []T {
	var ys []T

	for _, x := range xs {
		if predicate(x) {
			ys = append(ys, x)
		}
	}

	return ys
}

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

func Collect[T any](xs <-chan T) []T {
	var collected []T

	for x := range xs {
		collected = append(collected, x)
	}

	return collected
}
