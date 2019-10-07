package dao

// Add{{.UpperCamelName}} 新增{{.Comment}}
func Add{{.UpperCamelName}}({{.LowerCamelName}} *model.{{.UpperCamelName}}) (err error) {
    _, err = db.Engine.Insert({{.LowerCamelName}})
    return errors.WithStack(err)
}

// Remove{{.UpperCamelName}} 删除{{.Comment}}
func Remove{{.UpperCamelName}}(id int) (err error) {
	_, err = db.Engine.ID(id).Delete(&model.{{.UpperCamelName}}{})
    return errors.WithStack(err)
}

// Update{{.UpperCamelName}} 修改{{.Comment}}
func Update{{.UpperCamelName}}({{.LowerCamelName}} *model.{{.UpperCamelName}}) (err error) {
	_, err = db.Engine.ID({{.LowerCamelName}}.ID).Update({{.LowerCamelName}})
    return errors.WithStack(err)
}

// Get{{.UpperCamelName}} 查询{{.Comment}} 单个
func Get{{.UpperCamelName}}(id int) ({{.LowerCamelName}} *model.{{.UpperCamelName}}, err error) {
	{{.LowerCamelName}} = &model.{{.UpperCamelName}}{}
	exist, err := db.Engine.ID(id).Get({{.LowerCamelName}})
	if err != nil {
        return nil, errors.WithStack(err)
    }
    if !exist {
        {{.LowerCamelName}} = nil
    }
	return
}

// Get{{.UpperCamelName}}s 查询{{.Comment}}列表 全部
func Get{{.UpperCamelName}}s({{.LowerCamelName}} *model.{{.UpperCamelName}}) ({{.LowerCamelName}}s *[]model.{{.UpperCamelName}}, err error) {
	{{.LowerCamelName}}s = &[]model.{{.UpperCamelName}}{}
    err = db.Engine.Find({{.LowerCamelName}}s, {{.LowerCamelName}})
    err = errors.WithStack(err)
    return
}

// Get{{.UpperCamelName}}sPage 查询{{.Comment}} 分页
func Get{{.UpperCamelName}}sPage(page, pageSize int, {{.LowerCamelName}} *model.{{.UpperCamelName}}) ({{.LowerCamelName}}s *[]model.{{.UpperCamelName}}, total int64, err error) {
	{{.LowerCamelName}}s = &[]model.{{.UpperCamelName}}{}
    total, err = pager.Help(page, pageSize, {{.LowerCamelName}}s, {{.LowerCamelName}}, db.Engine.NewSession())
    err = errors.WithStack(err)
    return
}
