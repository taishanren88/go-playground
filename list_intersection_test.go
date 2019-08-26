package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	var headA ListNode ;
	headA.Val = 1;

	var headASecond ListNode ;
	headASecond.Val = 2;

	var headAThird ListNode ;
	headAThird.Val = 3;

	headA.Next = &headASecond;
	headASecond.Next = &headAThird;


	var headB ListNode;
	headB.Val = 4;

	var headBSecond ListNode;
	headBSecond.Val = 5;

	headB.Next = &headBSecond;
	headBSecond.Next = &headASecond;

	assert.Equal(t, 2, getIntersectionNodeWrapper(&headA, &headB));

}


func TestNoIntersection(t *testing.T) {
	var headA ListNode ;
	headA.Val = 1;

	var headASecond ListNode ;
	headASecond.Val = 2;

	var headAThird ListNode ;
	headAThird.Val = 3;

	headA.Next = &headASecond;
	headASecond.Next = &headAThird;


	var headB ListNode;
	headB.Val = 4;

	var headBSecond ListNode;
	headBSecond.Val = 5;

	var headBThird ListNode;
	headBThird.Val = 6;


	headB.Next = &headBSecond;
	headBSecond.Next = &headBThird;

	assert.Equal(t, 0, getIntersectionNodeWrapper(&headA, &headB));

}