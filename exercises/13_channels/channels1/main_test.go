package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"gostlings/internal/testutil"
	"testing"
	"time"
)

func TestMainUsesUnbufferedChannelWithGoroutineSend(t *testing.T) {
	file, err := parser.ParseFile(token.NewFileSet(), "main.go", nil, 0)
	if err != nil {
		t.Fatalf("parse main.go: %v", err)
	}

	var channelName string
	var hasGoroutineSend bool
	for _, declaration := range file.Decls {
		mainFunc, ok := declaration.(*ast.FuncDecl)
		if !ok || mainFunc.Name.Name != "main" {
			continue
		}

		for _, statement := range mainFunc.Body.List {
			assignment, ok := statement.(*ast.AssignStmt)
			if !ok || len(assignment.Lhs) != 1 || len(assignment.Rhs) != 1 {
				continue
			}
			name, ok := assignment.Lhs[0].(*ast.Ident)
			if ok && isUnbufferedStringChannel(assignment.Rhs[0]) {
				channelName = name.Name
			}
		}

		ast.Inspect(mainFunc.Body, func(node ast.Node) bool {
			goStatement, ok := node.(*ast.GoStmt)
			if !ok {
				return true
			}
			function, ok := goStatement.Call.Fun.(*ast.FuncLit)
			if !ok {
				return true
			}
			ast.Inspect(function.Body, func(node ast.Node) bool {
				send, ok := node.(*ast.SendStmt)
				if !ok {
					return true
				}
				name, ok := send.Chan.(*ast.Ident)
				if ok && name.Name == channelName {
					hasGoroutineSend = true
				}
				return true
			})
			return true
		})
	}

	if channelName == "" {
		t.Fatal("main() must create an unbuffered chan string")
	}
	if !hasGoroutineSend {
		t.Fatal("main() must send on its unbuffered channel from a goroutine")
	}
}

func isUnbufferedStringChannel(expression ast.Expr) bool {
	call, ok := expression.(*ast.CallExpr)
	if !ok || len(call.Args) != 1 {
		return false
	}
	makeName, ok := call.Fun.(*ast.Ident)
	if !ok || makeName.Name != "make" {
		return false
	}
	channel, ok := call.Args[0].(*ast.ChanType)
	if !ok || channel.Dir != ast.SEND|ast.RECV {
		return false
	}
	element, ok := channel.Value.(*ast.Ident)
	return ok && element.Name == "string"
}

func TestMainPrintsValueAfterUnbufferedHandoff(t *testing.T) {
	gotCh := make(chan string, 1)
	go func() { gotCh <- testutil.CaptureStdout(t, main) }()
	select {
	case got := <-gotCh:
		const want = "hi\n"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("main() did not finish; unbuffered handoff is still blocked")
	}
}
