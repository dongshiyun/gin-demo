//  知识点：数学 二分 牛顿法
//  字节跳动,快手,携程,作业帮,美团点评
//  简单
//  考察次数7
// 题目描述
// 实现函数 int sqrt(int x).
// 计算并返回x的平方根
// 输入
// 2
// 输出
// 1

//牛顿法
//其实就是求f(x)=num - x ^ 2的零点
//公式：用代码表示为num = (num + x / num) / 2

package main

import (
	"fmt"
	"math"
)

func main() {
	m := sqrt(100)
	fmt.Println(m)
}

func sqrt(x float64) float64 {
	n := x
	m := n
	for math.Abs(m*m-n) >= 0.00001 {
		fmt.Println(m)
		m = (m + n/m) / 2.0
	}
	return m
}