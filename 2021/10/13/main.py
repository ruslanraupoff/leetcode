# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def bstFromPreorder(self, preorder: List[int]) -> Optional[TreeNode]:
        self.i = 0
        n = len(preorder)
        
        def bfs(b):
            if n == self.i or preorder[self.i] > b:
                return None
            
            root = TreeNode(val=preorder[self.i])
            self.i += 1
            root.left = bfs(root.val)
            root.right = bfs(b)
            return root
        
        return bfs(1 << 32)