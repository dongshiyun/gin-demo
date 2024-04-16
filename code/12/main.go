//  知识点：链表
//  字节跳动,猿辅导,美团点评,百度,shopee
//  中等
//  考察次数8

// 题目描述
// 假设链表中每一个节点的值都在 0 - 9 之间，那么链表整体就可以代表一个整数。
// 给定两个这种链表，请生成代表两个整数相加值的结果链表。
// 例如：链表 1 为 9->3->7，链表 2 为 6->3，最后生成新的结果链表为 1->0->0->0。
// 示例1
// 输入
// [9,3,7],[6,3]
// 输出
// {1,0,0,0}

//基本原理：加法从个位开始相加，所以要先反转两个链表，反转后遍历，第一个相加作为新链表的第一个节点，注意是否有进位，依次类推

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a1 := []int{9, 3, 7}
	r1 := arrayToListNode(a1)
	//shownode(r1)

	a2 := []int{6, 3}
	r2 := arrayToListNode(a2)
	//shownode(r2)

	a3 := addInList(r1, r2)
	shownode(a3)
}

/**
 *
 * @param head1 ListNode类
 * @param head2 ListNode类
 * @return ListNode类
 */
func addInList(head1 *ListNode, head2 *ListNode) *ListNode {
	// write code here
	if head1 == nil && head2 == nil {
		return nil
	}
	//反转链表
	h1 := reverseList(head1)
	h2 := reverseList(head2)
	//初始化一个新链表
	head := &ListNode{}

	//记录进位
	carry := 0
	//遍历相加，注意进位
	var val1, val2 int
	for h1 != nil || h2 != nil || carry > 0 {
		val1 = 0
		val2 = 0
		if h1 != nil {
			val1 = h1.Val
			h1 = h1.Next
		}
		if h2 != nil {
			val2 = h2.Val
			h2 = h2.Next
		}
		sum := val1 + val2 + carry
		//判断是否需要进位
		if sum >= 10 {
			carry = sum / 10
			sum = sum % 10
		} else {
			carry = 0
		}

		rNode := &ListNode{Val: sum}
		rNode.Next = head.Next
		head.Next = rNode
	}
	return head.Next
}

//封装反转函数，并返回表头
func reverseList(pHead *ListNode) *ListNode {
	// write code here
	if pHead == nil {
		return nil
	}

	if pHead.Next == nil {
		return pHead
	}

	a := reverseList(pHead.Next)
	pHead.Next.Next = pHead
	pHead.Next = nil

	return a
}

//数组转换为链表
func arrayToListNode(a []int) *ListNode {
	resultSentinel := &ListNode{}
	for i := len(a) - 1; i >= 0; i-- {
		//fmt.Println(a[i])
		// 结果节点， 头插法
		rNode := &ListNode{Val: a[i]}
		rNode.Next = resultSentinel.Next
		resultSentinel.Next = rNode
	}
	return resultSentinel.Next
}

func shownode(p *ListNode) { //遍历
	for p != nil {
		fmt.Println(*p)
		p = p.Next //移动指针
	}
}
