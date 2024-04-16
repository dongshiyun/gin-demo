//知识点，链表 双指针
//中等
//美团点评,shopee,猿辅导,拼多多,百度,腾讯,字节跳动
//考察次数10
// 题目描述
// 给定一个链表，删除链表的倒数第n个节点并返回链表的头指针
// 例如，
//  给出的链表为:1->2->3->4->5, n= 2.
//  删除了链表的倒数第n个节点之后,链表变为1->2->3->5.
// 备注：
// 题目保证n一定是有效的
// 请给出请给出时间复杂度为O(n)的算法

// 示例1
// 输入
// {1,2},2
// 输出
// {2}

//先要理解什么是链表
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//初始化一个链表
	head := createNode()
	//遍历打印结果
	shownode(head)

	//删除
	removeNthFromEnd(head, 3)

	shownode(head)
}

func createNode() *ListNode {
	//初始化一个链表
	var head = new(ListNode)
	head.Val = 0
	var tail *ListNode
	tail = head //tail用于记录最末尾的节点的地址，刚开始tail的的指针指向头节点
	for i := 1; i < 10; i++ {
		var listNode = ListNode{Val: i}
		//设置上一个节点next为本节点的指针
		(*tail).Next = &listNode
		tail = &listNode
	}
	return head
}

func shownode(p *ListNode) { //遍历
	for p != nil {
		fmt.Println(*p)
		p = p.Next //移动指针
	}
}

/**
 *
 * @param head ListNode类
 * @param n int整型
 * @return ListNode类
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// write code here
	//设置双指针
	one := head
	two := head
	//one先走N步
	for i := 0; i < n; i++ {
		one = one.Next
	}
	//one继续走直到最后，每次two跟着走一步
	for {
		if one.Next != nil {
			one = one.Next
			two = two.Next
		} else {
			break
		}
	}
	two.Next = two.Next.Next
	return two
}
