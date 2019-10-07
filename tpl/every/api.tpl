package api

// Get{{ .UpperCamelName}}RespAPI 查询{{ .Comment}} 单个
type Get{{ .UpperCamelName}}RespAPI struct {
	// 返回码
    Code int `json:"code"`
    // 返回信息
    Message string `json:"message"`
	// 内容
	Data val.Get{{ .UpperCamelName}}Resp `json:"data"`
}

// Get{{ .UpperCamelName}}sRespAPI 查询{{ .Comment}}列表 全部
type Get{{ .UpperCamelName}}sRespAPI struct {
	// 返回码
    Code int `json:"code"`
    // 返回信息
    Message string `json:"message"`
	// 内容
	Data val.Get{{ .UpperCamelName}}Resp `json:"data"`
}

// Get{{ .UpperCamelName}}sWithPageAPI 查询{{ .Comment}}列表 分页
type Get{{ .UpperCamelName}}sWithPageAPI struct {
	// 返回码
    Code int `json:"code"`
    // 返回信息
    Message string `json:"message"`
	// 内容
	Data {{ .UpperCamelName}}sPageModelAPI `json:"data"`
}

// {{ .UpperCamelName}}sPageModel 查询{{ .Comment}}分页内容
type {{ .UpperCamelName}}sPageModelAPI struct {
	// 页码
    PageNum int `json:"page_num"`
    // 每页大小
    PageSize int `json:"page_size"`
    // 总页数
    TotalPages int64 `json:"total_pages"`
    // 总记录条数
    TotalRecords int64 `json:"total_records"`
	// 数据
	Records      []val.Get{{ .UpperCamelName}}Resp `json:"records"`
}