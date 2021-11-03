# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def sumNumbers(self, root: Optional[TreeNode]) -> int:
        self.s = 0
        def dfs(root, t):
            if root == None:
                return 0
            
            t = t*10 + root.val
            if root.left == None and root.right == None:
                self.s += t
                return
            
            dfs(root.left, t)
            dfs(root.right, t)
        dfs(root, 0)
        return self.s

            