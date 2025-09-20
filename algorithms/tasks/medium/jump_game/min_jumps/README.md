Jump Game II
==============
Leetcode https://leetcode.com/problems/jump-game-ii/description/?envType=study-plan-v2&envId=top-interview-150

You are given a 0-indexed array of integers nums of length n. 
You are initially positioned at index 0.

Each element nums[i] represents the maximum length of a forward jump from index i. 
In other words, if you are at index i, you can jump to any index (i + j) where:

0 <= j <= nums[i] and i + j < n

Return the minimum number of jumps to reach index n - 1. 
The test cases are generated such that you can reach index n - 1.

Example 1:
```
Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2. 
Jump 1 step from index 0 to 1, then 3 steps to the last index.
```

Example 2:
```
Input: nums = [2,3,0,1,4]
Output: 2
```

Example 3:
```
Input: nums = [2]
Output: 0
```

Example 4:
```
Input: nums = [3, 1000, 1, 1, 2, 0]
Output: 2
```

Example 5:
```
Input: nums = [3, 2, 3, 1, 1, 0]
Output: 2
```
3 -> j2 -> 3 -> j3 -> 0

Constraints:
```
1 <= nums.length <= 104
0 <= nums[i] <= 1000
It's guaranteed that you can reach nums[n - 1].
```

