package main

import (
	"fmt"
	"math"
	"strings"
)

/**
自己写的有点糟糕，各种条件判断和分之。

下面是题解中看到的，代码和思路清晰简单

1 去除字符的空格
2 然后字符遍历 如果数字就相加
3 根据首个字符是否是+- 符号
4 最大值判断

PS：题目中并没有说环境中是否允许超出 int32 范围。如果环境不允许，这里需要多几行代码来提前判断是否会超出范围。
*/

func myAtoi(str string) int {
	//去掉收尾空格
	str = strings.TrimSpace(str)
	result := 0
	sign := 1

	for i, v := range str {
		if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
		} else if v == '-' && i == 0 {
			sign = -1
		} else if v == '+' && i == 0 {
			sign = 1
		} else {
			break
		}
		// 数值最大检测
		if result > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}

	return sign * result
}

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("21474836460"))
}
