package sqlang

import (
	"strconv"

	"github.com/frk/ast"
)

var (
	DEFAULT   DefaultMarker
	NULL      Null
	ISNULL    = IsX{NULL}
	ISNOTNULL = IsNotX{NULL}
)

type NoOp struct{}

func (NoOp) Walk(w *ast.Writer) {}

type SpaceChar struct{}

func (SpaceChar) Walk(w *ast.Writer) {
	w.Write(" ")
}

type NewLine struct{}

func (NewLine) Walk(w *ast.Writer) {
	w.NewLine()
}

type DefaultMarker struct{}

func (DefaultMarker) Walk(w *ast.Writer) {
	w.Write("DEFAULT")
}

type Null struct{}

func (Null) Walk(w *ast.Writer) {
	w.Write("NULL")
}

type IsX struct {
	X Expr
}

func (x IsX) Walk(w *ast.Writer) {
	w.Write(" IS ")
	x.X.Walk(w)
}

type IsNotX struct {
	X Expr
}

func (x IsNotX) Walk(w *ast.Writer) {
	w.Write(" IS NOT ")
	x.X.Walk(w)
}

type Keyword string

func (k Keyword) Walk(w *ast.Writer) {
	w.Write(string(k))
}

type Name string

func (n Name) Walk(w *ast.Writer) {
	w.Write(`"`)
	w.Write(string(n))
	w.Write(`"`)
}

type NameSlice []Name

func (ns NameSlice) Walk(w *ast.Writer) {
	w.Write("(")
	for i, c := range ns {
		w.NewLine()
		if i > 0 {
			w.Write(", ")
		}
		c.Walk(w)
	}
	w.NewLine()
	w.Write(")")
}

type Ident struct {
	Name  Name
	Qual  string
	Alias string
}

func (id Ident) Walk(w *ast.Writer) {
	if len(id.Qual) > 0 {
		w.Write(id.Qual)
		w.Write(".")
	}
	id.Name.Walk(w)

	if len(id.Alias) > 0 {
		w.Write(" AS ")
		w.Write(id.Alias)
	}
}

type ColumnIdent struct {
	Name Name
	Qual string
}

func (id ColumnIdent) Walk(w *ast.Writer) {
	if len(id.Qual) > 0 {
		w.Write(id.Qual)
		w.Write(".")
	}
	id.Name.Walk(w)
}

type ColumnIdentSlice []ColumnIdent

func (ids ColumnIdentSlice) Walk(w *ast.Writer) {
	for i, id := range ids {
		if i > 0 {
			w.NewLine()
			w.Write(", ")
		}
		id.Walk(w)
	}
}

type ColumnExpr interface {
	Node
	columnExprNode()
}

type ColumnExprSlice []ColumnExpr

func (exprs ColumnExprSlice) Walk(w *ast.Writer) {
	for i, x := range exprs {
		if i > 0 {
			w.NewLine()
			w.Write(", ")
		}
		x.Walk(w)
	}
}

type Coalesce struct {
	A ValueExpr
	B ValueExpr
	// For the majority of COALESCE expressions the two arguments A and B
	// should be enough, but in case more are needed the Additional slice
	// can be used.
	Additional []ValueExpr
}

func (c Coalesce) Walk(w *ast.Writer) {
	w.Write("COALESCE(")
	c.A.Walk(w)
	w.Write(", ")
	c.B.Walk(w)
	for _, x := range c.Additional {
		w.Write(", ")
		x.Walk(w)
	}
	w.Write(")")
}

type NULLIF struct {
	Value ValueExpr
	Expr  ValueExpr
}

func (n NULLIF) Walk(w *ast.Writer) {
	w.Write("NULLIF(")
	n.Value.Walk(w)
	w.Write(", ")
	n.Expr.Walk(w)
	w.Write(")")
}

type PositionalParameter int

func (o PositionalParameter) Walk(w *ast.Writer) {
	w.Write("$")
	w.Write(strconv.Itoa(int(o)))
}

type NodeSlice []Node

func (slice NodeSlice) Walk(w *ast.Writer) {
	for _, n := range slice {
		n.Walk(w)
	}
}

type NodeList []Node

func (list NodeList) Walk(w *ast.Writer) {
	for i, n := range list {
		if i > 0 {
			w.NewLine()
		}
		n.Walk(w)
	}
}

type BINARY_OP string

const (
	BINARY_ADD BINARY_OP = "+"
	BINARY_SUB BINARY_OP = "-"
)

type BinaryExpr struct {
	X  Expr
	Op BINARY_OP
	Y  Expr
}

func (x BinaryExpr) Walk(w *ast.Writer) {
	x.X.Walk(w)
	w.Write(" ")
	w.Write(string(x.Op))
	w.Write(" ")
	x.Y.Walk(w)
}

func (NoOp) exprNode()                {}
func (Literal) exprNode()             {}
func (ColumnIdent) exprNode()         {}
func (DefaultMarker) exprNode()       {}
func (PositionalParameter) exprNode() {}
func (Coalesce) exprNode()            {}
func (NULLIF) exprNode()              {}
func (IsX) exprNode()                 {}
func (IsNotX) exprNode()              {}
func (Null) exprNode()                {}
func (BinaryExpr) exprNode()          {}

func (ColumnIdent) columnExprNode() {}
func (Coalesce) columnExprNode()    {}

func (DefaultMarker) valueExprNode() {}
func (ColumnIdent) valueExprNode()   {}
func (Coalesce) valueExprNode()      {}
func (NULLIF) valueExprNode()        {}
