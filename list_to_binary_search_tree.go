package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// *
//  * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		single := new(TreeNode)
		single.Val = head.Val
		single.Left = nil
		single.Right = nil
		return single
	}

	slow := head
	fast := head
	var slowParent *ListNode = nil
	for fast != nil && fast.Next != nil {
		slowParent = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if slow != nil {
		parent := new(TreeNode)
		if slowParent != nil {
			slowParent.Next = nil
		}

		parent.Val = slow.Val
		parent.Left = buildTree(head)
		parent.Right = buildTree(slow.Next)
		return parent
	}
	return nil
}
func sortedListToBST(head *ListNode) *TreeNode {
	// Loop through the list with fast (moves 2x) and slow (moves 1x) pointers
	// Stop when fast == null || fast->next == null, slow will be right at the center, start tree here
	// center->left->right needs to be set correctly
	// Left is left child ,right is right child
	return buildTree(head)
}

func main() {

}
