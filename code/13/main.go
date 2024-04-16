//  知识点：链表
//  字节跳动,招银网络,百度,网易
//  简单
//  考察次数8
// 题目描述
// 输入两个链表，找出它们的第一个公共结点。（注意因为传入数据是链表，所以错误测试数据的提示是用其他方式显示的，保证传入数据是正确的）

//思路
//第一步：
//     1  2
//           3  4  5
//  7  8  9
//假如两个链表公共结点前面是相等的，就可以用双指针方法
// 0  1  2
//          3  4  5
// 7  8  9
//两个链表依次往后走，直到相等就是第一个公共结点
//如何让公共结点前面相等呢？两个链表相加的长度就会相同，变成下面
// 7  8  9  3  4  5 + 1  2  3  4  5
//
// 1  2  3  4  5 + 7  8  9  3  4  5
//编码时不用真实补充链表
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//模拟测试数据
	//第一个链表前半段
	a1 := []int{1, 2}
	//第二个链表前半段
	a2 := []int{7, 8, 9}
	//公共部分
	a3 := []int{3, 4, 5}
	r1, r2 := createNode(a1, a2, a3)
	//
	//shownode(r1)
	//shownode(r2)

	p := FindFirstCommonNode(r1, r2)
	fmt.Println(p)
}

/**
 *
 * @param pHead1 ListNode类
 * @param pHead2 ListNode类
 * @return ListNode类
 */
func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	if pHead1 == nil || pHead2 == nil {
		return nil
	}

	p1, p2 := pHead1, pHead2

	for p1 != p2 {
		if p1 == nil {
			p1 = pHead2
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = pHead1
		} else {
			p2 = p2.Next
		}
	}
	return p1
}

//数组转换为链表
func createNode(a1 []int, a2 []int, a3 []int) (*ListNode, *ListNode) {
	//生成链表公共部分
	var h3 = new(ListNode)
	h3.Val = a3[0]
	var tail3 *ListNode
	tail3 = h3
	for i := 1; i < len(a3); i++ {
		var listNode = ListNode{Val: a3[i]}
		//设置上一个节点next为本节点的指针
		(*tail3).Next = &listNode
		tail3 = &listNode
	}
	(*tail3).Next = nil

	//生成第一个链表前半段
	var h1 = new(ListNode)
	h1.Val = a1[0]
	var tail1 *ListNode
	tail1 = h1
	for i := 1; i < len(a1); i++ {
		var listNode = ListNode{Val: a1[i]}
		//设置上一个节点next为本节点的指针
		(*tail1).Next = &listNode
		tail1 = &listNode
	}
	(*tail1).Next = h3

	//生成第二个链表前半段
	var h2 = new(ListNode)
	h2.Val = a2[0]
	var tail2 *ListNode
	tail2 = h2
	for i := 1; i < len(a2); i++ {
		var listNode = ListNode{Val: a2[i]}
		//设置上一个节点next为本节点的指针
		(*tail2).Next = &listNode
		tail2 = &listNode
	}
	(*tail2).Next = h3

	return h1, h2
}

func shownode(p *ListNode) { //遍历
	for p != nil {
		fmt.Println(*p)
		p = p.Next //移动指针
	}
}
