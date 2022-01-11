# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def sumRootToLeaf(self, root: Optional[TreeNode]) -> int:
        def dfs(node, b):
            if node == None:
                return 0
            b = b * 2 + node.val
            if node.left == node.right:
                return b
            return dfs(node.left, b) + dfs(node.right, b)
        return dfs(root, 0)
