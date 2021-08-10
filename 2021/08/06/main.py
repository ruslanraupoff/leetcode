"""
# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""

class Solution:
    def levelOrder(self, root: 'Node') -> List[List[int]]:
        r = []
        q = [root]
        
        while q:
            r.append([n.val for n in q])
            q = [c for n in q for c in n.children]

        return r

        

