//  知识点：数组 双指针
//  携程,作业帮,快手,字节跳动,拼多多,百度
//  中等
//  考察次数6

// 题目描述
// 给出一个有n个元素的数组S，S中是否有元素a,b,c满足a+b+c=0？找出数组S中所有满足条件的三元组。
// 注意：
// 三元组（a、b、c）中的元素必须按非降序排列。（即a≤b≤c）
// 解集中不能包含重复的三元组。
// 例如，给定的数组 S = {-1 0 1 2 -1 -4},解集为(-1, 0, 1) (-1, -1, 2)

//思路
//1. 求和，指定位置i的数，从i的后一位和数组最后一位向中间闭紧，求出符合条件的值
//2. 去重，首先数组按照升序排序，最外层for循环时，开始下次循环时，只要保证本次循环的值和上一次不同就可以，因为得出的三元数组第一个值肯定不同

package main

import (
	"fmt"
	"sort"
)

func main() {
	num := []int{-1, 0, 1, 2, -1, -4, -3, 3, 4}
	threeNum := threeSum(num)
	fmt.Println(threeNum)
}

/**
 *
 * @param num int整型一维数组
 * @return int整型二维数组
 */
func threeSum(num []int) [][]int {
	// write code here
	//升序，去重的关键
	sort.Ints(num)

	threeNum := make([][]int, 0)

	for i := 0; i < len(num); {
		//取第i个为第一个数
		//第二个数left为i+1
		//第三个数right为len(num)-1
		left := i + 1
		right := len(num) - 1
		for left < right {
			if num[left]+num[right]+num[i] == 0 {
				//加入到三元数组中
				tmp := make([]int, 0)
				tmp = append(tmp, num[i])
				tmp = append(tmp, num[left])
				tmp = append(tmp, num[right])
				threeNum = append(threeNum, tmp)

				left++
				//假如下一个数相同，越过找下一个
				for num[left] == num[left-1] {
					left++
				}

				right--
				//假如下一个数相同，越过找下一个
				for num[right] == num[right+1] {
					right--
				}

			} else if num[left]+num[right]+num[i] > 0 {
				right--
				//假如下一个数相同，越过找下一个
				for num[right] == num[right+1] {
					right--
				}
			} else if num[left]+num[right]+num[i] < 0 {
				left++
				//假如下一个数相同，越过找下一个
				for num[left] == num[left-1] {
					left++
				}
			}
		}

		i++
		for i < len(num) && num[i] == num[i-1] {
			i++
		}
	}

	return threeNum
}
