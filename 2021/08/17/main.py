# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def goodNodes(self, root: TreeNode) -> int:
        def dfs(node, mx):
            if not node:
                return 0

            curr = (node.val >= mx)
            mx = max(mx, node.val)
            
            return dfs(node.left, mx) + dfs(node.right, mx) + curr

        return dfs(root, -1000001)
