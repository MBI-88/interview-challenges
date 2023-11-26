package cmd






type Interface interface {
	Add(data int) *AVL
	PrintAVL(space int)
	RootNode() int

}



func NewBalancedTree(data int) Interface {
	return &AVL{node: data}
}