package main

import "fmt"

/**
滑动窗口思路

窗口的长度尾 k，窗口的取值范围就是 k ～ l - k
*/

func maxScore(cardPoints []int, k int) int {

	l := len(cardPoints)
	if k > l {
		return 0
	}

	m := 0
	for i := 0; i < k; i++ {
		m += cardPoints[i]
	}

	tmp := m
	for i := 0; i < k; i++ {
		tmp = tmp - cardPoints[k-1-i] + cardPoints[l-1-i]
		m = max(m, tmp)
	}

	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3))
	fmt.Println(maxScore([]int{2, 2, 2}, 2))
	fmt.Println(maxScore([]int{9, 7, 7, 9, 7, 7, 9}, 7))
	fmt.Println(maxScore([]int{1, 1000, 1}, 1))
	fmt.Println(maxScore([]int{1, 79, 80, 1, 1, 1, 200, 1}, 3))
}
