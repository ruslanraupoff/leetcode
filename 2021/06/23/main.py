# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseBetween(self, head: ListNode, left: int, right: int) -> ListNode:
        pre = new = ListNode(0)
        new.next = head
        for _ in range(left-1):
            pre = pre.next
        curr, prev = pre.next, None
        for _ in range(right - left + 1):
            prev, curr.next, curr = curr, prev, curr.next
        pre.next.next = curr
        pre.next = prev
        return new.next
        