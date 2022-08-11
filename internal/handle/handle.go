package handle

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/LinkinStars/baileys/internal/converter"
	"github.com/LinkinStars/baileys/internal/generator"
	"github.com/LinkinStars/baileys/internal/parsing"
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

// ConverterSql2Code 进入主页加载配置文件、模板、数据库数据
func ConverterSql2Code(context *gin.Context) {
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
	data := make(map[string]interface{}, 4)
	data["tableData"] = cache.TableData
	data["everyTplList"] = cache.EveryTplList
	data["oneTplList"] = cache.OneTplList
	context.HTML(http.StatusOK, "sql_2_code.html", data)
}

// ConvertSql2GoCode 将 sql 数据转换为对应模板的 go 代码
func ConvertSql2GoCode(ctx *gin.Context) {
	genReq := &entity.GenReq{}
	if err := ctx.Bind(genReq); err != nil {
		fmt.Println(err)
		return
	}

	// 将用户选中的模板和表格保存起来
	chooseTableMap, chooseTplMap := make(map[string]bool, 4), make(map[string]bool, 4)
	for _, tableName := range genReq.GenTableNameList {
		chooseTableMap[tableName] = true
	}
	for _, tplName := range genReq.GenTplNameList {
		chooseTplMap[tplName] = true
	}

	// 生成
	if err := gen(chooseTableMap, chooseTplMap); err != nil {
		util.SendFailResp(ctx, "生成失败："+err.Error())
	} else {
		util.SendSuccessResp(ctx, "生成成功~!")
	}
}

// ConvertGoStruct2PbMessage 将 golang 结构体转换为 Protocol Buffers
func ConvertGoStruct2PbMessage(ctx *gin.Context) {
	req := &entity.ConvertGoStruct2PbMessageReq{}
	if err := ctx.Bind(req); err != nil {
		_ = ctx.Error(err)
		return
	}

	structList, err := parsing.StructParser(req.GoStruct)
	if err != nil {
		util.SendFailResp(ctx, "生成失败："+err.Error())
		return
	}

	pbList := converter.GoStruct2PB(structList)
	message, err := generator.GenPBMessage(pbList)
	if err != nil {
		util.SendFailResp(ctx, "生成失败："+err.Error())
		return
	}
	convertFunc, err := generator.GenerateStruct2PBFunc(structList)
	if err != nil {
		util.SendFailResp(ctx, "生成失败："+err.Error())
		return
	}

	buf := bytes.NewBufferString(message)
	buf.WriteString("\n")
	buf.WriteString(strings.Repeat("-", 20))
	buf.WriteString("\n\n")
	buf.WriteString(convertFunc)
	util.SendResp(ctx, 200, 200, "生成成功", buf.String())
}

// ConvertGoStruct2Json 将 golang 结构体转换为 json
func ConvertGoStruct2Json(ctx *gin.Context) {
	req := &entity.ConvertGoStruct2JsonReq{}
	if err := ctx.Bind(req); err != nil {
		_ = ctx.Error(err)
		return
	}

	structList, err := parsing.StructParser(req.GoStruct)
	if err != nil {
		util.SendFailResp(ctx, "生成失败："+err.Error())
		return
	}

	jsonStr := converter.GoStruct2Json(structList)
	util.SendResp(ctx, 200, 200, "生成成功", jsonStr)
}
