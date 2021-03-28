import unittest

import max_profit


class TestCase:
    __test__ = False
    test_input: list
    expected_output: int

    def __init__(self, test_input=None, expected_output=0):
        self.test_input = test_input
        self.expected_output = expected_output


class TestProductExceptSelf(unittest.TestCase):

    def test_easy_numbers(self):
        for r in [
            TestCase(test_input=[7, 1, 5, 3, 6, 4], expected_output=5),
            TestCase(test_input=[7, 6, 4, 3, 1], expected_output=0),
        ]:
            solution = max_profit.Solution()
            product = solution.max_profit(r.test_input)
            self.assertEqual(r.expected_output, product)


if __name__ == '__main__':
    unittest.main()
