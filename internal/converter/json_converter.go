package converter

import (
	"encoding/json"
	"strings"

	"github.com/LinkinStars/baileys/internal/parsing"
)

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

// GoStruct2Json convert golang struct to json
func GoStruct2Json(structList []*parsing.StructFlat) (jsonStr string) {
	// 所有结构体 json 序列化的数据
	allJsonMap := make(map[string]map[string]interface{})
	// 用于标识是否为根节点结构体
	allStructRootFlag := make(map[string]bool)

	// 首先遍历所有结构体，将所有结构体序列化为 json 数据并保存
	for _, s := range structList {
		jsonMap := make(map[string]interface{})
		for _, field := range s.Fields {
			tag := field.GetJsonTag()
			jsonMap[tag] = GoType2JsonDefaultValue(field.Type)
		}
		allJsonMap[s.Name] = jsonMap
		allStructRootFlag[s.Name] = true
	}

	// 再次遍历每个结构体，获取其中嵌套的结构
	for _, s := range structList {
		for _, field := range s.Fields {
			// 如果当前的类型是别的一个结构体的名称，证明当前结构体嵌套了另一个结构体
			if t, ok := allJsonMap[field.Type]; ok {
				tag := field.GetJsonTag()
				allJsonMap[s.Name][tag] = t
				allStructRootFlag[field.Type] = false
			}
		}
	}

	// 找到根节点结构体
	rootStructName := ""
	for root, ok := range allStructRootFlag {
		if ok {
			rootStructName = root
			break
		}
	}

	// 最终输出根节点结构体所对应的 json
	jsonBytes, _ := json.Marshal(allJsonMap[rootStructName])
	jsonStr = string(jsonBytes)
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
