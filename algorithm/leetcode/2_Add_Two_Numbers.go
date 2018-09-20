// LeetCode Medium
// 2 Add Two Numbers
// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order and each of their nodes contain a single digit.
// Add the two numbers and return it as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Example:
// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 20 ms
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    _sumVal := 0
    _round := 0
    _start := l1
    _start_l2 := l2
    _len_l1 := 0
    _len_l2 := 0
    for l1 != nil || l2 != nil {
        if(l1 == nil) {
            _sumVal = l2.Val + _round
            if _sumVal >= 10 {
                _sumVal = _sumVal - 10
                _round = 1
            } else {
                _round = 0
            }
            l2.Val = _sumVal
            if l2.Next == nil && _round == 1 {
                var _last *ListNode = new(ListNode)
                _last.Val = 1
                _round = 0
                l2.Next = _last
            }
            l2 = l2.Next
            _len_l2 =+ 1
            continue;
        }
        if(l2 == nil) {
            _sumVal = l1.Val + _round
            if _sumVal >= 10 {
                _sumVal = _sumVal - 10
                _round = 1
            } else {
                _round = 0
            }
            l1.Val = _sumVal
            if l1.Next == nil && _round == 1 {
                var _last *ListNode = new(ListNode)
                _last.Val = 1
                _round = 0
                l1.Next = _last
            }
            l1 = l1.Next
            _len_l1 =+ 1
            continue;
        }
        
        _sumVal = l1.Val + l2.Val + _round
        _round = 0
        if _sumVal >= 10 {
            _round = 1
            l1.Val = _sumVal - 10
            l2.Val = _sumVal - 10
        } else {
            l1.Val = _sumVal
            l2.Val = _sumVal
        }
         if l1.Next == nil && l2.Next == nil && _round == 1 {
             var _last *ListNode = new(ListNode)
             _last.Val = 1
             l1.Next = _last
             return _start
        }
        l1 = l1.Next
        l2 = l2.Next
    }
    if _len_l2 > _len_l1 {
        return _start_l2
    }
    return _start
}
