package util

import (
	"go/ast"
)

// ReflectAccess 解析 ast 中对应类型的内容
func ReflectAccess(a ast.Expr) string {
	switch t := a.(type) {
	case *ast.StarExpr:
		identifier, ok := t.X.(*ast.Ident)
		if !ok {
			return ""
		}
		return identifier.Name
	case *ast.Ident:
		return t.Name
	}
	return ""
}
