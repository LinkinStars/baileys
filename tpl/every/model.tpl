package model

// {{ .UpperCamelName}} {{ .Comment}} 
type {{ .UpperCamelName}} struct {
{{range .Fields}}	{{ .UpperCamelName}}	{{.Type}} {{.ORMTag}}
{{end}}
}

