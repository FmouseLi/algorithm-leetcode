package main

import "fmt"

/**
首先想到的是能不能更进一步拆解，可以看到 Z 字型的行和列是相等的，其次想到的是一直重复一竖一斜，所以以一竖一斜为单位来解此题

一竖一斜使用了多少个字母呢，2*numRows-2 个，用变量 m 表示

可以看到第一行和最后一行其实比较简单，是有规律的
第一行从 0 开始，间隔 2*numRows-2 出现一个
最后一行从 numRows-1 开始，间隔 2*numRows-2 出现一个

中间行有两个字母，其中一竖上的字母和上面第一行和最后一行的规律一样，现在来看斜线上的字母规律
第二行斜线上的字母，要从最小的一竖一斜重复单位入手考虑，可以看到第二行斜线上的字母是最小单位的倒数第二个
假设当前重复一竖一斜是从 s 的第 j 个字母开始的，那么第二行斜线上的字母下标就是 j+(2*numRows-2)-1，
第三行斜线上的字母下标就是 j+(2*numRows-2)-2，变动的和其所在行有关。所以只要不是第一行和最后一行，需要额外处理这个字母
*/

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	l, m, z := len(s), 2*numRows-2, 0
	buffer := make([]byte, l)
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < l; j += m {
			buffer[z] = s[i+j]
			z++
			if i != 0 && i != numRows-1 && j+m-i < l {
				buffer[z] = s[j+m-i]
				z++
			}
		}
	}
	return string(buffer)
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("A", 1))
}
