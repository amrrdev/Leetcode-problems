import collections

class Solution:
    def groupAnagrams(self, strs: list[str]) -> list[list[str]]:
        hashMap = collections.defaultdict(list)
        for word in strs:
            count = [0] * 26
            for char in word:
                count[ord(char) - ord("a")] += 1
            hashMap[tuple(count)].append(word)
        return hashMap.values()