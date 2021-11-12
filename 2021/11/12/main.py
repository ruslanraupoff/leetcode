# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeElements(self, head: Optional[ListNode], val: int) -> Optional[ListNode]:
        root = ListNode(0)
        root.next = head
        prev = root
        while prev != None:
            if prev.next != None and prev.next.val == val:
                prev.next = prev.next.next
            else:
                prev = prev.next
        return root.next