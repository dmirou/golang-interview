package main

import "iter"

func SliceAllValues[Slice []V, V any](S Slice) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, val := range S {
			if !yield(val) {
				return
			}
		}
	}
}

func SliceAllKeyValues[Slice []V, V any](S []V) iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		for i, val := range S {
			if !yield(i, val) {
				return
			}
		}
	}
}

func SliceAppendSeq[Slice []V, V any](S Slice, seq iter.Seq[V]) Slice {
	result := make(Slice, len(S), len(S))
	copy(result, S)

	for val := range seq {
		result = append(result, val)
	}

	return result
}
