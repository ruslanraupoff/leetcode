# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:

    def __init__(self, head: Optional[ListNode]):
        self.l = []
        while head:
            self.l.append(head.val)
            head = head.next

    def getRandom(self) -> int:
        r = int(random.random() * len(self.l))
        return self.l[r]


# Your Solution object will be instantiated and called as such:
# obj = Solution(head)
# param_1 = obj.getRandom()