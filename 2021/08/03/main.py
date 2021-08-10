class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        r = []
        n = len(nums)
        nums.sort()
        def backtrack(path, j):
            r.append(path)
            for i in range(j, n):
                if i > j and nums[i] == nums[i-1]:
                    continue
                backtrack([*path, nums[i]], i+1)
        backtrack([], 0)
        return r