class Solution:
    def maxLength(self, arr: List[str]) -> int:
        dp = [set()]
        for st in arr:
            if len(set(st)) < len(st):
                continue
            st = set(st)
            for c in dp[:]:
                if st & c:
                    continue
                dp.append(st | c)

        return max(len(a) for a in dp)