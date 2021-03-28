import unittest

import three_sum


class TestThreeSum(unittest.TestCase):

    def test_three_sum(self):
        for nums, solution in [
            [[-1, 0, 1, 2, -1, -4], [[-1, -1, 2], [-1, 0, 1]]],
            [[], []],
            [[0, 1], []],
            [[0], []],
            [[1, 2, 3, 4, 5, 6], []],
            [[-1, 0, -1, 1, 2, 2, -4], [[-1, 0, 1], [-4, 2, 2], [-1, -1, 2]]]
        ]:
            s = three_sum.Solution()
            res = s.three_sum(nums)
            self.assertEqual(len(solution), len(res))
            for r in res:
                if r not in solution:
                    self.fail(f"illegal element {r}, should not be in solution ({solution})")


if __name__ == '__main__':
    unittest.main()
