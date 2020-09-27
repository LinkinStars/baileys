package conf

import (
	"github.com/spf13/viper"
)

// All 全部配置索引
var All *AllConfig
var configVip = viper.New()

// InitConfig 初始化读取配置文件
func InitConfig(path string) (err error) {
	configVip.SetConfigFile(path)

	// 读取配置
	if err := configVip.ReadInConfig(); err != nil {
		return err
	}

	// 配置映射到结构体
	All = &AllConfig{}

	return configVip.Unmarshal(All)
}

// ORM 框架的名称枚举
type ORM string

var (
	// XORMName xrom名称
	XORMName ORM = "xorm"
	// GORMName xrom名称
	GORMName ORM = "gorm"
)

// AllConfig 全部配置文件
type AllConfig struct {
	Connection string `mapstructure:"connection"`

	TableNameSuffix    string `mapstructure:"table_name_suffix"`
	TableNamePrefix    string `mapstructure:"table_name_prefix"`
	TableCommentSuffix string `mapstructure:"table_comment_suffix"`

	EveryTplGenPath map[string]string `mapstructure:"every_tpl_gen_path"`
	OneTplGenPath   map[string]string `mapstructure:"one_tpl_gen_path"`

	GenFileSuffix    bool `mapstructure:"gen_file_suffix"`
	AutoImport       bool `mapstructure:"auto_import"`
	IsLowerCamelName bool `mapstructure:"table_field_name_is_lower_camel_name"`

	ORMName ORM `mapstructure:"orm_name"`

	SpecialMapping []string `mapstructure:"mapping"`
}
