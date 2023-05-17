package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"panda/ast"
	"panda/code"
	"strings"
)

type Type string

type BuiltinFunction func(args ...Object) Object

const (
	IntegerObj = "Integer"
	BooleanObj = "Boolean"
	NullObj    = "NULL"

	ReturnValueObj = "ReturnValue"

	FunctionObj         = "Function"
	CompiledFunctionObj = "CompiledFunction"

	StringObj  = "String"
	ArrayObj   = "Array"
	HashObj    = "Hash"
	BuiltinObj = "Builtin"

	ErrorObj = "Error"
)

type Object interface {
	Type() Type
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() Type {
	return IntegerObj
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() Type {
	return BooleanObj
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

type Null struct{}

func (n *Null) Type() Type {
	return NullObj
}

func (n *Null) Inspect() string {
	return "null"
}

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() Type {
	return ReturnValueObj
}

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() Type {
	return FunctionObj
}

func (f *Function) Inspect() string {
	var out bytes.Buffer

	var params []string
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type CompiledFunction struct {
	Instructions  code.Instructions
	LocalsNum     int
	ParametersNum int
}

func (cf *CompiledFunction) Type() Type {
	return CompiledFunctionObj
}

func (cf *CompiledFunction) Inspect() string {
	return fmt.Sprintf("CompiledFunction[%p]", cf)
}

type String struct {
	Value string
}

func (s *String) Type() Type {
	return StringObj
}

func (s *String) Inspect() string {
	return s.Value
}

type Array struct {
	Elements []Object
}

func (a *Array) Type() Type {
	return ArrayObj
}

func (a *Array) Inspect() string {
	var out bytes.Buffer

	var elements []string
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ","))
	out.WriteString("]")

	return out.String()
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() Type {
	return HashObj
}

func (h *Hash) Inspect() string {
	var out bytes.Buffer

	paris := []string{}
	for _, pair := range h.Pairs {
		paris = append(paris, fmt.Sprintf("%s, %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(paris, ", "))
	out.WriteString("}")

	return out.String()
}

type Error struct {
	Message string
}

func (e *Error) Type() Type {
	return ErrorObj
}

func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() Type {
	return BuiltinObj
}

func (b *Builtin) Inspect() string {
	return "builtin function"
}

type HashKey struct {
	Type  Type
	Value uint64
}

type Hashable interface {
	HashKey() HashKey
}

func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	_, err := h.Write([]byte(s.Value))
	if err != nil {
		return HashKey{}
	}

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}
