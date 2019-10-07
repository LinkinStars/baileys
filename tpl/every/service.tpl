package service

// Add{{.UpperCamelName}} 新增{{.Comment}}
func Add{{.UpperCamelName}}(add{{.UpperCamelName}}Req *val.Add{{.UpperCamelName}}Req) (err error) {
    {{.LowerCamelName}} := &model.{{.UpperCamelName}}{}
    _ = copier.Copy({{.LowerCamelName}}, add{{.UpperCamelName}}Req)
    return dao.Add{{.UpperCamelName}}({{.LowerCamelName}})
}

// Remove{{.UpperCamelName}} 删除{{.Comment}}
func Remove{{.UpperCamelName}}(id int) (err error) {
    return dao.Remove{{.UpperCamelName}}(id)
}

// Update{{.UpperCamelName}} 修改{{.Comment}}
func Update{{.UpperCamelName}}(update{{.UpperCamelName}}Req *val.Update{{.UpperCamelName}}Req) (err error) {
    {{.LowerCamelName}} := &model.{{.UpperCamelName}}{}
    _ = copier.Copy({{.LowerCamelName}}, update{{.UpperCamelName}}Req)
    return dao.Update{{.UpperCamelName}}({{.LowerCamelName}})
}

// Get{{.UpperCamelName}} 查询{{.Comment}} 单个
func Get{{.UpperCamelName}}(id int) (get{{.UpperCamelName}}Resp *val.Get{{.UpperCamelName}}Resp, err error) {
	{{.LowerCamelName}}, err := dao.Get{{.UpperCamelName}}(id)
    if err != nil || {{.LowerCamelName}} == nil {
        return
    }

    get{{.UpperCamelName}}Resp = &val.Get{{.UpperCamelName}}Resp{}
    _ = copier.Copy(get{{.UpperCamelName}}Resp, {{.LowerCamelName}})
    return get{{.UpperCamelName}}Resp, nil
}

// Get{{.UpperCamelName}}s 查询{{.Comment}}列表 全部
func Get{{.UpperCamelName}}s(get{{.UpperCamelName}}sReq *val.Get{{.UpperCamelName}}sReq) ({{.LowerCamelName}}sResp *[]val.Get{{.UpperCamelName}}Resp, err error) {
	{{.LowerCamelName}} := &model.{{.UpperCamelName}}{}
	_ = copier.Copy({{.LowerCamelName}}, get{{.UpperCamelName}}sReq)

	{{.LowerCamelName}}s, err := dao.Get{{.UpperCamelName}}s({{.LowerCamelName}})
	if err != nil {
		return
	}

	{{.LowerCamelName}}sResp = &[]val.Get{{.UpperCamelName}}Resp{}
	_ = copier.Copy({{.LowerCamelName}}sResp, {{.LowerCamelName}}s)
	return
}

// Get{{.UpperCamelName}}sWithPage 查询{{.Comment}}列表 分页
func Get{{.UpperCamelName}}sWithPage(get{{.UpperCamelName}}sWithPageReq *val.Get{{.UpperCamelName}}sWithPageReq) (pageModel *pager.PageModel, err error) {
	{{.LowerCamelName}} := &model.{{.UpperCamelName}}{}
	_ = copier.Copy({{.LowerCamelName}}, get{{.UpperCamelName}}sWithPageReq)

	page := get{{.UpperCamelName}}sWithPageReq.Page
	pageSize := get{{.UpperCamelName}}sWithPageReq.PageSize
	
	{{.LowerCamelName}}s, total, err := dao.Get{{.UpperCamelName}}sPage(page, pageSize, {{.LowerCamelName}})
	if err != nil {
		return
	}
	
	{{.LowerCamelName}}sResp := &[]val.Get{{.UpperCamelName}}Resp{}
	_ = copier.Copy({{.LowerCamelName}}sResp, {{.LowerCamelName}}s)
	
	return pager.NewPageModel(page, pageSize, total, {{.LowerCamelName}}sResp), nil
}