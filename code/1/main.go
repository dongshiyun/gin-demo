//2017滴滴校招
//简单 知识点：贪心
//题目描述
//一个数组有 N 个元素，求连续子数组的最大和。 例如：[-1,2,1]，和最大的连续子数组为[2,1]，其和为 3
//输入描述:
//输入为两行。 第一行一个整数n(1 <= n <= 100000)，表示一共有n个元素 第二行为n个数，即每个元素,每个整数都在32位int范围内。以空格分隔。
//输出描述:
//所有连续子数组中和最大的值。
//输入
//3 -1 2 1
//输出
//3

package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	data := make([]int, n)
	for i := range data {
		fmt.Scan(&data[i])
	}
	res, now := data[0], 0
	for i := range data {
		if now < 0 {
			now = data[i]
		} else {
			now += data[i]
		}
		if now > res {
			res = now
		}
	}
	fmt.Println(res)
}
