# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def rangeSumBST(self, root: Optional[TreeNode], low: int, high: int) -> int:
        self.s = 0
        if root == None:
            return 0

        if root.val < low:
            self.s = self.rangeSumBST(root.right, low, high)
        elif root.val > high:
            self.s = self.rangeSumBST(root.left, low, high)
        else:
            self.s = root.val
            self.s += self.rangeSumBST(root.left, low, high)
            self.s += self.rangeSumBST(root.right, low, high)
        
        return self.s
        