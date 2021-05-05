class Solution:
    def jump(self, nums: List[int]) -> int:
        r, mx, mj = 0, 0, 0
        for i in range(len(nums) - 1):
            mx = max(mx, i + nums[i])
            if i == mj:
                r += 1
                mj = mx
        return r

if __name__ == '__main__':
    print('__main__')
            
