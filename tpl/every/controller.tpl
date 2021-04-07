package controller

// Add{{.UpperCamelName}} 新增{{.Comment}}
// @Summary 新增{{.Comment}}
// @Description 新增{{.Comment}}
// @Tags {{.UpperCamelName}}
// @Accept json
// @Produce json
// @Param data body val.Add{{.UpperCamelName}}Req true "{{.Comment}}"
// @Success 200 {object} api.BaseRespBody
// @Router /{{.UnderlineName | UnderlineStr2Strikethrough}} [post]
func Add{{.UpperCamelName}}(ctx *gin.Context) {
	add{{.UpperCamelName}}Req := &val.Add{{.UpperCamelName}}Req{}
    if httper.BindAndCheck(ctx, add{{.UpperCamelName}}Req) {
        return
    }

    err := service.Add{{.UpperCamelName}}(add{{.UpperCamelName}}Req)
    httper.HandleResponse(ctx, err, nil)
}

// Remove{{.UpperCamelName}} 删除{{.Comment}}
// @Summary 删除{{.Comment}}
// @Description 删除{{.Comment}}
// @Tags {{.UpperCamelName}}
// @Accept json
// @Produce json
// @Param data body val.Remove{{.UpperCamelName}}Req true "{{.Comment}}"
// @Success 200 {object} api.BaseRespBody
// @Router /{{.UnderlineName | UnderlineStr2Strikethrough}} [delete]
func Remove{{.UpperCamelName}}(ctx *gin.Context) {
    remove{{.UpperCamelName}}Req := &val.Remove{{.UpperCamelName}}Req{}
	if httper.BindAndCheck(ctx, remove{{.UpperCamelName}}Req) {
		return
	}

	err := service.Remove{{.UpperCamelName}}(remove{{.UpperCamelName}}Req.ID)
	httper.HandleResponse(ctx, err, nil)
}

// Update{{.UpperCamelName}} 修改{{.Comment}}
// @Summary 修改{{.Comment}}
// @Description 修改{{.Comment}}
// @Tags {{.UpperCamelName}}
// @Accept json
// @Produce json
// @Param data body val.Update{{.UpperCamelName}}Req true "{{.Comment}}"
// @Success 200 {object} api.BaseRespBody
// @Router /{{.UnderlineName | UnderlineStr2Strikethrough}} [put]
func Update{{.UpperCamelName}}(ctx *gin.Context) {
    update{{.UpperCamelName}}Req := &val.Update{{.UpperCamelName}}Req{}
	if httper.BindAndCheck(ctx, update{{.UpperCamelName}}Req) {
		return
	}

	err := service.Update{{.UpperCamelName}}(update{{.UpperCamelName}}Req)
	httper.HandleResponse(ctx, err, nil)
}

// Get{{.UpperCamelName}} 查询{{.Comment}} 单个
// @Summary 查询{{.Comment}} 单个
// @Description 查询{{.Comment}} 单个
// @Tags {{.UpperCamelName}}
// @Accept json
// @Produce json
// @Param id path int true "{{.Comment}}id"
// @Success 200 {object} httper.RespBody{data=val.Get{{.UpperCamelName}}Resp}
// @Router /{{.UnderlineName | UnderlineStr2Strikethrough}}/{id} [get]
func Get{{.UpperCamelName}}(ctx *gin.Context) {
    id, _ := strconv.Atoi(ctx.Param("id"))
    if id == 0 {
        httper.HandleResponse(ctx, myerr.NewParameterError("id为必填参数"), nil)
        return
    }
    
    resp, err := service.Get{{.UpperCamelName}}(id)
    httper.HandleResponse(ctx, err, resp)
}

// Get{{.UpperCamelName}}s 查询{{.Comment}}列表 全部
// @Summary 查询{{.Comment}}列表 全部
// @Description 查询{{.Comment}}列表 全部
// @Tags {{.UpperCamelName}}
// @Produce json
{{range .Fields -}}
    {{if and (ne .UpperCamelName "ID") (ne .UpperCamelName "CreatedAt") (ne .UpperCamelName "UpdatedAt") (ne .UpperCamelName "DeletedAt") -}}
    // @Param {{ .UnderlineName}} query string false "{{ .Comment}}"
    {{end}}
{{- end -}}
// @Success 200 {object} httper.RespBody{data=[]val.Get{{.UpperCamelName}}Resp}
// @Router /{{.UnderlineName | UnderlineStr2Strikethrough}}s [get]
func Get{{.UpperCamelName}}s(ctx *gin.Context) {
	get{{.UpperCamelName}}sReq := &val.Get{{.UpperCamelName}}sReq{}
	if httper.BindAndCheck(ctx, get{{.UpperCamelName}}sReq) {
		return
	}

	resp, err := service.Get{{.UpperCamelName}}s(get{{.UpperCamelName}}sReq)
	httper.HandleResponse(ctx, err, resp)
}

// Get{{.UpperCamelName}}sWithPage 查询{{.Comment}}列表 分页
// @Summary 查询{{.Comment}}列表 分页
// @Description 查询{{.Comment}}列表 分页
// @Tags {{.UpperCamelName}}
// @Produce json
// @Param page query int false "页码 默认为1"
// @Param page_size query int false "每页大小 默认为10"
{{range .Fields -}}
    {{if and (ne .UpperCamelName "ID") (ne .UpperCamelName "CreatedAt") (ne .UpperCamelName "UpdatedAt") (ne .UpperCamelName "DeletedAt") -}}
    // @Param {{ .UnderlineName}} query string false "{{ .Comment}}"
    {{end}}
{{- end -}}
// @Success 200 {object} httper.RespBody{data=pager.PageModel{records=[]val.Get{{.UpperCamelName}}Resp}}
// @Router /{{.UnderlineName | UnderlineStr2Strikethrough}}s/page [get]
func Get{{.UpperCamelName}}sWithPage(ctx *gin.Context) {
	get{{.UpperCamelName}}sWithPageReq := &val.Get{{.UpperCamelName}}sWithPageReq{}
	if httper.BindAndCheck(ctx, get{{.UpperCamelName}}sWithPageReq) {
		return
	}

	resp, err := service.Get{{.UpperCamelName}}sWithPage(get{{.UpperCamelName}}sWithPageReq)
	httper.HandleResponse(ctx, err, resp)
}
