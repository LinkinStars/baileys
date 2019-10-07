package util

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"xorm.io/core"
)

var (
	created, updated, deleted              = []string{"created_at"}, []string{"updated_at"}, []string{"deleted_at"}
	mapper                    core.IMapper = SpecialMapper
)

// 数据库字段、表名转换为大驼峰的golang命名，并且会替换一些特殊的映射规则如：ID
func SqlStr2GoStr(str string) string {
	return mapper.Table2Obj(str)
}

func CreateTypeString(col *core.Column) string {
	st := col.SQLType
	t := core.SQLType2Type(st)
	s := t.String()
	if s == "[]uint8" {
		return "[]byte"
	}
	return s
}

func CreateXORMTag(table *core.Table, col *core.Column) string {
	var res []string
	if !col.Nullable {
		res = append(res, "not null")
	}
	if col.IsPrimaryKey {
		res = append(res, "pk")
	}
	if col.Default != "" {
		res = append(res, "default "+col.Default)
	}
	if col.IsAutoIncrement {
		res = append(res, "autoincr")
	}

	if col.SQLType.IsTime() && include(created, col.Name) {
		res = append(res, "created")
	}

	if col.SQLType.IsTime() && include(updated, col.Name) {
		res = append(res, "updated")
	}

	if col.SQLType.IsTime() && include(deleted, col.Name) {
		res = append(res, "deleted")
	}

	if col.Comment != "" {
		res = append(res, fmt.Sprintf("comment('%s')", col.Comment))
	}

	names := make([]string, 0, len(col.Indexes))
	for name := range col.Indexes {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		index := table.Indexes[name]
		var uistr string
		if index.Type == core.UniqueType {
			uistr = "unique"
		} else if index.Type == core.IndexType {
			uistr = "index"
		}
		if len(index.Cols) > 1 {
			uistr += "(" + index.Name + ")"
		}
		res = append(res, uistr)
	}

	nstr := col.SQLType.Name
	if col.Length != 0 {
		if col.Length2 != 0 {
			nstr += fmt.Sprintf("(%v,%v)", col.Length, col.Length2)
		} else {
			nstr += fmt.Sprintf("(%v)", col.Length)
		}
	} else if len(col.EnumOptions) > 0 { //enum
		nstr += "("
		opts := ""

		enumOptions := make([]string, 0, len(col.EnumOptions))
		for enumOption := range col.EnumOptions {
			enumOptions = append(enumOptions, enumOption)
		}
		sort.Strings(enumOptions)

		for _, v := range enumOptions {
			opts += fmt.Sprintf(",'%v'", v)
		}
		nstr += strings.TrimLeft(opts, ",")
		nstr += ")"
	} else if len(col.SetOptions) > 0 { //enum
		nstr += "("
		opts := ""

		setOptions := make([]string, 0, len(col.SetOptions))
		for setOption := range col.SetOptions {
			setOptions = append(setOptions, setOption)
		}
		sort.Strings(setOptions)

		for _, v := range setOptions {
			opts += fmt.Sprintf(",'%v'", v)
		}
		nstr += strings.TrimLeft(opts, ",")
		nstr += ")"
	}
	res = append(res, nstr)

	res = append(res, col.Name)

	var tags []string
	if len(res) > 0 {
		tags = append(tags, "xorm:\""+strings.Join(res, " ")+"\"")
	}
	if len(tags) > 0 {
		return "`" + strings.Join(tags, " ") + "`"
	} else {
		return ""
	}
}

func include(source []string, target string) bool {
	for _, s := range source {
		if s == target {
			return true
		}
	}
	return false
}

// 生成实体类验证标签
func CreateValTag(col *core.Column, typeStr string) string {
	tag := "validate:"

	if col.Nullable {
		tag += `"omitempty`
	} else {
		tag += `"required`
	}

	if col.SQLType.Name == core.Enum {
		tag += ",oneof="
		for option, _ := range col.EnumOptions {
			tag += option + " "
		}
		tag = strings.TrimSpace(tag) + `"`
	} else if strings.EqualFold(typeStr, "string") && col.Length > 0 {
		tag += `,gt=0,lte=` + strconv.Itoa(col.Length) + `"`
	} else {
		tag += `"`
	}
	tag += ` comment:"` + col.Comment + `"`
	return tag
}

// 因为修改的时候字段均不是必填项
func ChangeValTagForUpdate(tag string) string {
	return strings.ReplaceAll(tag, "required", "omitempty")
}
