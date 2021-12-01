# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def deleteNode(self, root: Optional[TreeNode], key: int) -> Optional[TreeNode]:
        def findMax(node):
            if node.right == None:
                return node.val
            return findMax(node.right)

        if root == None:
            return root

        if root.val == key:
            if root.left == None and root.right == None:
                return None
            elif root.left == None:
                return root.right
            elif root.right == None:
                return root.left
            else:
                max = findMax(root.left)
                root.val = max
                root.left = self.deleteNode(root.left, max)
        elif root.val > key:
            root.left = self.deleteNode(root.left, key)
        else:
            root.right = self.deleteNode(root.right, key)

        return root
