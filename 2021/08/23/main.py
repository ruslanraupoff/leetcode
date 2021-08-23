# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def findTarget(self, root: Optional[TreeNode], k: int) -> bool:
        def helper(root, sroot, k):
            if root == None:
                return False

            return (2 * root.val != k and bs(sroot, k - root.val)) or helper(root.left, sroot, k) or helper(root.right, sroot, k)
        def bs(root, k):
            if root == None:
                return False
            if root.val == k:
                return True
            if root.val < k:
                return bs(root.right, k)
            return bs(root.left, k)
        return helper(root, root, k)