package leetcode1637

import (
	"sort"
)

/**
1637. Widest Vertical Area Between Two Points Containing No Points

Given n points on a 2D plane where points[i] = [xi, yi], Return the widest vertical area between two points such that no points are inside the area.

A vertical area is an area of fixed-width extending infinitely along the y-axis (i.e., infinite height). The widest vertical area is the one with the maximum width.

Note that points on the edge of a vertical area are not considered included in the area.

Input: points = [[8,7],[9,9],[7,4],[9,7]]
Output: 1
Explanation: Both the red and the blue area are optimal.

Example 2:

Input: points = [[3,1],[9,0],[1,0],[1,4],[5,3],[8,8]]
Output: 3


Constraints:

    n == points.length
    2 <= n <= 105
    points[i].length == 2
    0 <= xi, yi <= 109

*/

func maxWidthOfVerticalArea(points [][]int) int {
	ans := 0
	x := make([]int, len(points))
	for i, point := range points {
		x[i] = point[0]
	}
	sort.Ints(x)
	for i := 1; i < len(x); i++ {
		if x[i]-x[i-1] > ans {
			ans = x[i] - x[i-1]
		}
	}
	return ans
}
