package creator

import (
	"fmt"
	"sort"
	"strings"

	"xorm.io/core"
)

var created, updated, deleted = []string{"created_at"}, []string{"updated_at"}, []string{"deleted_at"}

// XormStructInfoCreator 生成 xorm 相关标签和结构体
type XormStructInfoCreator struct {
	Column  *core.Column
	Table   *core.Table
	typeStr string
}

// CreateTypeString 生成字段的数据类型
func (d *XormStructInfoCreator) CreateTypeString() string {
	st := d.Column.SQLType
	t := core.SQLType2Type(st)
	s := t.String()
	if s == "[]uint8" {
		return "[]byte"
	}
	d.typeStr = s
	return s
}

// CreateORMTag 生成字段的orm框架标签
func (d *XormStructInfoCreator) CreateORMTag() string {
	col := d.Column
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
		index := d.Table.Indexes[name]
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

	sqlTypeStr := colSQLType(col)
	res = append(res, sqlTypeStr)
	res = append(res, col.Name)

	var tags []string
	if len(res) > 0 {
		tags = append(tags, "xorm:\""+strings.Join(res, " ")+"\"")
	}
	if len(tags) > 0 {
		return "`" + strings.Join(tags, " ") + "`"
	}
	return ""
}

func colSQLType(col *core.Column) string {
	nstr := col.SQLType.Name
	if col.Length != 0 {
		if col.Length2 != 0 {
			nstr += fmt.Sprintf("(%d,%d)", col.Length, col.Length2)
		} else {
			nstr += fmt.Sprintf("(%d)", col.Length)
		}
	} else if len(col.EnumOptions) > 0 {
		nstr += "("
		opts := ""

		enumOptions := make([]string, 0, len(col.EnumOptions))
		for enumOption := range col.EnumOptions {
			enumOptions = append(enumOptions, enumOption)
		}
		sort.Strings(enumOptions)

		for _, v := range enumOptions {
			opts += fmt.Sprintf(",'%s'", v)
		}
		nstr += strings.TrimLeft(opts, ",")
		nstr += ")"
	} else if len(col.SetOptions) > 0 {
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
	return nstr
}

func include(source []string, target string) bool {
	for _, s := range source {
		if s == target {
			return true
		}
	}
	return false
}

// CreateValTag 生成字段的val框架标签
func (d *XormStructInfoCreator) CreateValTag() string {
	return CreateValidatorTag(d.Column, d.typeStr)
}
