import unittest

import remove_duplicates


class TestCase:
    __test__ = False
    test_input: list
    expected_output: list

    def __init__(self, test_input=None, expected_output=None):
        self.test_input = test_input
        self.expected_output = expected_output


class TestRemoveDuplicates(unittest.TestCase):

    def test_remove_duplicates(self):
        for r in [
            TestCase(test_input=[], expected_output=[]),
            TestCase(test_input=[1], expected_output=[1]),
            TestCase(test_input=[1, 1], expected_output=[1]),
            TestCase(test_input=[1, 1, 2], expected_output=[1, 2]),
            TestCase(test_input=[1, 1, 1, 2, 3, 4, 4], expected_output=[1, 2, 3, 4]),
            TestCase(test_input=[1, 2, 3, 3, 4, 5, 5, 6], expected_output=[1, 2, 3, 4, 5, 6]),
        ]:
            sol = remove_duplicates.Solution()
            res = sol.removeDuplicates(r.test_input)
            self.assertEqual(len(r.expected_output), res)


if __name__ == '__main__':
    unittest.main()
