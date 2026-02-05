package kthmissingnumber

/*
1539. Kth Missing Positive Number

https://leetcode.com/problems/kth-missing-positive-number/?envType=problem-list-v2&envId=binary-search

Easy
Topics
premium lock icon
Companies
Hint

Given an array arr of positive integers sorted in a strictly increasing order, and an integer k.

Return the kth positive integer that is missing from this array.

Example 1:

Input: arr = [2,3,4,7,11], k = 5
Output: 9
Explanation: The missing positive integers are [1,5,6,8,9,10,12,13,...]. The 5th missing positive integer is 9.

Example 2:
Input: arr = [1,2,3,4], k = 2
Output: 6
Explanation: The missing positive integers are [5,6,7,...]. The 2nd missing positive integer is 6.

Constraints:

1 <= arr.length <= 1000
1 <= arr[i] <= 1000
1 <= k <= 1000
arr[i] < arr[j] for 1 <= i < j <= arr.length

Follow up:
Could you solve this problem in less than O(n) complexity?

Input: arr = [2,3,4,7,11], k = 5
Output: 9
l r mid arr[mid] mc poss pmc
0 4 2   4		 1  2    1
3 4 3   7 		 3  3    3
4 4 4	11		 6  3    3

ans = 7 + 5 - 3 = 9

input: arr = [2] k = 1
1 < 2, ans k = 1

input: arr = [1 3] k = 1
l r mid arr[mid] mc poss pmc
0 1 0	1		 0  0	 0
1 1 1   3		 1  0    0
1 + 1 - 0 = 2
*/
func findKthPositive(arr []int, k int) int {
	if k < arr[0] {
		return k
	}
	l := 0
	r := len(arr) - 1
	mid := -1
	possible := -1
	pMissCount := -1
	for l <= r {
		mid = l + (r-l)/2
		missCount := arr[mid] - (mid + 1)
		if missCount < k {
			possible = mid
			pMissCount = missCount
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if possible == -1 {
		return k
	}
	return arr[possible] + k - pMissCount
}
