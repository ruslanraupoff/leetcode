# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> TreeNode:
        mp = {o: i for i, o in enumerate(inorder)}
        def toTree(index, start, end):
            if start > end:
                return None, index
            
            treeNode = TreeNode(preorder[index])
            index += 1
            idx = mp[treeNode.val]
            
            treeNode.left, index = toTree(index, start, idx - 1)
            treeNode.right, index = toTree(index, idx + 1, end)

            return treeNode, index
        res, index = toTree(0, 0, len(inorder) - 1)
        return res
