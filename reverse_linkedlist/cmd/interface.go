package cmd




type LinkedList interface {
	Push(data int ) *linkedList
	ReverseLinkedList()
	PrintLinkedList()
	StrToInt(arr []string) []int
}



func NewLinkedLis() LinkedList {
	return &linkedList{}
}