package entity

// TableData 表的数模型
type TableData struct {
	UpperCamelName string // 大驼峰名称
	LowerCamelName string // 小驼峰名称
	UnderlineName  string // 下划线名称
	Comment        string // 注释
	Fields         []FieldData
}

// FieldData 字段的数据模型
type FieldData struct {
	UpperCamelName string // 大驼峰名称
	LowerCamelName string // 小驼峰名称
	UnderlineName  string // 下划线名称
	Type           string // 对应go的类型
	Comment        string // 字段注释
	ORMTag         string // orm框架的标签
	ValTag         string // 验证框架的标签
}
