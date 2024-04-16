//知识点，堆 分治
//中等
//字节跳动,猿辅导,作业帮,利库软件,招银网络,百度,远景能源
//考察次数12

// 快速排序基本思想
// 1．先从数列中取出一个数作为基准数。
// 2．分区过程，将比这个数大的数全放到它的右边，小于或等于它的数全放到它的左边。
// 3．再对左右区间重复第二步，直到各区间只有一个数。

// 题目描述
// 有一个整数数组，请你根据快速排序的思路，找出数组中第K大的数。

// 给定一个整数数组a,同时给定它的大小n和要找的K(K在1到n之间)，请返回第K大的数，保证答案存在。

// 测试样例：
// [1,3,5,2,2],5,3
// 返回：2
package main

import "fmt"

func main() {
	a := []int{8, 3, 5, 2, 1, 9, 7}
	n := 7
	k := 3
	fmt.Println(findKth(a, n, k))
}

/**
 *
 * @param a int整型一维数组
 * @param n int整型
 * @param K int整型
 * @return int整型
 */
func findKth(a []int, n int, k int) int {
	// write code here
	left, right := []int{}, []int{}
	v := a[0]

	for i := 1; i < n; i++ {
		if a[i] < v {
			left = append(left, a[i])
		} else {
			right = append(right, a[i])
		}
	}
	fmt.Println("left=", left)
	fmt.Println("right=", right)
	//假如right(大数)的值等于k-1返回v
	if len(right) == k-1 {
		return v
	}
	if len(right) > k-1 {
		return findKth(right, len(right), k)
	} else {
		return findKth(left, len(left), k-len(right)-1) //减一是为了减去自己
	}
}
