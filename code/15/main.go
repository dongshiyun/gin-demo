//  知识点：递归
//  字节跳动,作业帮,百度,shopee,面试经典题
//  简单
//  考察次数5

// 题目描述
// 一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法（先后次序不同算不同的结果）。

//思路
//跳台阶 等于 斐波拉契数列 有n个台阶，就是f[n] = f[n-1] + f[n-2].但是时间复杂度呈指数级增长，肯定不是面试的目的
//最优方法
//自底向上循环实现，斐波拉契数列第i的值等于 i-2 加 i-1的值，因为前两个数已经知道
//1 2 3 5 8 13 21 34      8
package main

import "fmt"

func main() {
	fmt.Println(jumpFloor(10))
	fmt.Println(jumpFloor1(10))
}

/**
 *
 * @param number int整型
 * @return int整型
 */
func jumpFloor(number int) int {
	// write code here
	if number <= 1 {
		return 1
	}
	if number == 2 {
		return 2
	}
	return jumpFloor(number-1) + jumpFloor(number-2)
}

func jumpFloor1(number int) int {
	// write code here
	if number <= 1 {
		return 1
	}
	if number == 2 {
		return 2
	}

	a, b := 2, 1
	for i := 3; i <= number; i++ {
		a = a + b
		b = a - b
	}
	return a
}
