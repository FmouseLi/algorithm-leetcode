package main

import (
	"fmt"
	"math"
)

/**
因为题目中假设环境不允许存储 64 位整数，所以每次先检查是否会超出 int32 的范围
math.MinInt32 = -2147483648
math.MaxInt32 = 2147483647
因为在计算过程中是将 x 一位一位拆解再计算的，所以将 MaxInt32 和 MinInt32 也分成两部分，这样就可以预判是否会超出 int32 的范围
*/

func reverse(x int) int {

	ans := 0
	min := math.MinInt32 / 10
	max := math.MaxInt32 / 10

	for x != 0 {

		m := x % 10

		if ans > max || (ans == max && m > 7) {
			return 0
		}

		if ans < min || (ans == min && m < -8) {
			return 0
		}

		ans = ans*10 + m
		x = x / 10
	}

	return ans
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(0))
	fmt.Println(reverse(1534236469))
}
