package main

import (
	"fmt"
	"math"
)

func main() {
	n3 := ListNode{
		1,
		nil,
	}
	n2 := ListNode{
		1,
		&n3,
	}
	n1 := ListNode{
		0,
		&n2,
	}
	n0 := ListNode{
		1,
		&n1,
	}
	curNode := &n0
	var sliceVal []int
	if curNode.Val == 1 {
		sliceVal = append(sliceVal, 1)
	}
	for curNode.Next != nil {
		curNode = curNode.Next
		for index, val := range sliceVal {
			sliceVal[index] = val * 2
		}
		if curNode.Val == 1 {
			sliceVal = append(sliceVal, 1)
		}
	}
	ret := 0
	for _, val := range sliceVal {
		ret += val
	}
	fmt.Println(ret)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	curNode := head
	index := 0
	ret := 0
	if curNode.Val == 1 {
		ret += int(math.Pow(float64(2), float64(index)))
	}
	index++
	for curNode.Next != nil {
		curNode = curNode.Next
		if curNode.Val == 1 {
			ret += int(math.Pow(float64(2), float64(index)))
		}
		index++
	}
	return ret
}

func reverseList(head *ListNode) *ListNode {
	curNode := head
	newCurNode := ListNode{
		Val:  curNode.Val,
		Next: nil,
	}
	for curNode.Next != nil {
		newLastNode := newCurNode
		curNode = curNode.Next
		newNextNode := ListNode{
			Val:  curNode.Val,
			Next: &newLastNode,
		}
		newCurNode = newNextNode
	}
	return &newCurNode
}
