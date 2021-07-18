package converter

import (
	"encoding/json"
	"strings"

	"github.com/LinkinStars/baileys/internal/parsing"
)

// GoStruct2Json convert golang struct to json
func GoStruct2Json(structList []*parsing.StructFlat) (jsonStr string) {
	//existMapping := make(map[string]interface{}, len(structList))
	for _, s := range structList {
		jsonMap := make(map[string]interface{})
		for _, field := range s.Fields {
			tag := field.GetTag("json")
			jsonMap[tag] = GoType2JsonDefaultValue(field.Type)
		}
		jsonBytes, _ := json.Marshal(jsonMap)
		jsonStr = string(jsonBytes)
	}
	return jsonStr
}

// GoType2JsonDefaultValue convert golang type to json default value
func GoType2JsonDefaultValue(goType string) (jsonDefaultValue interface{}) {
	jsonDefaultValue, ok := go2jsonDefaultValueMapping[goType]
	if ok {
		return jsonDefaultValue
	}
	// 处理 array 的情况
	if strings.HasPrefix(goType, "[]") {
		return make([]string, 0)
	}
	// 处理 map 的情况
	if strings.HasPrefix(goType, "map[") {
		return map[string]interface{}{}
	}
	return nil
}

var (
	go2jsonDefaultValueMapping = map[string]interface{}{
		"float32":    0.01,
		"float64":    0.01,
		"complex64":  0.01,
		"complex128": 0.01,
		"int":        0,
		"int8":       0,
		"int16":      0,
		"int32":      0,
		"int64":      0,
		"uint":       0,
		"uint8":      0,
		"uint16":     0,
		"uint32":     0,
		"uint64":     0,
		"bool":       true,
		"string":     "",
		"[]byte":     "",
		"uintptr":    nil,
		"interface":  nil,
		"struct":     "",
		"time.Time":  "",
	}
)
