//2020-8-21 滴滴笔试
// 小明昨晚做了一个梦。在梦里，很多很多斐波那契数连成了一条蛇。突然，最大的那个数变成了蛇头，把小明一口给吞到肚子里去了。

// 小明被吓醒了，他赶紧拿笔在纸上面画了一条斐波那契蛇。

// 这是一个蛇形迂回的斐波那契数列，它是一个n*n的矩阵，在上面的矩阵中n=3。第1行第1列是最大值，然后按照顺时针的次序数字逐渐变小。

// 下面是n=4时的情况：

// 小明希望你能够编写一个程序，输入一个正整数n，然后逐行逐列输出斐波那契蛇形矩阵中的元素。

// 输入描述
// 单组输入，输入数据占一行，包含一个正整数n，表示斐波那契蛇形矩阵的大小。(n<10)

// 输出描述
// 输出数据占一行，逐行逐列（从第1行开始到第n行，每一行从第1列开始到第n列）输出斐波那契蛇形矩阵中的元素，每两个数字之间用一个空格隔开。

// 样例输入
// 3
// 样例输出
// 34 21 13
// 1 1 8
// 2 3 5
//斐波那契数

// 0*0 0*1 0*2 0*3
// 1*0 1*1 1*2 1*3
// 2*0 2*1 2*2 2*3
// 3*0 3*1 3*2 3*3

// 987 610 377 233
// 5   3   2   144
// 8   1   1   89
// 13  21  34  55
package main

import "fmt"

func main() {
	n := 4
	//根据n获取斐波那契数
	data := Feibo(n)
	//创建并初始化切片
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
	}

	One(&a, &data, n*n-1, 0, n-1)

	//输出到命令行
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(a[i][j], " ")
		}
		fmt.Printf("\n")
	}
}

func Feibo(n int) []int {
	data := []int{}
	for i := 0; i <= n*n-1; i++ {
		if i == 0 || i == 1 {
			data = append(data, 1)
		} else {
			data = append(data, data[i-1]+data[i-2])
		}
	}
	return data
}

func One(a *[][]int, data *[]int, index int, left int, right int) {
	//从左到右 0*0 0*1 0*2 0*3 1*1 1*2
	for i := left; i <= right; i++ {
		(*a)[left][i] = (*data)[index]
		index--
	}
	//从上到下 0*3 1*3 2*3 3*3    1*2 2*2
	for i := left + 1; i <= right; i++ {
		(*a)[i][right] = (*data)[index]
		index--
	}
	//从右到左 3*3 3*2 3*1 3*0    2*2 2*1
	for i := right - 1; i >= left; i-- {
		(*a)[right][i] = (*data)[index]
		index--
	}
	//从下到上 3*0 2*0 1*0 0*0
	for i := right - 1; i > left; i-- {
		(*a)[i][left] = (*data)[index]
		index--
	}
	left++
	if left < right {
		right--
		One(a, data, index, left, right)
	}
}
