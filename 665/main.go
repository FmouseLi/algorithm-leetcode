package main

import "fmt"

/**
贪心算法
依次遍历 nums 数组，如果出现 nums[i] > nums[i+1]，需要修改一个元素，修改 nums[i] 或 nums[i+1] 都可以
但是当 nums[i+1] < nums[i-1] 时，必须要修改 nums[i+1]
因为由 nums[i+1] < nums[i-1] <= nums[i] 变成 nums[i-1] <= nums[i] <= nums[i+1] 必须修改 nums[i+1]
其他情况修改 nums[i] 来满足条件，这样可以尽量减少对后续元素的影响，减少修改次数，也是题目要求最少修改
*/

func checkPossibility(nums []int) bool {
	cnt := 0
	for i := 0; i < len(nums)-1; i++ {
		x, y := nums[i], nums[i+1]
		if x > y {
			cnt++
			if cnt > 1 {
				return false
			}
			if i > 0 && y < nums[i-1] {
				nums[i+1] = x
			}
		}
	}
	return true
}

func main() {
	fmt.Println(checkPossibility([]int{4, 2, 3}))
	fmt.Println(checkPossibility([]int{4, 2, 1}))
	fmt.Println(checkPossibility([]int{3, 4, 2, 3}))
}
