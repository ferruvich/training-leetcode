class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def add_two_numbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        return self._add_to_numbers_recursive(l1, l2, 0)

    def _add_to_numbers_recursive(self, l1: ListNode, l2: ListNode, rest: int) -> ListNode:
        if not l1 and not l2:
            if rest == 0:
                return None

            return ListNode(rest)

        l1_val = l1.val if l1 else 0
        l1_next = l1.next if l1 else None
        l2_val = l2.val if l2 else 0
        l2_next = l2.next if l2 else None

        next_rest = 0
        unit = l1_val + l2_val + rest
        if unit >= 10:
            next_rest = unit // 10
            unit %= 10

        return ListNode(unit, self._add_to_numbers_recursive(l1_next, l2_next, next_rest))
