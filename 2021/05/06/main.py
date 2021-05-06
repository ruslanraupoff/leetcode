# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def sortedListToBST(self, head: ListNode) -> TreeNode:
        
        def sortToBST(begin, end):
            if begin == end:
                return None
            
            if begin.next == end:
                return TreeNode(begin.val)
            
            fast, slow = begin, begin
            while fast != end and fast.next != end:
                fast = fast.next.next
                slow = slow.next
            mid = slow
            
            return TreeNode(mid.val, sortToBST(begin, mid), sortToBST(mid.next, end))
        
        return sortToBST(head, None)