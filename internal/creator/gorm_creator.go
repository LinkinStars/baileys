package creator

import (
	"fmt"
	"strings"

	"xorm.io/core"
)

// GormStructInfoCreator 生成 gorm 相关标签和结构体
type GormStructInfoCreator struct {
	Column  *core.Column
	Table   *core.Table
	typeStr string
}

// CreateTypeString 生成字段的数据类型
func (d *GormStructInfoCreator) CreateTypeString() string {
	colName := strings.ToLower(d.Column.Name)
	if colName == "deleted" || colName == "deleted_at" {
		return "gorm.DeletedAt"
	}
	st := d.Column.SQLType
	t := core.SQLType2Type(st)
	s := t.String()
	if s == "[]uint8" {
		return "[]byte"
	}
	if s == "int" {
		return "int64"
	}
	d.typeStr = s
	return s
}

// CreateORMTag 生成字段的orm框架标签
func (d *GormStructInfoCreator) CreateORMTag() string {
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
func (d *GormStructInfoCreator) CreateValTag() string {
	return CreateValidatorTag(d.Column, d.typeStr)
}
