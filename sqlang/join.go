package sqlang

import (
	"github.com/frk/ast"
)

type TableExpr interface {
	Node
	tableExprNode()
}

type JoinType string

func (typ JoinType) Walk(w *ast.Writer) {
	w.Write(string(typ))
}

const (
	JoinNone  JoinType = ""
	JoinLeft  JoinType = "LEFT JOIN"
	JoinRight JoinType = "RIGHT JOIN"
	JoinFull  JoinType = "FULL JOIN"
	JoinCross JoinType = "CROSS JOIN"
)

type TableJoin struct {
	Type JoinType
	Rel  Ident
	Cond JoinCondition // can be left empty for cross joins
}

func (tj TableJoin) Walk(w *ast.Writer) {
	tj.Type.Walk(w)
	w.Write(" ")
	tj.Rel.Walk(w)

	// write the join_condition if not cross join
	if tj.Type != JoinCross {
		w.Write(" ")
		tj.Cond.Walk(w)
	}
}

type JoinCondition struct {
	SearchCondition BoolValueExpr
}

func (cond JoinCondition) Walk(w *ast.Writer) {
	if cond.SearchCondition != nil {
		w.Write("ON ")
		cond.SearchCondition.Walk(w)
	}
}

func (Ident) tableExprNode()     {}
func (TableJoin) tableExprNode() {}
