//冒泡排序和优化
package main

import "fmt"

func main() {
	data := []int{1, 6, 7, 3, 9, 2, 4, 0, 5, 8}
	fmt.Println(data)
	data = Bubble(data)
	fmt.Println(data)
}

func Bubble(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] < data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	return data
}
