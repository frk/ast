package golang

import (
	"github.com/frk/ast"
)

// RawStringInsertExpr produces an expression that inserts the node X into a raw string.
type RawStringInsertExpr struct {
	X ExprNode
}

func (s RawStringInsertExpr) Walk(w *ast.Writer) {
	w.Write("` + ")
	s.X.Walk(w)
	w.Write(" + `")
}

type NL struct{}

func (NL) Walk(w *ast.Writer) {
	w.Write("\n")
}

func (NL) stmtNode()    {}
func (NL) commentNode() {}

type StmtList []StmtNode

func (list StmtList) Walk(w *ast.Writer) {
	for _, stmt := range list {
		stmt.Walk(w)
		w.NewLine()
	}
	w.NoNewLine()
}

func (list *StmtList) Add(ss ...StmtNode) {
	*list = append(*list, ss...)
}

func (StmtList) stmtNode() {}

var iferrreturn = IfStmt{
	Cond: BinaryExpr{X: Ident{"err"}, Op: BinaryNeq, Y: Ident{"nil"}},
	Body: BlockStmt{List: []StmtNode{ReturnStmt{Ident{"err"}}}},
}

type IfErrReturn struct{}

func (IfErrReturn) Walk(w *ast.Writer) {
	iferrreturn.Walk(w)
}

func (IfErrReturn) stmtNode() {}

type NoOp struct{}

func (NoOp) Walk(w *ast.Writer) {}

func (NoOp) stmtNode()    {}
func (NoOp) exprNode()    {}
func (NoOp) declNode()    {}
func (NoOp) specNode()    {}
func (NoOp) commentNode() {}

func (n NoOp) exprListNode() []ExprNode { return []ExprNode{n} }
