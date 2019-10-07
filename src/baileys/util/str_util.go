package util

import (
	"bytes"
	"strings"
)

// 大驼峰转小驼峰
func UpperToLowerCamel(s string) string {
	if s == "" {
		return s
	}

	str := bytes.Buffer{}

	for i := 0; i < len(s); i++ {
		if r := rune(s[i]); r >= 'A' && r <= 'Z' {
			str.WriteString(strings.ToLower(string(r)))
		} else {
			str.WriteString(s[i:])
			break
		}
	}

	return str.String()
}

// 将下划线字符串转换为中划线
func UnderlineStr2Strikethrough(str string) string {
	return strings.ReplaceAll(str, "_", "-")
}
