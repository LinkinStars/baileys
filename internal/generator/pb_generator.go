package generator

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	"github.com/LinkinStars/baileys/internal/converter"
)

const (
	pbTpl = `message {{.Name}} {
{{range .PBFieldList}}  {{ .Type}} {{.Name}} = {{ .Index }};
{{end}}}
`
)

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
	fmt.Println(resBytes.String())
	return resBytes.String(), nil
}
