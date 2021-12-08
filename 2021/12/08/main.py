# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def findTilt(self, root: Optional[TreeNode]) -> int:
        self.tilt = 0
        def helper(root):
            if root == None:
                return 0

            l = helper(root.left)
            r = helper(root.right)

            if l > r:
                self.tilt += l - r
            else:
                self.tilt += r - l

            return l + r + root.val
        
        helper(root)
        
        return self.tilt
