package bubble

func Sort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	var swapped bool

	for i := 0; i < len(arr); i++ {
		swapped = false

		for j := 1; j < len(arr)-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				swapped = true
			}
		}

		if !swapped {
			return
		}
	}
}
