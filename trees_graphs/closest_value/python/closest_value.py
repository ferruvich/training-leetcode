# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def closest_value(self, root: TreeNode, target: float) -> int:
        closest: int = -1
        minDifference = float("inf")
            
        el = root
        while el != None:
            diff = abs(el.val - target)
            if diff < minDifference:
                minDifference = diff
                closest = el.val
            
            if el.val < target:
                el = el.right
            else:
                el = el.left
                
        return closest