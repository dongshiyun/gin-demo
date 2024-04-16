//简单
//百度,字节跳动,腾讯,京东,小米,携程,作业帮,一点资讯,网易,拼多多,美团点评,招银网络,南北软件
//考察次数31
// 题目描述
// 请实现有重复数字的有序数组的二分查找。
// 输出在数组中第一个大于等于查找值的位置，如果数组中不存在这样的数，则输出数组长度加一。

// 示例1
// 输入
// 5,4,[1,2,4,4,5]
// 输出
// 3
package main

import (
	"fmt"
)

func main() {
	n := 9
	v := 0
	a := []int{0, 1, 2, 3, 3, 4, 5, 7, 9}
	fmt.Println(upper_bound_(n, v, a))
}

/**
 * 二分查找
 * @param n int整型 数组长度
 * @param v int整型 查找值
 * @param a int整型一维数组 有序数组
 * @return int整型
 */
func upper_bound_(n int, v int, a []int) int {
	// write code here
	left, right := 0, n-1
	for left < right {
		mid := (left + right) / 2
		if a[mid] >= v {
			if mid == 0 || a[mid-1] < v {
				return mid + 1
			} else {
				right = mid
			}
		} else {
			left = mid + 1
		}
	}
	return n + 1
}
