package sqlang

import (
	"github.com/frk/ast"
)

type DeleteStatement struct {
	Table     Ident
	Using     UsingClause
	Where     WhereClause
	Returning ReturningClause
}

func (s DeleteStatement) Walk(w *ast.Writer) {
	w.NoIndent()
	w.Write("DELETE FROM ")
	s.Table.Walk(w)
	w.NewLine()
	w.Indent()
	s.Using.Walk(w)
	w.NewLine()
	s.Where.Walk(w)
	w.NewLine()
	s.Returning.Walk(w)
	w.NoNewLine()
}

type UsingClause struct {
	List []TableExpr
}

func (c UsingClause) Walk(w *ast.Writer) {
	if len(c.List) == 0 {
		return
	}

	w.Write("USING ")
	for i, x := range c.List {
		if i > 0 {
			w.NewLine()
		}
		x.Walk(w)
	}
}
