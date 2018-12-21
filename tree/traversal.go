package tree

func (node *Node) Traverse()  {
	if nil == node {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
