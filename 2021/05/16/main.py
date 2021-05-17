# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def minCameraCover(self, root: TreeNode) -> int:
        self.ans = 0
        def dfs(root: TreeNode):
            if not root:
                return 1
            l = dfs(root.left)
            r = dfs(root.right)
            if l == 2 or r == 2:
                self.ans += 1
                return 3
            elif l == 3 or r == 3:
                return 1
            else:
                return 2
        if dfs(root) == 2:
            self.ans += 1
        return self.ans
