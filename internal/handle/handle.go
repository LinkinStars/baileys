package handle

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"

	"github.com/LinkinStars/baileys/internal/cache"
	"github.com/LinkinStars/baileys/internal/conf"
	"github.com/LinkinStars/baileys/internal/deal"
	"github.com/LinkinStars/baileys/internal/entity"
	"github.com/LinkinStars/baileys/internal/util"
)

var (
	// 模板中需要使用的自定义函数
	tplFunc = template.FuncMap{
		"UnderlineStr2Strikethrough": util.UnderlineStr2Strikethrough,
		"ChangeValTagForUpdate":      util.ChangeValTagForUpdate,
		"ReplaceTime2TimesISOTime":   util.ReplaceTime2TimesISOTime,
		"ToUpper":                    strings.ToUpper,
		"ToLower":                    strings.ToLower,
	}
)

// LoadingIndex 进入主页加载配置文件、模板、数据库数据
func LoadingIndex(context *gin.Context) {
	err := conf.InitConfig(cache.ConfPath)
	if err != nil {
		context.JSON(http.StatusOK, "读取配置文件出现异常："+err.Error())
		return
	}

	// 设置表名和表注释后缀忽略
	deal.TableNameSuffix = conf.All.TableNameSuffix
	deal.TableNamePrefix = conf.All.TableNamePrefix
	deal.TableCommentSuffix = conf.All.TableCommentSuffix

	// 将配置文件中的映射关系表保存
	for _, value := range conf.All.SpecialMapping {
		util.SpecialMapper[value] = true
	}

	// 连接数据库
	tables, err := deal.GetRawTablesData(conf.All.Connection)
	if err != nil {
		context.JSON(http.StatusOK, "读取数据库出现异常："+err.Error())
		return
	}

	// 将原数据转换为数据模型
	cache.TableData = deal.ConvertRawData2Model(tables)

	// 根据模板的目录 生成模板
	cache.EveryTplList, err = deal.ReadDirGetTemplate(cache.EveryTplPath, tplFunc)
	if err != nil {
		context.JSON(http.StatusOK, "读取every文件夹模板异常："+err.Error())
		return
	}
	cache.OneTplList, err = deal.ReadDirGetTemplate(cache.OneTplPath, tplFunc)
	if err != nil {
		context.JSON(http.StatusOK, "读取one文件夹模板异常："+err.Error())
		return
	}

	// 处理模板后缀和路径等问题
	for _, tpl := range cache.EveryTplList {
		if !conf.All.GenFileSuffix {
			tpl.FilenameSuffix = ".go"
		}
		if path, ok := conf.All.EveryTplGenPath[tpl.Filename]; ok {
			tpl.OutputPath = path
		}
	}
	for _, tpl := range cache.OneTplList {
		tpl.FilenameSuffix = tpl.Filename + ".go"
		if path, ok := conf.All.OneTplGenPath[tpl.Filename]; ok {
			tpl.OutputPath = path
		}
	}

	// 封装返回数据
	data := make(map[string]interface{}, 0)
	data["tableData"] = cache.TableData
	data["everyTplList"] = cache.EveryTplList
	data["oneTplList"] = cache.OneTplList
	context.HTML(http.StatusOK, "index.html", data)
}

// GenCode 代码生成
func GenCode(context *gin.Context) {
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
	if err := gen(chooseTableMap, chooseTplMap); err != nil {
		util.SendFailResp(context, "生成失败："+err.Error())
	} else {
		util.SendSuccessResp(context, "生成成功~!")
	}
}
