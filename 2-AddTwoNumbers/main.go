package main

import "fmt"

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

    Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
    Output: 7 -> 0 -> 8
    Explanation: 342 + 465 = 807.
*/

// 两个非空的链表，表示两个非负的整数，链表上的每一个节点(node)代表了一个数字(0-9)，将这两个数字相加，
// 返回一个新的链表

// ListNode is the struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func asNumber(l *ListNode) int {
	var result = 0
	var factor = 1
	for {
		result += l.Val * factor
		if l.Next != nil {
			l = l.Next
			factor = 10 * factor
		} else {
			break
		}
	}

	return result
}

func asList(v int) *ListNode {
	var l ListNode
	l.Val = v % 10
	if v/10 != 0 {
		l.Next = asList(v / 10)
	}

	return &l
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	r := asNumber(l1) + asNumber(l2)
	return asList(r)
}

func addTwoNumbers2(l1, l2 *ListNode) *ListNode {
	var sum int
	var preResult = &ListNode{Val: 0}

	var result = preResult
	for {
		if l1 == nil && l2 == nil {
			break
		}
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		result.Next = &ListNode{Val: sum % 10}
		sum = sum / 10

		result = result.Next

	}
	if sum == 1 {
		result.Next = &ListNode{Val: 1}
	}
	return preResult.Next
}

func main() {
	// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
	// Output: 7 -> 0 -> 8

	l1 := &ListNode{Val: 9, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	fmt.Printf("%v + %v\n", asNumber(l1), asNumber(l2))

	result := addTwoNumbers2(l1, l2)
	fmt.Printf("%v\n", asNumber(result))
}
