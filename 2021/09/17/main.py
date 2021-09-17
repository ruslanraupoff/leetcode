class Solution:
    def intersect(self, nums1: List[int], nums2: List[int]) -> List[int]:
        m1, m2, r = {}, {}, []
        for n in nums1:
            if n in m1:
                m1[n] += 1
            else:
                m1[n] = 1
        for n in nums2:
            if n in m2:
                m2[n] += 1
            else:
                m2[n] = 1
        
        if len(list(m1)) > len(list(m2)):
            m1, m2 = m2, m1

        for n in m1:
            if n in m2:
                m = min(m1[n], m2[n])
                for i in range(m):
                    r.append(n)
        return r
        