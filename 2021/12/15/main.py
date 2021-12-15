# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def insertionSortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        prehead = ListNode(0)
        prehead.next = head
        current = head
        while current != None and current.next != None:
            p = current.next
            if current.val <= p.val:
                current = p
                continue

            current.next = p.next
            pre, next = prehead, prehead.next
            while next.val < p.val:
                pre = next
                next = next.next
            pre.next = p
            p.next = next

        return prehead.next