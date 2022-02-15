package quick

import (
	"math/rand"
	"time"
)

func Sort(arr []int) {
	rand.Seed(time.Now().UnixNano())
	sort(arr, 0, len(arr)-1)
}

func sort(arr []int, left, right int) {
	if left < right {
		pivotIdx := left + rand.Intn(right-left)
		newPIdx := partition(arr, left, right, pivotIdx)
		sort(arr, left, newPIdx-1)
		sort(arr, newPIdx+1, right)
	}
}

func partition(arr []int, left, right, pivotIdx int) int {
	pV := arr[pivotIdx]
	arr[right], arr[pivotIdx] = arr[pivotIdx], arr[right]
	newPIdx := left
	for i := left; i < right; i++ {
		if arr[i] < pV {
			arr[newPIdx], arr[i] = arr[i], arr[newPIdx]
			newPIdx++
		}
	}
	arr[newPIdx], arr[right] = arr[right], arr[newPIdx]

	return newPIdx
}
