//  知识点：链表
//  快手,字节跳动,一点资讯,招银网络,腾讯,京东,小米,百度
//  简单
//  考察次数7
// 题目描述
// 将两个有序的链表合并为一个新链表，要求新的链表是通过拼接两个链表的节点来生成的
// 示例1
// 输入
// {1},{}
// 输出
// {1}
// 示例2
// 输入
// {1},{1}
// 输出
// {1,1}
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := createNode(0)
	l2 := createNode(0)

	head := mergeTwoLists(l1, l2)
	shownode(head)
}

/**
 *
 * @param l1 ListNode类
 * @param l2 ListNode类
 * @return ListNode类
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// write code here
	//注意这句话“要求新的链表是通过拼接两个链表的节点来生成的”
	//决定了我们不能重新建立一个链表，只能通过修改l1和l2两个链表的next合并为一个链表
	if l1 == nil {
		//l1没有了，所以后面节点都是l2
		return l2
	}
	if l2 == nil {
		//l2没有了，所以后面节点都是l1
		return l1
	}
	//假如l1的第一个值比l2的第一个值小，就返回l1否则l2，作为链表的开始
	if l1.Val < l2.Val {
		//l1的第一个节点已经被使用，递归 判断l1的下一个节点和l2的第一个节点进行判断，获得小的值作为下一个节点
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		//l2的第一个节点已经被使用，递归 判断l2的下一个节点和l1的第一个节点进行判断，获得小的值作为下一个节点
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
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
