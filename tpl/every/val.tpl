package val

// Add{{ .UpperCamelName}}Req 新增{{ .Comment}}请求结构 
type Add{{ .UpperCamelName}}Req struct {
{{range .Fields -}}
    {{if and (ne .UpperCamelName "ID") (ne .UpperCamelName "CreatedAt") (ne .UpperCamelName "UpdatedAt") (ne .UpperCamelName "DeletedAt") -}}
        // {{ .Comment}}
        {{ .UpperCamelName}}	{{.Type}} `{{.ValTag}} json:"{{ .UnderlineName}}"`
    {{else -}}
    {{end}}
{{- end}}
}

// Remove{{ .UpperCamelName}}Req 删除{{ .Comment}}请求结构 
type Remove{{ .UpperCamelName}}Req struct {
    {{range .Fields -}}
        {{if eq .UpperCamelName "ID" -}}
            // {{ .Comment}}
            {{ .UpperCamelName}}	{{.Type}} `{{.ValTag}} json:"{{ .UnderlineName}}"`
        {{end}}
    {{- end}}
}

// Update{{ .UpperCamelName}}Req 修改{{ .Comment}}请求结构 
type Update{{ .UpperCamelName}}Req struct {
    {{range .Fields -}} 
        {{if and (ne .UpperCamelName "CreatedAt") (ne .UpperCamelName "UpdatedAt") (ne .UpperCamelName "DeletedAt") -}}
            // {{ .Comment}}
            {{if eq .UpperCamelName "ID" -}}
                {{ .UpperCamelName}}	{{.Type}} `{{.ValTag}} json:"{{ .UnderlineName}}"`
            {{else -}}
                {{ .UpperCamelName}}	{{.Type}} `{{.ValTag | ChangeValTagForUpdate}} json:"{{ .UnderlineName}}"`
            {{end}}
        {{- end}}
    {{- end}}
}

// Get{{ .UpperCamelName}}sReq 查询{{ .Comment}}列表 全部 请求结构
type Get{{ .UpperCamelName}}sReq struct {
	{{range .Fields -}}
	    {{if and (ne .UpperCamelName "ID") (ne .UpperCamelName "CreatedAt") (ne .UpperCamelName "UpdatedAt") (ne .UpperCamelName "DeletedAt") -}}
	        // {{ .Comment}}
            {{ .UpperCamelName}}	{{.Type}} `{{.ValTag | ChangeValTagForUpdate}} form:"{{ .UnderlineName}}"`
        {{end}}
    {{- end}}
}

// Get{{ .UpperCamelName}}sWithPageReq 查询{{ .Comment}}列表 分页 请求结构
type Get{{ .UpperCamelName}}sWithPageReq struct {
	// 页码
	Page int `validate:"required,min=1" comment:"页码"`
	// 每页大小
	PageSize int `validate:"required,min=1" comment:"每页大小"`
	{{range .Fields -}}
	    {{if and (ne .UpperCamelName "ID") (ne .UpperCamelName "CreatedAt") (ne .UpperCamelName "UpdatedAt") (ne .UpperCamelName "DeletedAt") -}}
	        // {{ .Comment}}
            {{ .UpperCamelName}}	{{.Type}} `{{.ValTag | ChangeValTagForUpdate}} form:"{{ .UnderlineName}}"`
        {{end}}
    {{- end}}
}

// Get{{ .UpperCamelName}}Resp 查询{{ .Comment}} 返回结构 
type Get{{ .UpperCamelName}}Resp struct {
    {{range .Fields -}}
        {{if and (ne .UpperCamelName "DeletedAt") -}}
            // {{ .Comment}}
            {{ .UpperCamelName}}	{{.Type}} `json:"{{.UnderlineName}}"`
        {{end}}
    {{- end}}
}
