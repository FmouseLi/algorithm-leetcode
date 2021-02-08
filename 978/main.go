package main

import "fmt"

/**

此题自己暴力解出来，代码和思路不清晰，官方题解比较清晰易懂

动态规划

记 dp[i][0] 为以 arr[i] 结尾，且 arr[i−1] < arr[i] 的「湍流子数组」的最大长度；
dp[i][1] 为以 arr[i] 结尾，且 arr[i−1] > arr[i] 的「湍流子数组」的最大长度。

显然，以下标 0 结尾的「湍流子数组」的最大长度为 1，因此边界情况为 dp[0][0] = dp[0][1]=1。

当 i>0 时，考虑 arr[i−1] 和 arr[i] 之间的大小关系：

如果 arr[i−1] > arr[i]，则如果以下标 i−1 结尾的子数组是「湍流子数组」，应满足 i − 1 = 0，或者当 i − 1 > 0 时 arr[i−2] < arr[i−1]，因此 dp[i][0] = dp[i−1][1] + 1，dp[i][1] = 1；
如果 arr[i−1] < arr[i]，则如果以下标 i−1 结尾的子数组是「湍流子数组」，应满足 i − 1 = 0，或者当 i − 1 > 0 时 arr[i−2] > arr[i−1]，因此 dp[i][0] = 1，dp[i][1] = dp[i−1][0] + 1；
如果 arr[i−1] = arr[i]，则 arr[i−1] 和 arr[i] 不能同时出现在同一个湍流子数组中，因此 dp[i][0] = dp[i][1] = 1。

用两个变量 dp0 和 dp1 代替 dp[i][0] 和 dp[i][1]，将空间复杂度降到 O(1)。

*/

func maxTurbulenceSize(arr []int) int {
	ans := 1
	n := len(arr)
	dp0, dp1 := 1, 1
	for i := 1; i < n; i++ {
		if arr[i-1] > arr[i] {
			dp0, dp1 = dp1+1, 1
		} else if arr[i-1] < arr[i] {
			dp0, dp1 = 1, dp0+1
		} else {
			dp0, dp1 = 1, 1
		}
		ans = max(ans, max(dp0, dp1))
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxTurbulenceSize([]int{9, 4, 2, 10, 7, 8, 8, 1, 9}))
	fmt.Println(maxTurbulenceSize([]int{4, 8, 12, 16}))
	fmt.Println(maxTurbulenceSize([]int{100}))
	fmt.Println(maxTurbulenceSize([]int{9, 9}))
}
