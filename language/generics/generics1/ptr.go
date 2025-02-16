package main

func Ptr[T any](value T) *T {
	return &value
}

func Val[T any](value *T) T {
	if value == nil {
		return *new(T)
	}

	return *value
}

func Copy[T any](value *T) *T {
	if value == nil {
		return nil
	}
	clone := *value

	return &clone
}
