package util

import (
	"strings"

	"xorm.io/core"
)

var (
	mapper core.IMapper = SpecialMapper
)

// SQLStr2GoStr 数据库字段、表名转换为大驼峰的golang命名，并且会替换一些特殊的映射规则如：ID
func SQLStr2GoStr(str string) string {
	return mapper.Table2Obj(str)
}

// ChangeValTagForUpdate 因为修改的时候字段均不是必填项
func ChangeValTagForUpdate(tag string) string {
	return strings.ReplaceAll(tag, "required", "omitempty")
}
