package pkg

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "", args: args{fromArray([]int{})}, want: fromArray([]int{})},
		{name: "1", args: args{fromArray([]int{1})}, want: fromArray([]int{1})},
		{name: "2", args: args{fromArray([]int{1, 2})}, want: fromArray([]int{2, 1})},
		{name: "3", args: args{fromArray([]int{1, 2, 3})}, want: fromArray([]int{2, 1, 3})},
		{name: "4", args: args{fromArray([]int{1, 2, 3, 4})}, want: fromArray([]int{2, 1, 4, 3})},
		{name: "5", args: args{fromArray([]int{1, 2, 3, 4, 5})}, want: fromArray([]int{2, 1, 4, 3, 5})},
		{name: "6", args: args{fromArray([]int{1, 2, 3, 4, 5, 6})}, want: fromArray([]int{2, 1, 4, 3, 6, 5})},
		{name: "7", args: args{fromArray([]int{1, 2, 3, 4, 5, 6, 7})}, want: fromArray([]int{2, 1, 4, 3, 6, 5, 7})},
		{name: "8", args: args{fromArray([]int{1, 2, 3, 4, 5, 6, 7, 8})}, want: fromArray([]int{2, 1, 4, 3, 6, 5, 8, 7})},
		{name: "9", args: args{fromArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})}, want: fromArray([]int{2, 1, 4, 3, 6, 5, 8, 7, 9})},
		{name: "10", args: args{fromArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})}, want: fromArray([]int{2, 1, 4, 3, 6, 5, 8, 7, 10, 9})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := tt.args.head
			if got := swapPairs(head); check(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func fromArray(arr []int) *ListNode {
	var root = &ListNode{}
	var cur = root
	for _, v := range arr {
		cur.Next = &ListNode{Val: v, Next: nil}
		cur = cur.Next
	}
	return root.Next
}

func check(node1 *ListNode, node2 *ListNode) bool {
	var arr1 []int
	var arr2 []int
	for node1 != nil {
		arr1 = append(arr1, node1.Val)
		node1 = node1.Next
	}
	for node2 != nil {
		arr2 = append(arr2, node2.Val)
		node2 = node2.Next
	}
	fmt.Println(arr1)
	fmt.Println(arr2)
	return !reflect.DeepEqual(arr1, arr2)
}
