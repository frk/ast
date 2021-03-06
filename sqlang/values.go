package sqlang

import (
	"strconv"

	"github.com/frk/ast"
)

type ValueExpr interface {
	Node
	valueExprNode()
}

// ValueExprList produces a comma separated list of SQL value expressions.
type ValueExprList []ValueExpr

func (list ValueExprList) Walk(w *ast.Writer) {
	if len(list) == 0 {
		return
	}

	list[0].Walk(w)
	for _, x := range list[1:] {
		w.NewLine()
		w.Write(", ")
		x.Walk(w)
	}
}

const (
	// NOOP does not produce any SQL element.
	NOOP nooptype = 0
)

// ColumnReference produces an SQL column reference.
type ColumnReference struct {
	Qual string
	Name Name
}

func (r ColumnReference) Walk(w *ast.Writer) {
	if len(r.Qual) > 0 {
		w.Write(r.Qual)
		w.Write(".")
	}
	r.Name.Walk(w)
}

// DynamicParmeterSpec produces an SQL dynamic parameter specification.
type DynamicParmeterSpec struct{}

func (DynamicParmeterSpec) Walk(w *ast.Writer) {
	w.Write("?")
}

// OrdinalParameterSpec produces a PostgreSQL ordinal parameter specification.
type OrdinalParameterSpec struct {
	N int
}

func (s OrdinalParameterSpec) Walk(w *ast.Writer) {
	w.Write("$")
	w.Write(strconv.Itoa(s.N))
}

// Literal produces an SQL value verbatim.
type Literal struct {
	Value string
}

func (lit Literal) Walk(w *ast.Writer) {
	w.Write(lit.Value)
}

// RoutineInvocation produces an SQL function call.
type RoutineInvocation struct {
	Name string
	Args []ValueExpr
}

func (r RoutineInvocation) Walk(w *ast.Writer) {
	w.Write(r.Name)
	w.Write("(")

	for i, a := range r.Args {
		if i > 0 {
			w.Write(", ")
		}
		a.Walk(w)
	}

	w.Write(")")
}

// QuantifiedExpr...
type QuantifiedExpr struct {
	Qua  QUANTIFIER
	Expr ValueExpr
}

func (x QuantifiedExpr) Walk(w *ast.Writer) {
	x.Qua.Walk(w)
	w.Write("(")
	x.Expr.Walk(w)
	w.Write(")")
}

type CastExpr struct {
	Expr ValueExpr
	Type string
}

func (x CastExpr) Walk(w *ast.Writer) {
	x.Expr.Walk(w)
	w.Write("::")
	w.Write(x.Type)
}

// HostValue produces non-SQL code that should produce an SQL value
// when executed in the host environment.
type HostValue struct {
	Value Node
}

func (v HostValue) Walk(w *ast.Writer) {
	v.Value.Walk(w)
}

type nooptype uint8

func (nooptype) Walk(w *ast.Writer) {}

func (ValueExprList) valueExprNode()        {}
func (ColumnReference) valueExprNode()      {}
func (DynamicParmeterSpec) valueExprNode()  {}
func (OrdinalParameterSpec) valueExprNode() {}
func (Literal) valueExprNode()              {}
func (RoutineInvocation) valueExprNode()    {}
func (QuantifiedExpr) valueExprNode()       {}
func (CastExpr) valueExprNode()             {}
func (HostValue) valueExprNode()            {}
func (nooptype) valueExprNode()             {}
