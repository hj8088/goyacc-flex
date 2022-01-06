package goyacc_lex

type LogicRel int

const (
	LR_AND LogicRel = 1
	LR_OR  LogicRel = 2
)

func (lr LogicRel) String() string {
	switch lr {
	case LR_OR:
		return "OR"
	case LR_AND:
		return "AND"
	}
	return "ERR"
}

type NodeType int

const (
	NTLeaf NodeType = iota + 1
	NTLogic
)

type AbstractNode interface {
	getType() NodeType
}

type LeafNode struct {
	Key string
	Val string
}

func (node *LeafNode) getType() NodeType {
	return NTLeaf
}

type LogicNode struct {
	Rel LogicRel
	L   AbstractNode
	R   AbstractNode
}

func (node *LogicNode) getType() NodeType {
	return NTLogic
}

func newLeafNode(key, val string) AbstractNode {
	return &LeafNode{
		Key: key,
		Val: val,
	}
}
func newLogicNode(rel LogicRel, l, r AbstractNode) AbstractNode {
	return &LogicNode{
		Rel: rel,
		L:   l,
		R:   r,
	}
}
