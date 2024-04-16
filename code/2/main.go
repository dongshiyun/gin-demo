// 2018字节跳动校招
// 简单 穷举
// 题目描述
// 春天是鲜花的季节，水仙花就是其中最迷人的代表，数学上有个水仙花数，他是这样定义的：“水仙花数”是指一个三位数，它的各位数字的立方和等于其本身，比如：153=1^3+5^3+3^3。 现在要求输出所有在m和n范围内的水仙花数。

// 输入描述:
// 输入数据有多组，每组占一行，包括两个整数m和n（100<=m<=n<=999）。
// 输出描述:
// 对于每个测试实例，要求输出所有在给定范围内的水仙花数，就是说，输出的水仙花数必须大于等于m,并且小于等于n，如果有多个，则要求从小到大排列在一行内输出，之间用一个空格隔开;如果给定的范围内不存在水仙花数，则输出no;每个测试实例的输出占一行。
// 示例1
// 输入
// 复制
// 100 120
// 300 380
// 输出
// 复制
// no
// 370 371

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		return
	}
	strList := strings.Fields(input)
	if len(strList) != 2 {
		return
	}
	min, _ := strconv.ParseInt(strList[0], 10, 64)
	max, _ := strconv.ParseInt(strList[1], 10, 64)
	fmt.Println(strList)
	var data []int64
	if min >= 100 && max <= 999 && min <= max {
		for ; min <= max; min++ {
			num1 := min % 10
			num2 := (min / 10) % 10
			num3 := min / 100
			if (num1*num1*num1 + num2*num2*num2 + num3*num3*num3) == min {
				data = append(data, min)
			}
		}
	}
	if len(data) == 0 {
		fmt.Println("no")
	} else {
		for _, v := range data {
			fmt.Printf("%d ", v)
		}
	}
}
