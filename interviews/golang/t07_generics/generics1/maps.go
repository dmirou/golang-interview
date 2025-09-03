// PREPARE: Copy TestMapCompare func from maps_test.go to the candidate.
// GIVE TASK: Write a function MapCompare which compares maps with different types.
package generics1

func MapCompare[K, V comparable](m1, m2 map[K]V) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; !ok || v2 != v1 {
			return false
		}
	}

	return true
}

func MapApplyFunc[K comparable, V any](m map[K]V, fn func(item V) V) {
	for k, v := range m {
		m[k] = fn(v)
	}
}
