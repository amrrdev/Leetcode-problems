class Solution:
    def groupAnagrams(self, strs: list[str]) -> list[list[str]]:
        hashMap = {}
        for word in strs:
            sortedWord = "".join(sorted(word))
            if sortedWord not in hashMap:
                hashMap[sortedWord] = [word]
            else:
                hashMap[sortedWord].append(word)
        return list(hashMap.values())