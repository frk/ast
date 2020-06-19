package golang

import (
	"bytes"
	"io"

	"github.com/frk/ast"
)

type Node interface {
	Walk(w *ast.Writer)
}

func Write(n Node, w io.Writer) error {
	out := ast.NewWriter(w)
	n.Walk(out)
	return out.Err()
}

func ToString(n Node) (string, error) {
	b := new(bytes.Buffer)
	if err := Write(n, b); err != nil {
		return "", nil
	}
	return b.String(), nil
}
