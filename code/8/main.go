//  知识点：链表  双指针  快慢指针
//  字节跳动,腾讯,神策数据,网易,美团点评,招银网络,百度,京东,shopee
//  中等
//  考察次数11
//  题目描述
//  对于一个给定的链表，返回环的入口节点，如果没有环，返回null
//  拓展：
//  你能给出不利用额外空间的解法么？
//  应用，LRU算法 文件系统 内存分配

//简单
// 判断给定的链表中是否有环
// 扩展：
// 你能给出空间复杂度O(1)的解法么？

//定律
//如何判断是否存在环？
//追赶法，设定两个指针slow、fast，从头指针开始，每次分别前进1步、2步。如存在环，则两者相遇；如不存在环，fast遇到nil退出
//如何知道环的长度？
//记录下上一题的碰撞点p，slow、fast从该点开始，再次碰撞所走过的操作数就是环的长度s
//如何找出环的连接点在哪里？
//碰撞点p到连接点的距离=头指针到连接点的距离，因此，分别从碰撞点、头指针开始走，相遇的那个点就是连接点。
//带环链表的长度是多少？
//连接点距离头指针的长度，加上环的长度，二者之和就是带环单链表的长度
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//初始化一个链表
	head := createNode(3)
	//shownode(head)
	//判断是否有环
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next      //跳1步
		fast = fast.Next.Next //跳2步
		if slow == fast {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		fmt.Println(nil)
		return
	}

	//获取节点
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	fmt.Println(slow)
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
