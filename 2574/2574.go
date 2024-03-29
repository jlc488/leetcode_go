package leetcode2574

/**
2574. Left and Right Differences

Given a 0-indexed integer array nums, find a 0-indexed integer array answer where:

    answer.length == nums.length.
    answer[i] = |leftSum[i] - rightSum[i]|.

Where:

    leftSum[i] is the sum of elements to the left of the index i in the array nums. If there is no such element, leftSum[i] = 0.
    rightSum[i] is the sum of elements to the right of the index i in the array nums. If there is no such element, rightSum[i] = 0.

Return the array answer.



Example 1:

Input: nums = [10,4,8,3]
Output: [15,1,11,22]
Explanation: The array leftSum is [0,10,14,22] and the array rightSum is [15,11,3,0].
The array answer is [|0 - 15|,|10 - 11|,|14 - 3|,|22 - 0|] = [15,1,11,22].

Example 2:

Input: nums = [1]
Output: [0]
Explanation: The array leftSum is [0] and the array rightSum is [0].
The array answer is [|0 - 0|] = [0].



Constraints:

    1 <= nums.length <= 1000
    1 <= nums[i] <= 105


*/

// TODO need to refactor for time complexity
func leftRightDifference(nums []int) []int {
	answer := make([]int, len(nums))

	leftSum := make([]int, len(nums))
	rightSum := make([]int, len(nums))

	leftSum[0] = 0
	for i := 1; i < len(nums); i++ {
		leftSum[i] = leftSum[i-1] + nums[i-1]
	}

	rightSum[len(nums)-1] = 0
	for j := len(nums) - 2; j >= 0; j-- {
		rightSum[j] = rightSum[j+1] + nums[j+1]
	}

	for i := 0; i < len(nums); i++ {
		answer[i] = abs(leftSum[i] - rightSum[i])
	}

	return answer
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
