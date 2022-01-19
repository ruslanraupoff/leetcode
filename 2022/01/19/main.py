# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def detectCycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head == None or head.next == None:
            return None

        slow, fast = head.next, head.next.next
        while fast != None and fast.next != None and slow != fast:
            slow, fast = slow.next, fast.next.next

        if slow != fast:
            return None

        while slow != head:
            slow, head = slow.next, head.next

        return slow