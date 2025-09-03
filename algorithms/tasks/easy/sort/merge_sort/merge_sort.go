package main

func mergeSort(s []int) {
	buf := make([]int, len(s))
	sortRecursive(s, 0, len(s), buf)
	for i := 0; i < len(s); i++ {
		s[i] = buf[i]
	}
}

func sortRecursive(s []int, l, r int, buf []int) {
	if r-l <= 1 {
		return
	}
	m := (l + r) / 2
	sortRecursive(s, l, m, buf)
	sortRecursive(s, m, r, buf)
	mergeSorted(s, l, m, r, buf)
}

func mergeSorted(s []int, l, m, r int, buf []int) {
	i1, i2, ri := l, m, 0

	for i1 < m || i2 < r {
		if i1 == m {
			for i := i2; i < r; i++ {
				buf[ri] = s[i]
				ri++
			}
			break
		}
		if i2 == r {
			for i := i1; i < m; i++ {
				buf[ri] = s[i]
				ri++
			}
			break
		}

		if s[i1] <= s[i2] {
			buf[ri] = s[i1]
			ri++
			i1++
			continue
		}
		buf[ri] = s[i2]
		ri++
		i2++
	}

	for i := l; i < r; i++ {
		s[i] = buf[i-l]
	}
}
