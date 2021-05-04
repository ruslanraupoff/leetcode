class Solution:
    def checkPossibility(self, nums) -> bool:
        cnt = 0
        for i in range(0, len(nums)-1):
            if nums[i] > nums[i+1]:
                cnt += 1
                if cnt > 1 or i > 0 and nums[i-1] > nums[i+1] and i + 2 < len(nums) and nums[i+2] < nums[i]:
                    return False
        return True
if __name__ == '__main__':
    print('__main__')
    nums = [4, 2, 3]
    res = Solution().checkPossibility(nums)
    print(res)
    