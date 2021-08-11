package node

type INode interface {
	Print(string)
	Clone() INode
}
