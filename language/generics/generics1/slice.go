package main

type Number interface {
	int | int64 | float64
}

func SliceSum[T Number](values []T) T {
	var sum T
	for _, v := range values {
		sum += v
	}

	return sum
}

func SliceApplyFuncInMemory[T any](values []T, fn func(src *T)) {
	for i := range values {
		fn(&values[i])
	}
}

func SliceApplyFunc[T any](values []T, fn func(src T) T) []T {
	var results []T

	for i := range values {
		results = append(results, fn(values[i]))
	}
	return results
}
