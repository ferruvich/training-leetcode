from typing import List


class Solution:
    def three_sum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        res = []

        for i in range(0, len(nums)):
            if nums[i] > 0:
                return res

            if i == 0 or nums[i] != nums[i-1]:
                self.two_sum(nums, i, res)

        return res

    def two_sum(self, nums: List[int], i: int, res: List[List[int]]):
        found = {}

        for j in range(i+1, len(nums)):
            complement = -nums[i] - nums[j]

            if complement in found and found[complement] != j:
                triple = [nums[i], nums[j], complement]
                triple.sort()
                if triple not in res:
                    res.append(triple)

            found[nums[j]] = j

        return res
