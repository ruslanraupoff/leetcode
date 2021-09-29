# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def splitListToParts(self, head: Optional[ListNode], k: int) -> List[Optional[ListNode]]:
        def length(head):
            res = 0
            while head != None:
                res += 1 
                head = head.next
            return res

        i, n = 0, length(head)
        res = []
        subArray = []
        crnt = head
        dd = n // k
        md = n % k
        for i in range(k):
            head = crnt
            for j in range(dd + (i < md) - 1):
                if crnt: 
                    crnt = crnt.next
            if crnt: 
                crnt.next, crnt = None, crnt.next
            res.append(head)

        return res