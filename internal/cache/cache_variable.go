package cache

import "github.com/LinkinStars/baileys/internal/entity"

const (
	// EveryTplPath Every模板默认位置
	EveryTplPath = "./tpl/every/"
	// OneTplPath One模板默认位置
	OneTplPath = "./tpl/one/"
	// GenRootPath 最终生成的默认位置前缀
	GenRootPath = "./gen/"
)

var (
	// EveryTplList Every模板
	EveryTplList []*entity.TplModel
	// OneTplList One模板
	OneTplList []*entity.TplModel

	// TableData 数据库的表格数据
	TableData []*entity.TableData

	// ConfPath 默认配置文件路径
	ConfPath = "./conf/conf.yml"
	// WebPort 默认端口号
	WebPort = "5272"
	// OpenBrowser 是否打开浏览器
	OpenBrowser = false
)
