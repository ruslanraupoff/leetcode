class Solution:
    def breakPalindrome(self, palindrome: str) -> str:
        i, n = 0, len(palindrome)
        if n == 1:
            return ''
        l = list(palindrome)
        while i < n:
            if l[i] != 'a' and i != n-i-1:
                l[i] = 'a'
                break
            i += 1
        if i == n:
            l[i-1] = chr(ord(l[i-1]) + 1)
        
        return ''.join(l)