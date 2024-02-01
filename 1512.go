package main

/**
Given an array of integers nums, return the number of good pairs.

A pair (i, j) is called good if nums[i] == nums[j] and i < j.

Example 1:

Input: nums = [1,2,3,1,1,3]
Output: 4
Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.

Example 2:

Input: nums = [1,1,1,1]
Output: 6
Explanation: Each pair in the array are good.

Example 3:

Input: nums = [1,2,3]
Output: 0

Constraints:

    1 <= nums.length <= 100
    1 <= nums[i] <= 100

*/

import "fmt"

func numIdenticalPairs2(nums []int) int {
	ans := 0

	for i, val1 := range nums {
		for j, val2 := range nums {
			if val1 == val2 && i < j {
				ans++
			}
		}
	}

	return ans
}

func numIdenticalPairs1(nums []int) int {
	counts := make([]int, 100)
	ans := 0

	for _, val := range nums {
		counts[val-1]++
	}

	for _, count := range counts {
		if count > 1 {
			ans += (count * (count - 1)) / 2
		}
	}

	return ans
}

func main() {
	fmt.Println(numIdenticalPairs1([]int{1, 2, 3, 1, 1, 3}))
	fmt.Println(numIdenticalPairs2([]int{1, 2, 3, 1, 1, 3}))
}
