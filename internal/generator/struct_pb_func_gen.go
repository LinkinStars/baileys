package generator

import (
	"bytes"
	"log"
	"strings"
	"text/template"

	"github.com/LinkinStars/baileys/internal/parsing"
)

const (
	struct2PBFuncTpl = `{{if and (ne .Comment "")}}// Convert{{.Name}}2PB {{ .Comment}}
{{end -}}func Convert{{.Name}}2PB(goModel *baileys.{{.Name}}) (pbModel *baileys.{{.Name}}PB) { 
    pbModel = &baileys.{{.Name}}PB{} {{range .Fields}}
    {{if eq .Type "bool" -}}
    pbModel.{{.Name}} = goModel.{{.Name}}
    {{- else if eq .Type "int" -}}
    pbModel.{{.Name}} = int32(goModel.{{.Name}})
    {{- else if eq .Type "int8" -}}
    pbModel.{{.Name}} = int32(goModel.{{.Name}})
    {{- else if eq .Type "int16" -}}
    pbModel.{{.Name}} = int32(goModel.{{.Name}})
    {{- else if eq .Type "int32" -}}
    pbModel.{{.Name}} = int32(goModel.{{.Name}})
    {{- else if eq .Type "int64" -}}
    pbModel.{{.Name}} = int64(goModel.{{.Name}})
    {{- else if eq .Type "uint" -}}
    pbModel.{{.Name}} = int32(goModel.{{.Name}})
    {{- else if eq .Type "uint8" -}}
    pbModel.{{.Name}} = int32(goModel.{{.Name}})
    {{- else if eq .Type "uint16" -}}
    pbModel.{{.Name}} = int32(goModel.{{.Name}})
    {{- else if eq .Type "uint32" -}}
    pbModel.{{.Name}} = int32(goModel.{{.Name}})
    {{- else if eq .Type "uint64" -}}
    pbModel.{{.Name}} = int64(goModel.{{.Name}})
    {{- else if eq .Type "uintptr" -}}
    pbModel.{{.Name}} = goModel.{{.Name}}
    {{- else if eq .Type "float32" -}}
    pbModel.{{.Name}} = goModel.{{.Name}}
    {{- else if eq .Type "float64" -}}
    pbModel.{{.Name}} = goModel.{{.Name}}
    {{- else if eq .Type "complex64" -}}
    pbModel.{{.Name}} = goModel.{{.Name}}
    {{- else if eq .Type "complex128" -}}
    pbModel.{{.Name}} = goModel.{{.Name}}
    {{- else if eq .Type "interface{}" -}}
    pbModel.{{.Name}} = goModel.{{.Name}}
    {{- else if eq .Type "map[string]string" -}}
    pbModel.{{.Name}} = goModel.{{.Name}}
    {{- else if eq .Type "string" -}}
    pbModel.{{.Name}} = goModel.{{.Name}}
    {{- else if eq .Type "[]string" -}}
    pbModel.{{.Name}} = goModel.{{.Name}}
    {{- else if eq .Type "struct{}" -}}
    pbModel.{{.Name}} = goModel.{{.Name}}
    {{- else if eq .Type "time.Time" -}}
    pbModel.{{.Name}} = timestamppb.New(goModel.{{.Name}})
    {{- else -}}
    pbModel.{{.Name}} = goModel.{{.Name}}
    {{- end}}
{{- end}}
    return pbModel
}
`
)

// GenerateStruct2PBFunc 生成 golang struct 转换为 protobuf 的方法
func GenerateStruct2PBFunc(structList []*parsing.StructFlat) (res string, err error) {
	funcs := map[string]interface{}{
		"contains": strings.Contains,
	}
	t, err := template.New("struct2PBFuncTpl.tpl").Funcs(funcs).Parse(struct2PBFuncTpl)
	if err != nil {
		log.Printf("could not parse template: %s\n", err.Error())
		return "", err
	}

	for _, s := range structList {
		resBytes := bytes.NewBufferString("")
		err := t.Execute(resBytes, s)
		if err != nil {
			log.Printf("could not generate %s", err.Error())
		}
		resBytes.WriteString("\n")
		res += resBytes.String()
	}
	return
}
