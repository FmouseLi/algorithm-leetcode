package main

import "fmt"

/**
抓住题目中的重要信息，XOR 的特点以及题目中 nums 元素的取值范围
利用以上两点和贪心算法思路来逐位求证，最后得到最大的 x 值
*/
func findMaximumXOR(nums []int) (x int) {
	const highBit = 30
	for k := highBit; k >= 0; k-- {
		seen := map[int]bool{}
		for _, num := range nums {
			seen[num>>k] = true
		}

		xNext := x*2 + 1
		found := false

		for _, num := range nums {
			if seen[num>>k^xNext] {
				found = true
				break
			}
		}

		if found {
			x = xNext
		} else {
			x = xNext - 1
		}
	}
	return
}

func main() {
	fmt.Println(findMaximumXOR([]int{3, 10, 5, 25, 2, 8}))
	fmt.Println(findMaximumXOR([]int{0}))
	fmt.Println(findMaximumXOR([]int{2, 4}))
	fmt.Println(findMaximumXOR([]int{8, 10, 2}))
	fmt.Println(findMaximumXOR([]int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}))
}
