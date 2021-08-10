# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def pathSum(self, root: Optional[TreeNode], targetSum: int) -> List[List[int]]:
        if not root: 
            return []
        r = []
        self.dfs(root, targetSum, [], r)
        return r

    def dfs(self, root, target, path, r):
        if not root:
            return
        target -= root.val
        path.append(root.val)
        if not root.left and not root.right:
            if not target:
                r.append(path[:])
        self.dfs(root.left, target, path, r)
        self.dfs(root.right, target, path, r)
        path.pop()
