//  知识点：字符串，动态规划
//  小米,百度,猿辅导
//  简单
//  考察次数3

// 题目描述
// 对于一个字符串，请设计一个高效算法，计算其中最长回文子串的长度。

// 给定字符串A以及它的长度n，请返回最长回文子串的长度。

// 测试样例：
// "abc1234321ab",12
// 返回：7

//Manacher算法 马拉车算法

package main

import "fmt"

func main() {
	str := "abc1234321ab"
	n := 12
	len := getLongestPalindrome(str, n)
	fmt.Println(len)
}

func getLongestPalindrome(str string, n int) int {
	if n < 2 {
		return 1
	}
	//创建预处理字符串
	strArray := []string{}
	for i := 0; i < len(str); i++ {
		strArray = append(strArray, "#")
		strArray = append(strArray, string(str[i]))
	}
	strArray = append(strArray, "#")
	sLen := 2*n + 1

	// 当前遍历的中心最大扩散步数，其值等于原始字符串的最长回文子串的长度
	maxLen := 1

	for i := 0; i < sLen; i++ {
		iLen := centerLen(strArray, sLen, i)
		if iLen > maxLen {
			maxLen = iLen
		}
	}

	return maxLen
}

func centerLen(strArray []string, sLen int, c int) int {
	n := 0
	left := c - 1
	right := c + 1
	for left >= 0 && right < sLen && strArray[left] == strArray[right] {
		n++
		left--
		right++
	}
	return n
}
