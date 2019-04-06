package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	fmt.Println((l1.Next).Val)
	return l1
}
func main() {
	var ll3 = ListNode{Val: 3, Next: nil}
	var ll2 = ListNode{Val: 4, Next: &ll3}
	var ll1 = ListNode{Val: 2, Next: &ll2}
	var lll3 = ListNode{Val: 4, Next: nil}
	var lll2 = ListNode{Val: 6, Next: &lll3}
	var lll1 = ListNode{Val: 5, Next: &lll2}
	fmt.Println("result: ", addTwoNumbers(&ll1, &lll1))
}
