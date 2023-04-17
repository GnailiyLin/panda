package evaluator

import (
	"fmt"
	"panda/ast"
	"panda/object"
	"panda/token"
)

func quote(node ast.Node, env *object.Environment) object.Object {
	node = evalUnquoteCalls(node, env)
	return &object.Quote{Node: node}
}

func evalUnquoteCalls(quoted ast.Node, env *object.Environment) ast.Node {
	return ast.Modify(quoted, func(node ast.Node) ast.Node {
		if !isUnquoteCall(node) {
			return node
		}

		call, ok := node.(*ast.CallExpression)
		if !ok {
			return node
		}

		if len(call.Arguments) != 1 {
			return node
		}

		unquoted := Eval(call.Arguments[0], env)
		return convertObjectToASTNode(unquoted)
	})
}

func isUnquoteCall(node ast.Node) bool {
	callExpression, ok := node.(*ast.CallExpression)
	if !ok {
		return false
	}

	return callExpression.Function.TokenLiteral() == "unquote"
}

func convertObjectToASTNode(obj object.Object) ast.Node {
	switch obj := obj.(type) {
	case *object.Integer:
		tok := token.Token{
			Type:    token.Int,
			Literal: fmt.Sprintf("%d", obj.Value),
		}
		return &ast.IntegerLiteral{Token: tok, Value: obj.Value}

	case *object.Boolean:
		var tok token.Token
		if obj.Value {
			tok = token.Token{Type: token.True, Literal: "true"}
		} else {
			tok = token.Token{Type: token.False, Literal: "false"}
		}
		return &ast.Boolean{Token: tok, Value: obj.Value}

	case *object.Quote:
		return obj.Node

	default:
		return nil
	}
}
