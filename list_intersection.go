package main;
import "fmt"
/**
 * Definition for singly-linked list.
 */
  type ListNode struct {
      Val int
      Next *ListNode
  }
 
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    // A has a total of 'm' nodes 
    // B has a total of 'n' nodes
    // If they are 2 separate lists , we have a total of 'm' + 'n' nodes
    // However, if they meet, then if we traverse one list and then the other, we must 
    // end up at the end of the entire list at the same time, which means we meet at the head at the same time
    // Have a state variable , for each list saying we've already reset it. If it's set to false, and we reach end, point to the other list. If we reach end and it's been set ignore and return false.
    var headAOriginal *ListNode  = headA;
     var headBOriginal *ListNode  = headB;

     headAReset := false;
     headBReset := false;
     for headA != nil && headB != nil {
     	if headA == headB {
     		return headA;
     	}

     	headA = headA.Next;
     	headB = headB.Next;
     	if headA == nil && !headAReset {
     		headA = headBOriginal;
     		headAReset = true;

     	}
     	if headB == nil && !headBReset {
     		headB = headAOriginal;
     		headBReset = true;
     	}
     }
    return nil;
}

func getIntersectionNodeWrapper(headA, headB *ListNode) int {
	var ret *ListNode = getIntersectionNode(headA, headB);
	if ret == nil {
		return 0;
	} else {
		return ret.Val;
	}
} 
func main() {
	fmt.Println("here");
}