# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def maxProduct(self, root: Optional[TreeNode]) -> int:
        m = 10**9 + 7
        s = []
        def dfs(node, s):
            if not node: return 0
            s.append(node.val + dfs(node.left, s) + dfs(node.right, s))
            return s[-1]

        t = dfs(root, s)
        return max((x * (t - x) for x in s)) % m