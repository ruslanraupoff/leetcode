# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def pathSum(self, root: Optional[TreeNode], targetSum: int) -> int:
        if root == None:
            return 0
        self.cnt = 0
        def dfs(node, sum):
            if node == None:
                return

            sum -= node.val
            if sum == 0:
                self.cnt += 1

            dfs(node.left, sum)
            dfs(node.right, sum)

        dfs(root, targetSum)

        return self.cnt + self.pathSum(root.left, targetSum) + self.pathSum(root.right, targetSum)