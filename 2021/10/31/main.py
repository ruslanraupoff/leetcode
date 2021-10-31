"""
# Definition for a Node.
class Node:
    def __init__(self, val, prev, next, child):
        self.val = val
        self.prev = prev
        self.next = next
        self.child = child
"""

class Solution:
    def flatten(self, head: 'Node') -> 'Node':
        if head == None:
            return None
        curr = head
        while curr != None:
            if curr.child == None:
                curr = curr.next
                continue

            child = curr.child
            while child.next != None:
                child = child.next
            child.next = curr.next
            if curr.next != None:
                curr.next.prev = child
            
            curr.next = curr.child
            curr.child.prev = curr
            curr.child = None
        
        return head