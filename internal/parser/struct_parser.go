package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

// StructFlat 非嵌套结构体
type StructFlat struct {
	Name   string
	Fields []*StructField
}

// StructField 结构体字段
type StructField struct {
	Name string
	Type string
}

const (
	InterfaceTypeDef = "interface"
	StructTypeDef    = "struct"
	TimeTypeDef      = "time.Time"
)

// StructParser golang struct 解析器
func StructParser(src string) (structList []*StructFlat, err error) {
	src = addPackageIfNotExist(src)
	fileSet := token.NewFileSet()
	f, err := parser.ParseFile(fileSet, "src.go", src, 0)
	if err != nil {
		return nil, err
	}

	structList = make([]*StructFlat, 0)
	for _, node := range f.Decls {
		switch node.(type) {
		case *ast.GenDecl:
			genDecl := node.(*ast.GenDecl)
			for _, spec := range genDecl.Specs {
				switch spec.(type) {
				case *ast.TypeSpec:
					typeSpec := spec.(*ast.TypeSpec)

					// 获取结构体名称
					structFlat := &StructFlat{Name: typeSpec.Name.Name}
					structFlat.Fields = make([]*StructField, 0)
					log.Printf("read struct %s\n", typeSpec.Name.Name)

					switch typeSpec.Type.(type) {
					case *ast.StructType:
						structType := typeSpec.Type.(*ast.StructType)
						for _, field := range structType.Fields.List {
							f := &StructField{}
							switch field.Type.(type) {
							case *ast.Ident:
								iDent := field.Type.(*ast.Ident)
								f.Type = iDent.Name
							case *ast.InterfaceType:
								f.Type = InterfaceTypeDef
							case *ast.MapType:
								iDent := field.Type.(*ast.MapType)
								f.Type = fmt.Sprintf("map[%s]%s", iDent.Key, iDent.Value)
							case *ast.ArrayType:
								iDent := field.Type.(*ast.ArrayType)
								f.Type = fmt.Sprintf("[]%s", iDent.Elt)
							case *ast.StructType:
								f.Type = StructTypeDef
							case *ast.SelectorExpr:
								iDent := field.Type.(*ast.SelectorExpr)
								if iDent.Sel.Name == "Time" {
									f.Type = TimeTypeDef
								} else {
									log.Printf("undefined field type %+v", field.Type)
								}
							default:
								log.Printf("undefined field type %+v", field.Type)
							}

							for _, name := range field.Names {
								f.Name = name.Name
								structFlat.Fields = append(structFlat.Fields, f)
								log.Printf("name=%s type=%s\n", name.Name, f.Type)
							}
						}
					}
					structList = append(structList, structFlat)
				}
			}
		}
	}
	return structList, nil
}

func addPackageIfNotExist(src string) string {
	if strings.HasPrefix(src, "package") {
		return src
	}
	return "package mypackage\n" + src
}
