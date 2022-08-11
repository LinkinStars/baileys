package creator

import (
	"strconv"
	"strings"

	"xorm.io/core"
)

// CreateValidatorTag 创建 https://github.com/go-playground/validator/ 对应 tag
func CreateValidatorTag(column *core.Column, typeStr string) string {
	tag := "validate:"

	if column.Nullable {
		tag += `"omitempty`
	} else {
		tag += `"required`
	}

	if column.SQLType.Name == core.Enum {
		tag += ",oneof="
		for option := range column.EnumOptions {
			tag += option + " "
		}
		tag = strings.TrimSpace(tag) + `"`
	} else if strings.EqualFold(typeStr, "string") && column.Length > 0 {
		tag += `,gt=0,lte=` + strconv.Itoa(column.Length) + `"`
	} else {
		tag += `"`
	}
	tag += ` comment:"` + column.Comment + `"`
	return tag
}
