package main

// написать функцию,  которая на вход принимает массив целых чисел
// и возвращает индекс числа, без которого произведение будет максимальным.
//
// [2,1,5,3,4] -> 1, т. к. 2*5*3*4 - максимально
// [5, -6, 3, -2] ->  2
// [5, -6, 0, -2, 0] ->  2

func minIndex(items []int) int {
	if len(items) == 1 {
		return 0
	}

	nNeg := 0
	minPos := items[0]
	maxNeg := 0
	zeroIdx := -1
	hasZero := false

	for i, itm := range items {
		if itm < 0 {
			nNeg++

			if maxNeg < itm {
				maxNeg = itm
			}

			continue
		}
		if itm == 0 {
			hasZero = true
			zeroIdx = i
			continue
		}

		// itm > 0
		if minPos > itm {
			minPos = itm
		}
	}

	// [5, -6, 3, -2, 0] ->  4
	if nNeg > 0 && nNeg%2 == 0 && hasZero {
		return zeroIdx
	}

	// [5, -6, 3, -2, -1, 0] ->  4
	// [5, -6, 3, -2, -1] ->  4
	if nNeg > 0 && nNeg%2 == 0 {
		return maxNeg
	}

	if hasZero {
		return zeroIdx
	}

	return minPos
}

func main() {

}
