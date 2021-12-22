# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        """
        Do not return anything, modify head in-place instead.
        """
        if head == None:
            return

        cur = head
        size = 0
        while cur != None:
            cur = cur.next
            size+=1

        cur = head
        for i in range((size-1)//2):
            cur = cur.next

        next = cur.next
        while next != None:
            temp = next.next
            next.next = cur
            cur = next
            next = temp
        end = cur

        while head != end:
            hnext = head.next
            enext = end.next
            head.next = end
            end.next = hnext
            head = hnext
            end = enext

        end.next = None