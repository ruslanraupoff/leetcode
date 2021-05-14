# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def flatten(self, root: TreeNode) -> None:
        """
        Do not return anything, modify root in-place instead.
        """
        if root == None or (root.left == None and root.right == None):
            return root
        if root.left == None:
            return self.flatten(root.right)
        
        if root.right == None:
            root.right = root.left
            root.left = None
            return self.flatten(root.right)
        
        r = self.flatten(root.right)
        self.flatten(root.left).right = root.right
        root.right = root.left
        root.left = None
        return r
