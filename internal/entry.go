package internal

import (
	"go/ast"
	"log"
)

// corresponding Field, XxxSpec, FuncDecl, LabeledStmt, AssignStmt, Scope; or nil
func DoDecl(decl any) {
	f, ok := decl.(*ast.FuncDecl)
	if !ok {
		log.Println(decl)
		return
	}
	b := builder{}
	b.Visit(f.Body)
}
