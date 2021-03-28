import unittest

import closest_value


class TestCase:
    __test__ = False
    tree_node: closest_value.TreeNode
    target: float
    expected_output: int

    def __init__(
        self,
        tree_node=None, 
        target=0.0,
        expected_output=0
    ):
        self.tree_node = tree_node
        self.target = target
        self.expected_output = expected_output


class TestClosestValue(unittest.TestCase):

    def test_closest_value(self):
        for t in [
            TestCase(TreeNodeFactory.create_tree_from_list([4, 2, 5, 1, 3]), 3.714286, 4),
            TestCase(TreeNodeFactory.create_tree_from_list([4, 2, 5, 1, 3]), 3.214286, 3),
            TestCase(TreeNodeFactory.create_tree_from_list([1]), 3.214286, 1)
        ]:
            sol = closest_value.Solution()
            res = sol.closest_value(t.tree_node, t.target)
            self.assertEqual(t.expected_output, res)


class TreeNodeFactory:
    __test__ = False

    @classmethod
    def create_tree_from_list(cls, tree_list: list) -> closest_value.TreeNode:
        
        if not tree_list:
            return None
        
        root = closest_value.TreeNode(tree_list[0])
        for el in tree_list[1:]:
            root = cls._insert_element(root, el)

        return root
    
    @classmethod
    def _insert_element(cls, node: closest_value.TreeNode, value: int) -> closest_value.TreeNode:
        if not node:
            return closest_value.TreeNode(value)

        if node.val < value:
            node.right = cls._insert_element(node.right, value)
        else:
            node.left = cls._insert_element(node.left, value)

        return node

if __name__ == '__main__':
    unittest.main()
