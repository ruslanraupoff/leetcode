class Solution:
    def pushDominoes(self, dominoes: str) -> str:
        r, i, ll, n = '', 0, 0, len(dominoes)
        while i < n:
            if dominoes[i] == '.':
                i += 1
            elif dominoes[i] == 'L':
                r += 'L' * (i + 1 - ll)
                i += 1
                ll = i
            else:
                r += '.' * (i - ll)
                lr = i
                i += 1
                while i < n:
                    if dominoes[i] == '.':
                        i += 1
                    elif dominoes[i] == 'R':
                        r += 'R' * (i - lr)
                        lr = i
                        i += 1
                    else:
                        t = (i - lr + 1)
                        if t % 2 == 1:
                            r += 'R' * (t // 2) + '.' + 'L' * (t // 2)
                        else:
                            r += 'R' * (t // 2) + 'L' * (t // 2)
                        i += 1
                        lr = i
                        break
                r += 'R' * (i - lr)
                ll = i
        r += '.' * (i - ll)
        return r