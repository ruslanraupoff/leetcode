# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def generateTrees(self, n: int) -> List[Optional[TreeNode]]:
        a = {}
        def dfs(low, high):
            if low > high:
                return [None]
            if (low, high) in a:
                return a[(low,high)]
            ans = []
            for mid in range(low, high + 1):
                left = dfs(low, mid - 1)
                right = dfs(mid + 1, high)
                for l in left:
                    for r in right:
                        root = TreeNode(mid)
                        root.left = l
                        root.right = r
                        ans.append(root)
            a[(low,high)]=ans
            return ans
        return dfs(1, n)