package sumoddlensubarrays

// arr = [1]
// sums = [1]
// sum = [1]
//
// arr = [1, 2]
// sums = [1, 3]
// sum = [1] + [2]
// sum = sums[n-1]
//
// arr = [1, 2, 3]
// sums = [1, 3, 6]
// sum = [1] + [2] + [3] + [1 + 2 + 3]
// sum = sums[n-1] + sums[n-1]
//
// arr = [1, 2, 3, 4]
// sums = [1, 3, 6, 10]
// sum = [1] + [2] + [3] + [1 + 2 + 3] + [4] + [2 + 3 + 4]
// sum = sums[n-1] + sums[2] + sums[3] - sums[1]
//

// arr = [1, 2, 3, 4, 5]
// sums = [1, 3, 6, 10, 15]
// sum = [1] + [2] + [3] + [4] + [5] + [1, 2, 3] + [2, 3, 4] + [3, 4, 5] + [1, 2, 3, 4, 5]
// sum = [1] + [2] + [3] + [1, 2, 3] + [4] + [2, 3, 4] + [5] + [1, 2, 3, 4, 5] + [3, 4, 5]

// arr = [1, 2, 3, 4, 5, 6, 7]
// sum = [1] + [2] + [3] + [1 2 3] + [4] + ([1 2 3 4] - [1]) + [5] + [1 2 3 4 5] + ([[1 2 3 4 5] - [1 2])
// + [6] + ([1 2 3 4 5 6] - [1]) + ([1 2 3 4 5 6] - [1 2 3]) + [7] + [1 2 3 4 5 6 7] + ([1 2 3 4 5 6 7] - [1 2])
// + ([1 2 3 4 5 6 7] - [1 2 3 4])

// Input: arr = [1,4,2,5,3]
// Output: 58
// Explanation: The odd-length subarrays of arr and their sums are:
//
// [1] + [4] + [2] + [1,4,2] + [5] + [4,2,5] + [3] + [2,5,3] + [1,4,2,5,3]
//
// If we add all these together we get 1 + 4 + 2 + 5 + 3 + 7 + 11 + 10 + 15 = 58

// Constraints:
// 1 <= arr.length <= 100
// 1 <= arr[i] <= 1000
func sumOddLengthSubarrays(arr []int) int {
	sums := make([]int, len(arr))
	sums[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		sums[i] = sums[i-1] + arr[i]
	}

	// [1] 1
	// ts 1
	//
	// [1 2]
	// ts 1
	// i 1, i < 2
	// j 0, j < 0
	// ts 1

	/*
		[1 2 3 4 5 6 7]
		[1 3 6 10 15 21 28]
		ts 1
		i 1
			ts = 1 + 2 = 3
			j 0
			0 < 0
			ts exp 1+2 = 3
		i 2
			ts = 3 + 3 = 6
			j 0
			j 0 < 1
			2 % 2 = 0
			j 1 < 1
			ts = 6 + 6 = 12
			ts exp 1+2+3 + [1 2 3] = 12
		i 3
			ts = 12 + 4 = 16
			j 0 < 2
			3 % 2 = 1, ts = 16 + 10 - 1 = 25
			j 1 < 2
			2 % 2 = 0
			j 2 < 2
			i % 2 = 1
			ts exp = 12 + 4 + [2 3 4] = 16 + 9 = 25
		i 4
		...
	*/
	ts := sums[0]
	for i := 1; i < len(arr); i++ {
		ts += arr[i]
		for j := 0; j < i-1; j++ {
			if (i-j)%2 == 1 {
				ts += sums[i] - sums[j]
			}
		}
		// i starts from 0, so odd subarray will be with i 0, 2, 4, 6
		if i%2 == 0 {
			ts += sums[i]
		}
	}

	return ts
}
