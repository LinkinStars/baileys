package entity

import "text/template"

// TplModel 生成模板时数据整合
type TplModel struct {
	Tpl             *template.Template // 模板
	Filename        string             //模板名称
	FilenameWithExt string             //模板名称包含后缀
	OutputPath      string             // 输出路径
	FilenameSuffix  string             // 输出文件名后缀如 _model -》 user_model.go
}
