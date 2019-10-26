package handle

import (
	"errors"
	"fmt"

	"baileys/cache"
	"baileys/conf"
	"baileys/deal"
	"baileys/entity"
)

// gen 开始生成
func gen(chooseTableMap, chooseTplMap map[string]bool) (err error) {
	// 循环一遍数据表，确定用户选中的表
	userChooseTable := make([]*entity.TableData, 0)
	for _, v := range cache.TableData {
		if chooseTableMap[v.UpperCamelName] {
			userChooseTable = append(userChooseTable, v)
		}
	}
	if len(userChooseTable) == 0 {
		return errors.New("未选中任何表")
	}

	// 生成every模板对应的代码
	err = genEveryTplCode(chooseTplMap, userChooseTable)
	if err != nil {
		return
	}

	// 生成one模板对应的代码
	err = genOneTplCode(chooseTplMap, userChooseTable)
	if err != nil {
		return
	}

	return
}

// genEveryTplCode 生成every模板对应的代码
func genEveryTplCode(chooseTplMap map[string]bool, userChooseTable []*entity.TableData) (err error) {
	// 只操作用户选中的模板
	tempEveryTplList := make([]*entity.TplModel, 0)
	for i := 0; i < len(cache.EveryTplList); i++ {
		if chooseTplMap[cache.EveryTplList[i].Tpl.Name()] {
			tempEveryTplList = append(tempEveryTplList, cache.EveryTplList[i])
		}
	}
	// 创建对应文件生成的目录
	err = deal.CreateGenPathList(tempEveryTplList)
	if err != nil {
		return errors.New("创建文件夹异常：" + err.Error())
	}

	// 生成every文件夹下模板对应的文件
	for _, v := range userChooseTable {
		fmt.Println("生成：", v.UpperCamelName, v.Comment)
		for _, tpl := range tempEveryTplList {
			err := deal.GenFile(tpl.Tpl, tpl.OutputPath, v.UnderlineName+tpl.FilenameSuffix, v)
			if err != nil {
				return errors.New("生成 模板：" + tpl.Filename + "表：" + v.UnderlineName + " 操作时异常：" + err.Error())
			}
		}

		// 如果不需要导包那就继续
		if !conf.All.AutoImport {
			continue
		}

		// 全部生成之后再重新导入包这样才能导入完完整
		for _, tpl := range tempEveryTplList {
			err = deal.FormatAndImport(tpl.OutputPath + v.UnderlineName + tpl.FilenameSuffix)
			if err != nil {
				return errors.New("导包 " + tpl.Filename + " 操作时异常：" + err.Error())
			}
		}
	}
	return
}

// genEveryTplCode 生成every模板对应的代码
func genOneTplCode(chooseTplMap map[string]bool, userChooseTable []*entity.TableData) (err error) {
	// 只操作用户选中的模板
	tempOneTplList := make([]*entity.TplModel, 0)
	for i := 0; i < len(cache.OneTplList); i++ {
		if chooseTplMap[cache.OneTplList[i].Tpl.Name()] {
			tempOneTplList = append(tempOneTplList, cache.OneTplList[i])
		}
	}

	// 创建对应文件生成的目录
	err = deal.CreateGenPathList(tempOneTplList)
	if err != nil {
		return errors.New("创建文件夹异常：" + err.Error())
	}

	// 生成one文件夹下模板对应的文件
	for _, tpl := range tempOneTplList {
		if err := deal.GenFile(tpl.Tpl, tpl.OutputPath, tpl.FilenameSuffix, userChooseTable); err != nil {
			return errors.New("生成 " + tpl.Filename + " 操作时异常：" + err.Error())
		}

		// 如果不需要导包那就继续
		if !conf.All.AutoImport {
			continue
		}

		if err := deal.FormatAndImport(tpl.OutputPath + tpl.FilenameSuffix); err != nil {
			return errors.New("导包 " + tpl.Filename + " 操作时异常：" + err.Error())
		}
	}
	return
}
