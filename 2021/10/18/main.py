# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isCousins(self, root: Optional[TreeNode], x: int, y: int) -> bool:
        
        def dfs(root, x):
            if root == None:
                return None, 0

            if (root.left != None and root.left.val == x) or (root.right != None and root.right.val == x):
                return root, 1
            parent, depth = dfs(root.left, x)
            if depth > 0:
                return parent, depth + 1

            parent, depth = dfs(root.right, x)
            if depth > 0:
                return parent, depth + 1

            return None, 0
        
        node = TreeNode()
        node.left = root
        px, dx = dfs(node, x)
        py, dy = dfs(node, y)
        return dx == dy and px != py