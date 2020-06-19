package golang

import (
	"github.com/frk/ast"
)

type stringnode string

const (
	Ellipsis stringnode = "..."
	True     stringnode = "true"
	False    stringnode = "false"
)

func (s stringnode) Walk(w *ast.Writer) {
	w.Write(string(s))
}

func (stringnode) exprNode() {}

func (x stringnode) exprListNode() []ExprNode { return []ExprNode{x} }
