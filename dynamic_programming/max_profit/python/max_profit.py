from typing import List


class Solution:
    """
    Complexity:
    O(n) time - as we traverse the list only once
    O(1) space - as we're using only two variables
    """
    @staticmethod
    def max_profit(prices: List[int]) -> int:
        max_profit = 0
        min_price = prices[0]
        for el in prices[1:]:
            if el < min_price:
                min_price = el
            elif el - min_price > max_profit:
                max_profit = el - min_price

        return max_profit
