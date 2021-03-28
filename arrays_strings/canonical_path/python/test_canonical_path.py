import unittest

import canonical_path


class TestRemoveDuplicates(unittest.TestCase):

    def test_remove_duplicates(self):
        for test_input, expected_output in {
            "/home/": "/home",
            "/../": "/",
            "/home//foo": "/home/foo",
            "/a/./b/../../c/": "/c"
        }.items():
            sol = canonical_path.Solution()
            res = sol.simplify_path(test_input)
            self.assertEqual(expected_output, res)


if __name__ == '__main__':
    unittest.main()
