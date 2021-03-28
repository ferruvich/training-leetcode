class Solution:
    roman_to_int_dict: dict = {
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000
    }

    def roman_to_int(self, s: str) -> int:
        res: int = 0
        prev: str = ""
        for c in s[::-1]:
            n = self.roman_to_int_dict[c]
            if prev and self.roman_to_int_dict[prev] > n:
                res -= n
            else:
                res += n

            prev = c

        return res
