package main

import "fmt"

/**
滑动窗口的思想

依次遍历字符串 s，维护当前消耗 curCost，将当前消耗 curCost 与最大消耗 maxCost 相比
	如果 curCost > maxCost，left 向右移动，此时 left 到 right 的长度保持不变，并且维护 curCost 的值
	如果 curCost <= maxCost，left 不移动，维护 curCost 的值

这里有可能是中间的某个子串是 s 和 t 的最长匹配字符串，这里继续上面的逻辑说明下：
随着 left 不断右移，left 的消耗 > right 的消耗时，每次加 right 的消耗（小值），减去 left 的消耗（大值），所以curCost 有可能会逐渐变小
	当 curCost < maxCost 时，继续遍历，curCost + right 的消耗 < maxCost 时，left 没有移动，right 移动，这时候 left 到 right 的长度会变大

遍历完成后，left 到 right 的长度为最大匹配字符串长度
*/

func equalSubstring(s string, t string, maxCost int) int {
	left, curCost := 0, 0
	for right := range s {
		curCost += Abs(int(s[right]) - int(t[right]))
		if curCost > maxCost {
			curCost -= Abs(int(s[left]) - int(t[left]))
			left++
		}
	}
	return len(s) - left
}

//Abs 绝对值
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(equalSubstring("abcd", "bcdf", 3))
	fmt.Println(equalSubstring("abcd", "cdef", 3))
	fmt.Println(equalSubstring("abcd", "acde", 0))
	fmt.Println(equalSubstring("krrgw", "zjxss", 19))
}
