# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        l = []
        def dfs(root, i):
            if root is None:
                return
            
            if len(l) < i + 1:
                l.append([])
            
            l[i].append(root.val)
            
            dfs(root.left, i+1)
            dfs(root.right, i+1)
            
        dfs(root, 0)
        
        return l
        