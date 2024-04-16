//  知识点：数组、动态规划、贪心
//  字节跳动,远景能源
//  简单
//  考察次数3

// 题目描述
// 假设你有一个数组，其中第 i 个元素是股票在第 i 天的价格。
// 你有一次买入和卖出的机会。（只有买入了股票以后才能卖出）。请你设计一个算法来计算可以获得的最大收益。
// 示例1
// 输入
// [1,4,2]
// 输出
// 3
// 示例2
// 输入
// [2,4,1]
// 输出
// 2

//贪心算法，从左往右遍历，遇到最小值，保存，
//如果不是最小值，则计算差距，判断保存最大利润

package main

import "fmt"

func main() {
	a1 := []int{1, 4, 2}
	a2 := []int{2, 4, 1}
	max1 := maxProfit(a1)
	max2 := maxProfit(a2)
	fmt.Println(max1, max2)
}

/**
 *
 * @param prices int整型一维数组
 * @return int整型
 */
func maxProfit(prices []int) int {
	// write code here
	min := prices[0]
	max := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			if prices[i]-min > max {
				max = prices[i] - min
			}
		}
	}
	return max
}
