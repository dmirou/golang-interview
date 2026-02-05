package faircandyswap

import "slices"

/*
888. Fair Candy Swap

https://leetcode.com/problems/fair-candy-swap/description/?envType=problem-list-v2&envId=binary-search

Easy
Topics
premium lock icon
Companies

Alice and Bob have a different total number of candies. You are given two integer arrays aliceSizes and bobSizes
where aliceSizes[i] is the number of candies of the ith box of candy that Alice has and bobSizes[j] is the number
of candies of the jth box of candy that Bob has.

Since they are friends, they would like to exchange one candy box each so that after the exchange,
they both have the same total amount of candy. The total amount of candy a person has is the sum
of the number of candies in each box they have.

Return an integer array answer where answer[0] is the number of candies in the box that Alice must exchange,
and answer[1] is the number of candies in the box that Bob must exchange.
If there are multiple answers, you may return any one of them.
It is guaranteed that at least one answer exists.

Example 1:
Input: aliceSizes = [1,1], bobSizes = [2,2]
Output: [1,2]

Example 2:
Input: aliceSizes = [1,2], bobSizes = [2,3]
Output: [1,2]

Example 3:
Input: aliceSizes = [2], bobSizes = [1,3]
Output: [2,3]

Constraints:
1 <= aliceSizes.length, bobSizes.length <= 104
1 <= aliceSizes[i], bobSizes[j] <= 105
Alice and Bob have a different total number of candies.
There will be at least one valid answer for the given input.

Input: aliceSizes = [1,1], bobSizes = [2,2]
aSum = 2, bSum = 4
bobSizes [2, 2]
i target 	j
0 2			0
*/
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	aSum := 0
	for i := 0; i < len(aliceSizes); i++ {
		aSum += aliceSizes[i]
	}

	bSum := 0
	for i := 0; i < len(bobSizes); i++ {
		bSum += bobSizes[i]
	}

	if bSum > aSum {
		slices.Sort(bobSizes)
		for i := 0; i < len(aliceSizes); i++ {
			target := aliceSizes[i] + (bSum-aSum)/2
			if j := binarySearch(bobSizes, target); j != -1 {
				return []int{aliceSizes[i], bobSizes[j]}
			}
		}
	}

	for i := 0; i < len(bobSizes); i++ {
		slices.Sort(aliceSizes)
		target := bobSizes[i] + (aSum-bSum)/2
		if j := binarySearch(aliceSizes, target); j != -1 {
			return []int{aliceSizes[j], bobSizes[i]}
		}
	}

	return nil
}

func binarySearch(arr []int, target int) int {
	l := 0
	r := len(arr) - 1
	c := -1
	for l <= r {
		c = l + (r-l)/2
		if arr[c] == target {
			return c
		}
		if arr[c] < target {
			l = c + 1
		} else {
			r = c - 1
		}
	}
	return -1
}
