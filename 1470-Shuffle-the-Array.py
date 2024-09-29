class Solution:
    def shuffle(self, nums: list[int], n: int) -> list[int]:
        answer = []
        p1 = 0
        p2 = n
        while p1 < n:
            answer += [nums[p1], nums[p2]]
            p1 += 1
            p2 += 1
        return answer