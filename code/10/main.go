//  知识点：链表
//  shopee,百度,字节跳动,京东,快手,猿辅导,美团点评,腾讯,作业帮,昆仑万维,奇安信,网易,早稻科技,招银网络
//  简单
//  考察次数31
// 题目描述
// 输入一个链表，反转链表后，输出新链表的表头。

//栈：先进后出，只能从一端出
//队列：先进先出
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//初始化一个链表
	head := createNode(0)
	shownode(head)

	a := ReverseList(head)
	shownode(a)
}

/**
 *
 * @param pHead ListNode类
 * @return ListNode类
 */
func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	if pHead == nil {
		return nil
	}

	if pHead.Next == nil {
		return pHead
	}

	a := ReverseList(pHead.Next)
	pHead.Next.Next = pHead
	pHead.Next = nil

	return a
}

func createNode(n int) *ListNode {
	//初始化一个链表
	var head = new(ListNode)
	head.Val = 0
	var tail, tailNext *ListNode
	tail = head //tail用于记录最末尾的节点的地址，刚开始tail的的指针指向头节点
	for i := 1; i < 10; i++ {
		var listNode = ListNode{Val: i}
		//设置上一个节点next为本节点的指针
		(*tail).Next = &listNode
		tail = &listNode

		//返回到第4个节点
		if i == n && n != 0 {
			tailNext = &listNode
		}
	}
	//让最后一个节点跳转到环节点形成环
	if n == 0 {
		tailNext = nil
	}
	(*tail).Next = tailNext

	return head
}

func shownode(p *ListNode) { //遍历
	for p != nil {
		fmt.Println(*p)
		p = p.Next //移动指针
	}
}
