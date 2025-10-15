package internal

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

var (
	_ Flowable = RangeStmt{}
	_ Stmt     = RangeStmt{}
)

type RangeStmt struct {
	*ast.RangeStmt
	Key, Value Expr // Key, Value may be nil
	X          Expr
	Body       *BlockStmt
}

func (r RangeStmt) Eval(vm *VM) {
	rangeable := vm.returnsEval(r.X)
	vm.pushNewFrame()
	for i := 0; i < rangeable.Len(); i++ {
		if r.Key != nil {
			if ca, ok := r.Key.(CanAssign); ok {
				if i == 0 {
					ca.Define(vm, reflect.ValueOf(i))
				} else {
					ca.Assign(vm, reflect.ValueOf(i))
				}
			}
		}
		if r.Value != nil {
			if ca, ok := r.Value.(CanAssign); ok {
				if i == 0 {
					ca.Define(vm, rangeable.Index(i))
				} else {
					ca.Assign(vm, rangeable.Index(i))
				}
			}
		}
		vm.eval(r.Body)
	}
	vm.popFrame()
}

func (r RangeStmt) Flow(g *grapher) (head Step) {
	head = r.X.Flow(g)
	push := g.newPushStackFrame()
	g.nextStep(push)

	// index := 0
	indexVar := Ident{Ident: &ast.Ident{Name: "index"}}
	zeroInt := BasicLit{BasicLit: &ast.BasicLit{Kind: token.INT, Value: "0"}}
	assign := AssignStmt{
		AssignStmt: &ast.AssignStmt{
			Tok: token.DEFINE,
		},
		Lhs: []Expr{indexVar},
		Rhs: []Expr{zeroInt},
	}
	assign.Flow(g)

	// index < len(x)
	condition := BinaryExpr{
		BinaryExpr: &ast.BinaryExpr{
			Op: token.LSS,
		},
		X: indexVar,
		Y: CallExpr{Fun: Ident{Ident: &ast.Ident{Name: "len"}}, Args: []Expr{r.X}},
	}
	condition.Flow(g)

	// index++
	indexInc := IncDecStmt{
		IncDecStmt: &ast.IncDecStmt{
			Tok: token.INC,
		},
		X: indexVar,
	}
	indexInc.Flow(g)

	pop := g.newPopStackFrame()
	g.nextStep(pop)
	return
}

func (r RangeStmt) String() string {
	return fmt.Sprintf("RangeStmt(%v, %v, %v, %v)", r.Key, r.Value, r.X, r.Body)
}

func (r RangeStmt) stmtStep() Evaluable { return r }
