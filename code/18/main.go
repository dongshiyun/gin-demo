//  知识点：双指针 字符串
//  腾讯,字节跳动
//  简单
//  考察次数2

// 题目描述
// 写出一个程序，接受一个字符串，然后输出该字符串反转后的字符串。（字符串长度不超过1000）
// 示例1
// 输入
// "abcd"
// 输出
// "dcba"
package main

import "fmt"

func main() {
	s := "abcd"
	s1 := solve(s)
	fmt.Print(s1)
}

/**
 * 反转字符串
 * @param str string字符串
 * @return string字符串
 */
func solve(str string) string {
	// write code here
	left := 0
	right := len(str) - 1
	resStr := []byte(str)
	for left < right {
		resStr[left], resStr[right] = resStr[right], resStr[left]
		left++
		right--
	}
	return string(resStr)
}
