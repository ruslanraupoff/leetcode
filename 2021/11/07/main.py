class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        if num1 == "0" or num2 == "0":
            return "0"
        def to_dig(c):
            return ord(c) - ord('0')
        temp = [0] * (len(num1) +len(num2))
        for i in range(len(num1)):
            for j in range(len(num2)):
                temp[i+j+1] += to_dig(num1[i]) * to_dig(num2[j])
        for i in range(len(temp)-1, -1, -1):
            temp[i-1] += temp[i] // 10
            temp[i] = temp[i] % 10
        
        if temp[0] == 0:
            temp = temp[1:]

        return ''.join(map(str, temp))