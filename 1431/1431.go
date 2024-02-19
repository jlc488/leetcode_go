package leetcode1431

func kidsWithCandies(candies []int, extraCandies int) []bool {
	ans := make([]bool, len(candies))
	m := 0
	for _, val := range candies {
		if val > m {
			m = val
		}
	}

	for i := range candies {
		ans[i] = (candies[i] + extraCandies) >= m
	}

	return ans
}
