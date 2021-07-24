class Solution:
    def findLadders(self, beginWord: str, endWord: str, wordList: List[str]) -> List[List[str]]:
        result = []
        mp = {w: True for w in wordList}
        if endWord not in mp:
            return result
        queue = [[beginWord]]
        queueLen = 1
        levelMap = {}
        while len(queue):
            path = queue[0]
            queue = queue[1:]
            lastWord = path[-1]
            for i in range(len(lastWord)):
                for c in range(26):
                    nextWord = lastWord[:i] + chr(c + ord('a')) + lastWord[i+1:]
                    if nextWord == endWord:
                        path.append(endWord)
                        result.append(path)
                        continue
                    if nextWord in mp:
                        levelMap[nextWord] = True
                        newPath = path.copy()
                        newPath.append(nextWord)
                        queue.append(newPath)
            queueLen -= 1
            if queueLen == 0:
                if len(result) > 0:
                    return result
                for k in levelMap:
                    del mp[k]
                levelMap = {}
                queueLen = len(queue)
        return result