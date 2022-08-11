package generator

import (
	"bytes"
	"log"
	"text/template"

	"github.com/LinkinStars/baileys/internal/converter"
)

const (
	pbTpl = `{{if and (ne .Comment "")}}// {{ .Comment}}
{{end -}}message {{.Name}} { {{range .PBFieldList}}
    {{if and (ne .Comment "") -}}
    // {{ .Comment}}
    {{ .Type}} {{.Name}} = {{ .Index }};
    {{- else -}}
    {{ .Type}} {{.Name}} = {{ .Index }};
    {{- end}}
{{- end}}
}
`
)

// GenPBMessage 生成 pb message 结构
func GenPBMessage(flatList []*converter.PBFlat) (res string, err error) {
	t, err := template.New("pb.tpl").Parse(pbTpl)
	if err != nil {
		log.Printf("could not parse template: %s\n", err.Error())
		return "", err

	}
	resBytes := bytes.NewBufferString("")
	for _, flat := range flatList {
		err := t.Execute(resBytes, flat)
		if err != nil {
			log.Printf("could not generate %s", err.Error())
		}
		resBytes.WriteString("\n")
	}
	return resBytes.String(), nil
}
