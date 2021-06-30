# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        def findLCA(root: 'TreeNode', p: 'TreeNode', q: 'TreeNode'):
            if root == None or root == p or root == q:
                return root
            
            l = findLCA(root.left, p, q)
            r = findLCA(root.right, p, q)

            if l and r:
                return root
            
            return l if l is not None else r
        
        return findLCA(root, p, q)