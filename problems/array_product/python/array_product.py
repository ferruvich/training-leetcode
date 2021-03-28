from typing import List


class Solution:
    """
    This version does not use division:
    Complexity:
    O(n) time - as we traverse the list only twice, so O(2n)->O(n)
    (1) space - if we do not consider the output list
    """

    @staticmethod
    def product_except_self(nums: List[int]) -> List[int]:
        res = [0] * len(nums)

        for i in range(len(nums)):
            if i == 0:
                res[i] = 1
                continue

            res[i] = nums[i - 1] * res[i - 1]

        prec = 1
        for i in reversed(range(1, len(nums))):
            prec = prec * nums[i]
            res[i - 1] *= prec

        return res
