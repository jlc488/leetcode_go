package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	ans := make([]bool, len(candies))
	m := 0
	for _, val := range candies {
		if val > m {
			m = val
		}
	}

	for i := range candies {
		fmt.Println(i)
		ans[i] = (candies[i] + extraCandies) >= m
	}

	return ans
}

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
	//fmt.Println(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1))
	//fmt.Println(kidsWithCandies([]int{12, 1, 12}, 10))

}
