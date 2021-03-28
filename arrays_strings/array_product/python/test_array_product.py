import unittest

import array_product


class TestCase:
    __test__ = False
    test_input: list
    expected_output: list

    def __init__(self, test_input=None, expected_output=None):
        self.test_input = test_input
        self.expected_output = expected_output


class TestProductExceptSelf(unittest.TestCase):

    def test_easy_numbers(self):
        for r in [
            TestCase(test_input=[1, 2, 3, 4], expected_output=[24, 12, 8, 6]),
            TestCase(test_input=[-1, 1, 0, -3, 3], expected_output=[0, 0, 9, 0, 0]),
        ]:
            solution = array_product.Solution()
            product = solution.product_except_self(r.test_input)
            self.assertEqual(r.expected_output, product)


if __name__ == '__main__':
    unittest.main()
