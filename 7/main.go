package main

import (
	"fmt"
	"math"
)

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
	/*fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(0))*/
	fmt.Println(reverse(1534236469))
}
