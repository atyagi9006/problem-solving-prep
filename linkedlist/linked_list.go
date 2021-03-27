package main

func main() {

}

type node struct {
	Data int
	Next *node
}

func deleteMid(head *node) {

	slow:=head
	fast:=head.Next
	for fast.Next
}

func length(head *node) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}

func getNthElement(head *node, n int) *node {
	if n < 0 {
		return nil
	}
	index := 0
	for head != nil && index < n {
		index++
		head = head.Next
	}
	return head //if index is out of bound i return nil
}

func append(head *node, data int) {
	newNode := node{
		Data: data,
		Next: nil,
	}
	if head == nil {
		head = &newNode
		return
	}
	for head != nil {
		head = head.Next
	}

	head.Next = &newNode
}


func pop(head *node) int {
	if head == nil {
		return -1
	}
	data := head.Data
	head = head.Next

	return data

}
