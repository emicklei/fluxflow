package internal

import (
	"fmt"
	"go/ast"
	"go/token"
)

var _ Stmt = AssignStmt{}

type AssignStmt struct {
	*ast.AssignStmt
	Lhs []Expr
	Rhs []Expr
}

func (a AssignStmt) stmtStep() Evaluable { return a }

func (a AssignStmt) Eval(vm *VM) {
	if !vm.isStepping {
		// when stepping, the rhs are already evaluated
		// so here we need to eval each rhs to push values onto the operand stack
		for _, each := range a.Rhs {
			// values are on operand stack
			vm.eval(each)
		}
	}
	// operands are stacked in reverse order
	for i := len(a.Lhs) - 1; i != -1; i-- {
		each := a.Lhs[i]
		if trace {
			if len(vm.callStack.top().operandStack) == 0 {
				vm.fatal("operand stack empty before assignment")
			}
		}
		v := vm.callStack.top().pop()
		target, ok_ := each.(CanAssign)
		if !ok_ {
			vm.fatal("cannot assign to " + fmt.Sprintf("%T", each))
		}
		switch a.AssignStmt.Tok {
		case token.DEFINE: // :=
			target.Define(vm, v)
		case token.ASSIGN: // =
			target.Assign(vm, v)
		case token.ADD_ASSIGN: // +=
			current := vm.returnsEval(each)
			result := BinaryExprValue{left: current, op: token.ADD, right: v}.Eval()
			target.Assign(vm, result)
		case token.SUB_ASSIGN: // -=
			current := vm.returnsEval(each)
			result := BinaryExprValue{left: current, op: token.SUB, right: v}.Eval()
			target.Assign(vm, result)
		case token.MUL_ASSIGN: // *=
			current := vm.returnsEval(each)
			result := BinaryExprValue{left: current, op: token.MUL, right: v}.Eval()
			target.Assign(vm, result)
		case token.QUO_ASSIGN: // /=
			current := vm.returnsEval(each)
			result := BinaryExprValue{left: current, op: token.QUO, right: v}.Eval()
			target.Assign(vm, result)
		case token.REM_ASSIGN: // %=
			current := vm.returnsEval(each)
			result := BinaryExprValue{left: current, op: token.REM, right: v}.Eval()
			target.Assign(vm, result)
		case token.AND_ASSIGN: // &=
			current := vm.returnsEval(each)
			result := BinaryExprValue{left: current, op: token.AND, right: v}.Eval()
			target.Assign(vm, result)
		case token.OR_ASSIGN: // |=
			current := vm.returnsEval(each)
			result := BinaryExprValue{left: current, op: token.OR, right: v}.Eval()
			target.Assign(vm, result)
		case token.XOR_ASSIGN: // ^=
			current := vm.returnsEval(each)
			result := BinaryExprValue{left: current, op: token.XOR, right: v}.Eval()
			target.Assign(vm, result)
		case token.SHL_ASSIGN: // <<=
			current := vm.returnsEval(each)
			result := BinaryExprValue{left: current, op: token.SHL, right: v}.Eval()
			target.Assign(vm, result)
		case token.SHR_ASSIGN: // >>=
			current := vm.returnsEval(each)
			result := BinaryExprValue{left: current, op: token.SHR, right: v}.Eval()
			target.Assign(vm, result)
		case token.AND_NOT_ASSIGN: // &^=
			current := vm.returnsEval(each)
			result := BinaryExprValue{left: current, op: token.AND_NOT, right: v}.Eval()
			target.Assign(vm, result)
		default:
			panic("unsupported assignment " + a.AssignStmt.Tok.String())
		}
	}
	if len(a.Lhs) < len(a.Rhs) {
		_ = vm.callStack.top().pop()
	}
}
func (a AssignStmt) String() string {
	return fmt.Sprintf("AssignStmt(%v %s %v)", a.Lhs, a.AssignStmt.Tok, a.Rhs)
}

func (a AssignStmt) Flow(g *grapher) (head Step) {
	for i, each := range a.Rhs {
		if i == 0 {
			head = each.Flow(g)
			continue
		}
		each.Flow(g)
	}
	g.next(a)
	if head == nil {
		head = g.current
	}
	return head
}
