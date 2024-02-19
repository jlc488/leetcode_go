package leetcode771

/**
771. Jewels and Stones

You're given strings jewels representing the types of stones that are jewels, and stones representing the stones you have. Each character in stones is a type of stone you have. You want to know how many of the stones you have are also jewels.

Letters are case sensitive, so "a" is considered a different type of stone from "A".

Example 1:

Input: jewels = "aA", stones = "aAAbbbb"
Output: 3
Example 2:

Input: jewels = "z", stones = "ZZ"
Output: 0

Constraints:

1 <= jewels.length, stones.length <= 50
jewels and stones consist of only English letters.
All the characters of jewels are unique.
*/

// O(n) time complexity
func numJewelsInStones(jewels string, stones string) int {
	jMap := make(map[rune]int)

	for _, v := range jewels {
		jMap[v] += 1
	}
	ans := 0
	for _, v := range stones {
		if numberOfJewel, ok := jMap[v]; ok {
			ans += numberOfJewel
		}
	}
	return ans
}

// O(n^2) time complexity
//func numJewelsInStones(jewels string, stones string) int {
//	counter := 0
//
//	for i := 0; i < len(jewels); i++ {
//		for j := 0; j < len(stones); j++ {
//			if jewels[i] == stones[j] {
//				counter++
//			}
//		}
//	}
//
//	return counter
//}
