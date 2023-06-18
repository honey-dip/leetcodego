package design_linked_list

type MyLinkedListNode struct {
	Val  int
	Next *MyLinkedListNode
}

type MyLinkedList struct {
	head   *MyLinkedListNode
	length int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		head:   nil,
		length: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index > this.length-1 {
		return -1
	}
	cur := this.head
	i := 0
	for cur != nil {
		if i == index {
			return cur.Val
		}
		cur = cur.Next
		i++
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := MyLinkedListNode{
		Val:  val,
		Next: this.head,
	}
	this.head = &node
	this.length++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := MyLinkedListNode{
		Val:  val,
		Next: nil,
	}
	cur := this.head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &node
	this.length++

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {

}

func (this *MyLinkedList) DeleteAtIndex(index int) {

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
