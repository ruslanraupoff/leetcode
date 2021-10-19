class Solution:
    def nextGreaterElement(self, nums1: List[int], nums2: List[int]) -> List[int]:
        r, l = [], {}
        for i, n in enumerate(nums2):
            l[n] = i
        for i, n in enumerate(nums1):
            r.append(-1)
            for j in range(l[n]+1, len(nums2)):
                if n < nums2[j]:
                    r[i] = nums2[j]
                    break
        return r