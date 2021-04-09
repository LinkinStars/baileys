package creator

import (
	"fmt"
	"strconv"
	"strings"

	"xorm.io/core"
)

// MyStructInfoCreatorForGORM 创建者实现，生成 GORM, validatorV9用的相对应的标签
type MyStructInfoCreatorForGORM struct {
	Column  *core.Column
	Table   *core.Table
	typeStr string
}

// CreateTypeString 生成字段的数据类型
func (d *MyStructInfoCreatorForGORM) CreateTypeString() string {
	st := d.Column.SQLType
	t := core.SQLType2Type(st)
	s := t.String()
	if s == "[]uint8" {
		return "[]byte"
	}
	d.typeStr = s
	colName := strings.ToLower(d.Column.Name)
	if colName == "deleted" || colName == "deleted_at" {
		return "gorm.DeletedAt"
	}
	return s
}

// CreateORMTag 生成字段的orm框架标签
func (d *MyStructInfoCreatorForGORM) CreateORMTag() string {
	col := d.Column
	sqlTypeStr := colSQLType(col)

	res := ""
	colName := strings.ToLower(col.Name)
	if colName == "created" || colName == "created_at" {
		res += ";autoCreateTime"
	}
	if colName == "updated" || colName == "updated_at" {
		res += ";autoUpdateTime"
	}

	if col.IsPrimaryKey {
		res += ";primary_key"
	}
	if col.IsAutoIncrement {
		res += ";AUTO_INCREMENT"
	}

	return fmt.Sprintf("`"+`gorm:"column:%s;type:%s%s"`+"`", col.Name, sqlTypeStr, res)
}

// CreateValTag 生成字段的val框架标签
func (d *MyStructInfoCreatorForGORM) CreateValTag() string {
	tag := "validate:"

	if d.Column.Nullable {
		tag += `"omitempty`
	} else {
		tag += `"required`
	}

	if d.Column.SQLType.Name == core.Enum {
		tag += ",oneof="
		for option := range d.Column.EnumOptions {
			tag += option + " "
		}
		tag = strings.TrimSpace(tag) + `"`
	} else if strings.EqualFold(d.typeStr, "string") && d.Column.Length > 0 {
		tag += `,gt=0,lte=` + strconv.Itoa(d.Column.Length) + `"`
	} else {
		tag += `"`
	}
	tag += ` comment:"` + d.Column.Comment + `"`
	return tag
}
