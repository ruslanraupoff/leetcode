# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def oddEvenList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head == None:
            return head

        oe = head
        eb = head.next
        ee = eb

        while ee != None and ee.next != None:
            oe.next = ee.next
            oe = oe.next
            ee.next = oe.next
            ee = ee.next
            oe.next = eb

        return head