package goyacc_lex

import "fmt"

type Caller interface {
	Call(root AbstractNode)
}

type DefaultPrintCaller struct {
}

func (caller *DefaultPrintCaller) Call(root AbstractNode) {
	if root.getType() == NTLeaf {
		_node := root.(*LeafNode)
		fmt.Printf("rule: '%s' <=> '%s'\n", _node.Key, _node.Val)
	} else {
		_node := root.(*LogicNode)
		fmt.Printf("logic: '%s'\n", _node.Rel)
		caller.Call(_node.L)
		caller.Call(_node.R)
	}
}
