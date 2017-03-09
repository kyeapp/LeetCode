/*
 * You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
 *
 * You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 *
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var head ListNode
    constructList(l1, l2, &head, 0)
    return &head
}

func constructList(l1 *ListNode, l2 *ListNode, node *ListNode, carry int) {
    if l1 == nil {
        node.Val = l2.Val + carry
        l2 = l2.Next
    } else if l2 == nil {
        node.Val = l1.Val + carry
        l1 = l1.Next
    } else {
        node.Val = l1.Val + l2.Val + carry
        l1 = l1.Next
        l2 = l2.Next
    }
    
    carry = 0
    if node.Val > 9 {
        carry = 1
        node.Val -= 10
    }
    
    if l1 == nil && l2 == nil {
        if carry != 0 {
            var last ListNode
            last.Val = carry
            node.Next = &last
        }
        return
    }
    
    var temp ListNode
    node.Next = &temp
    constructList(l1, l2, &temp, carry)
}