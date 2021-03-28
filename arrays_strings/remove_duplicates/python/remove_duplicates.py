from typing import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        nums_map = {x: i for i, x in enumerate(nums)}
        for i, x in enumerate(nums_map.keys()):
            nums[i] = x

        return len(nums[:len(nums_map)])
