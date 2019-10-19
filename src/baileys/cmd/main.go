package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"baileys/conf"
	"baileys/entity"
	"baileys/util"
)

var (
	// 模板中需要使用的自定义函数
	tplFunc = template.FuncMap{
		"UnderlineStr2Strikethrough": util.UnderlineStr2Strikethrough,
		"ChangeValTagForUpdate":      util.ChangeValTagForUpdate,
	}

	// 读取模板文件列表
	everyTplPath = "./tpl/every/"
	oneTplPath   = "./tpl/one/"
	genRootPath  = "./gen/"

	// 模板和数据
	everyTplList []*entity.TplModel
	oneTplList   []*entity.TplModel
	tableData    []*entity.TableData

	// 默认配置文件路径
	confPath = "./conf/conf.yml"
	// 默认端口号
	webPort = "5272"
)

func init() {
	flag.StringVar(&confPath, "c", "./conf/conf.yml", "default config path")
	flag.StringVar(&webPort, "p", "5272", "default web port")
}

func main() {
	flag.Parse()

	if err := conf.InitConfig(confPath); err != nil {
		panic("读取配置文件异常：" + err.Error())
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static")

	router.POST("/gen", func(context *gin.Context) {
		genReq := &entity.GenReq{}
		if err := context.Bind(genReq); err != nil {
			fmt.Println(err)
			return
		}

		// 将用户选中的模板和表格保存起来
		chooseTableMap, chooseTplMap := make(map[string]bool, 0), make(map[string]bool, 0)
		for _, tableName := range genReq.GenTableNameList {
			chooseTableMap[tableName] = true
		}
		for _, tplName := range genReq.GenTplNameList {
			chooseTplMap[tplName] = true
		}

		// 生成
		if err := Gen(chooseTableMap, chooseTplMap); err != nil {
			util.SendFailResp(context, "生成失败："+err.Error())
		} else {
			util.SendSuccessResp(context, "生成成功~!")
		}
	})

	router.GET("", func(context *gin.Context) {
		err := conf.InitConfig("./conf/conf.yml")
		if err != nil {
			context.JSON(http.StatusOK, "读取配置文件出现异常："+err.Error())
			return
		}

		// 设置表名和表注释后缀忽略
		TableNameSuffix = conf.All.TableNameSuffix
		TableCommentSuffix = conf.All.TableCommentSuffix

		// 将配置文件中的映射关系表保存
		for _, value := range conf.All.SpecialMapping {
			util.SpecialMapper[value] = true
		}

		// 连接数据库
		tables, err := GetRawTablesData(conf.All.Connection)
		if err != nil {
			context.JSON(http.StatusOK, "读取数据库出现异常："+err.Error())
			return
		}

		// 将原数据转换为数据模型
		tableData = ConvertRawData2Model(tables)

		// 根据模板的目录 生成模板
		everyTplList, err = ReadDirGetTemplate(everyTplPath, tplFunc)
		if err != nil {
			context.JSON(http.StatusOK, "读取every文件夹模板异常："+err.Error())
			return
		}
		oneTplList, err = ReadDirGetTemplate(oneTplPath, tplFunc)
		if err != nil {
			context.JSON(http.StatusOK, "读取one文件夹模板异常："+err.Error())
			return
		}

		// 处理模板后缀和路径等问题
		for _, tpl := range everyTplList {
			if !conf.All.GenFileSuffix {
				tpl.FilenameSuffix = ".go"
			}
			if path, ok := conf.All.EveryTplGenPath[tpl.Filename]; ok {
				tpl.OutputPath = path
			}
		}

		for _, tpl := range oneTplList {
			tpl.FilenameSuffix = tpl.Filename + ".go"
			if path, ok := conf.All.OneTplGenPath[tpl.Filename]; ok {
				tpl.OutputPath = path
			}
		}

		data := make(map[string]interface{}, 0)
		data["tableData"] = tableData
		data["everyTplList"] = everyTplList
		data["oneTplList"] = oneTplList

		context.HTML(http.StatusOK, "index.html", data)
	})
	err := util.OpenBrowser("http://127.0.0.1:" + webPort)
	if err != nil {
		panic(err)
	}
	err = router.Run(":" + webPort)
	if err != nil {
		panic(err)
	}
}

func Gen(chooseTableMap, chooseTplMap map[string]bool) (err error) {
	// 只操作用户选中的模板
	tempEveryTplList := make([]*entity.TplModel, 0)
	tempOneTplList := make([]*entity.TplModel, 0)
	for i := 0; i < len(everyTplList); i++ {
		if chooseTplMap[everyTplList[i].Tpl.Name()] {
			tempEveryTplList = append(tempEveryTplList, everyTplList[i])
		}
	}
	for i := 0; i < len(oneTplList); i++ {
		if chooseTplMap[oneTplList[i].Tpl.Name()] {
			tempOneTplList = append(tempOneTplList, oneTplList[i])
		}
	}

	// 循环一遍数据表，确定用户选中的表
	userChooseTable := make([]*entity.TableData, 0)
	for _, v := range tableData {
		if chooseTableMap[v.UpperCamelName] {
			userChooseTable = append(userChooseTable, v)
		}
	}

	if len(userChooseTable) == 0 {
		return errors.New("未选中任何表")
	}

	// 创建对应文件生成的目录
	err = CreateGenPathList(tempEveryTplList)
	if err != nil {
		return errors.New("创建文件夹异常：" + err.Error())
	}
	err = CreateGenPathList(tempOneTplList)
	if err != nil {
		return errors.New("创建文件夹异常：" + err.Error())
	}

	// 生成every文件夹下模板对应的文件
	for _, v := range userChooseTable {
		fmt.Println("生成：", v.UpperCamelName, v.Comment)
		for _, tpl := range tempEveryTplList {
			err := GenFile(tpl.Tpl, tpl.OutputPath, v.UnderlineName+tpl.FilenameSuffix, v)
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
			err = FormatAndImport(tpl.OutputPath + v.UnderlineName + tpl.FilenameSuffix)
			if err != nil {
				return errors.New("导包 " + tpl.Filename + " 操作时异常：" + err.Error())
			}
		}
	}

	// 生成one文件夹下模板对应的文件
	for _, tpl := range tempOneTplList {
		err := GenFile(tpl.Tpl, tpl.OutputPath, tpl.FilenameSuffix, userChooseTable)
		if err != nil {
			return errors.New("生成 " + tpl.Filename + " 操作时异常：" + err.Error())
		}

		// 如果不需要导包那就继续
		if !conf.All.AutoImport {
			continue
		}

		err = FormatAndImport(tpl.OutputPath + tpl.FilenameSuffix)
		if err != nil {
			return errors.New("导包 " + tpl.Filename + " 操作时异常：" + err.Error())
		}
	}
	return nil
}
