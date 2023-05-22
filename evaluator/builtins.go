package evaluator

import (
	"panda/object"
)

var builtins = map[string]*object.Builtin{
	"len":   object.GetBuiltinByName("len"),
	"first": object.GetBuiltinByName("first"),
	"last":  object.GetBuiltinByName("last"),
	"rest":  object.GetBuiltinByName("rest"),
	"shift": object.GetBuiltinByName("shift"),
	"push":  object.GetBuiltinByName("push"),
	"puts":  object.GetBuiltinByName("puts"),
}
