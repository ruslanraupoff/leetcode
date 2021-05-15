class Solution:
    def isNumber(self, s: str) -> bool:
        return re.match(r'^[\+\-]?(\d+|\d+\.\d*|\.\d+)([eE][\+\-]?\d+)?$', s)
