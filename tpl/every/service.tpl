package service

// Add{{.UpperCamelName}} 新增{{.Comment}}
func Add{{.UpperCamelName}}(req *val.Add{{.UpperCamelName}}Req) (err error) {
    {{.LowerCamelName}} := &model.{{.UpperCamelName}}{}
    _ = copier.Copy({{.LowerCamelName}}, req)
    return dao.Add{{.UpperCamelName}}({{.LowerCamelName}})
}

// Remove{{.UpperCamelName}} 删除{{.Comment}}
func Remove{{.UpperCamelName}}(id int) (err error) {
    return dao.Remove{{.UpperCamelName}}(id)
}

// Update{{.UpperCamelName}} 修改{{.Comment}}
func Update{{.UpperCamelName}}(req *val.Update{{.UpperCamelName}}Req) (err error) {
    {{.LowerCamelName}} := &model.{{.UpperCamelName}}{}
    _ = copier.Copy({{.LowerCamelName}}, req)
    return dao.Update{{.UpperCamelName}}({{.LowerCamelName}})
}

// Get{{.UpperCamelName}} 查询{{.Comment}} 单个
func Get{{.UpperCamelName}}(id int) (req *val.Get{{.UpperCamelName}}Resp, err error) {
	{{.LowerCamelName}}, err := dao.Get{{.UpperCamelName}}(id)
    if err != nil || {{.LowerCamelName}} == nil {
        return
    }

    req = &val.Get{{.UpperCamelName}}Resp{}
    _ = copier.Copy(req, {{.LowerCamelName}})
    return req, nil
}

// Get{{.UpperCamelName}}s 查询{{.Comment}}列表 全部
func Get{{.UpperCamelName}}s(req *val.Get{{.UpperCamelName}}sReq) ({{.LowerCamelName}}sResp *[]val.Get{{.UpperCamelName}}Resp, err error) {
	{{.LowerCamelName}} := &model.{{.UpperCamelName}}{}
	_ = copier.Copy({{.LowerCamelName}}, req)

	{{.LowerCamelName}}s, err := dao.Get{{.UpperCamelName}}s({{.LowerCamelName}})
	if err != nil {
		return
	}

	{{.LowerCamelName}}sResp = &[]val.Get{{.UpperCamelName}}Resp{}
	_ = copier.Copy({{.LowerCamelName}}sResp, {{.LowerCamelName}}s)
	return
}

// Get{{.UpperCamelName}}sWithPage 查询{{.Comment}}列表 分页
func Get{{.UpperCamelName}}sWithPage(req *val.Get{{.UpperCamelName}}sWithPageReq) (pageModel *pager.PageModel, err error) {
	{{.LowerCamelName}} := &model.{{.UpperCamelName}}{}
	_ = copier.Copy({{.LowerCamelName}}, req)

	page := req.Page
	pageSize := req.PageSize
	
	{{.LowerCamelName}}s, total, err := dao.Get{{.UpperCamelName}}sPage(page, pageSize, {{.LowerCamelName}})
	if err != nil {
		return
	}
	
	{{.LowerCamelName}}sResp := &[]val.Get{{.UpperCamelName}}Resp{}
	_ = copier.Copy({{.LowerCamelName}}sResp, {{.LowerCamelName}}s)
	
	return pager.NewPageModel(page, pageSize, total, {{.LowerCamelName}}sResp), nil
}