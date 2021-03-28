import unittest

import add_two_numbers


class TestCase:
    __test__ = False

    def __init__(self, input_list_one, input_list_two, expected_result):
        self.input_list_one = LinkedListFactory.from_list(input_list_one)
        self.input_list_two = LinkedListFactory.from_list(input_list_two)
        self.expected_result = expected_result


class TestAddTwoNumbers(unittest.TestCase):

    def test_add_two_numbers(self):
        for t in [
            TestCase([2, 4, 3], [5, 6, 4], [7, 0, 8]),
            TestCase([0], [0], [0]),
            TestCase([9, 9, 9, 9, 9, 9, 9], [9, 9, 9, 9], [8, 9, 9, 9, 0, 0, 0, 1])
        ]:
            solution = add_two_numbers.Solution()
            result = solution.add_two_numbers(t.input_list_one, t.input_list_two)
            self._test_result(t.expected_result, result)

    def _test_result(self, expected: list, result: add_two_numbers.ListNode):
        if (not expected) and (not result):
            return

        self.assertIsNotNone(expected)
        self.assertIsNotNone(result)

        self.assertEqual(expected[0], result.val)

        self._test_result(expected[1:], result.next)


class LinkedListFactory:
    @classmethod
    def from_list(cls, to_be_linked: list) -> add_two_numbers.ListNode:
        if not to_be_linked:
            return None

        return add_two_numbers.ListNode(
            to_be_linked[0],
            cls.from_list(to_be_linked[1:])
        )
