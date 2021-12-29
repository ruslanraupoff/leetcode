"""
# Definition for a Node.
class Node:
    def __init__(self, val: int = 0, left: 'Node' = None, right: 'Node' = None, next: 'Node' = None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next
"""

class Solution:
    def connect(self, root: 'Optional[Node]') -> 'Optional[Node]':
        def helper(node):
            if not node: return None
            if node.left:
                node.left.next = node.right
            if node.next and node.left:
                node.right.next = node.next.left
            helper(node.left)
            helper(node.right)
        
        helper(root)
        return root