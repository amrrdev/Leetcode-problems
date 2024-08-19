class Solution:
    def twoSum(self, nums: list[int], target: int) -> list[int]:
        hashMap = {}
        for i in range(len(nums)):
            if target - nums[i] in hashMap:
                return [i, hashMap[(target - nums[i])]]
            hashMap[nums[i]] = i
