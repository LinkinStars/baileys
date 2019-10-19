package model

// {{ .UpperCamelName}} {{ .Comment}} 
type {{ .UpperCamelName}} struct {
{{range .Fields}}	{{ .UpperCamelName}}	{{.Type}} {{.ORMTag}}
{{end}}
}

// TableName {{ .Comment}} 表名
func ({{ .UpperCamelName}}) TableName() string {
    return "{{ .UnderlineName}}"
}
