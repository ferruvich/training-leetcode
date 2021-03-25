import unittest

import roman_to_int


class TestRomanToInt(unittest.TestCase):

    def test_easy_numbers(self):
        for roman, number in {
            "III": 3,
            "XXX": 30,
            "LVIII": 58,
        }.items():
            s = roman_to_int.Solution()
            n = s.roman_to_int(roman)
            self.assertEqual(number, n, f"should be {number}")

    def test_subtracting_numbers(self):
        for roman, number in {
            "IV": 4,
            "IX": 9,
            "MCMXCIV": 1994,
        }.items():
            s = roman_to_int.Solution()
            n = s.roman_to_int(roman)
            self.assertEqual(number, n, f"should be {number}")


if __name__ == '__main__':
    unittest.main()
