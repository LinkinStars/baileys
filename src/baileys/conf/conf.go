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
	if err := configVip.Unmarshal(All); err != nil {
		return err
	}

	return nil
}

// AllConfig 全部配置文件
type AllConfig struct {
	HttpPort   string `mapstructure:"http_port"`
	Connection string `mapstructure:"connection"`

	TableNameSuffix    string `mapstructure:"table_name_suffix"`
	TableCommentSuffix string `mapstructure:"table_comment_suffix"`

	EveryTplGenPath map[string]string `mapstructure:"every_tpl_gen_path"`
	OneTplGenPath   map[string]string `mapstructure:"one_tpl_gen_path"`

	GenFileSuffix  bool     `mapstructure:"gen_file_suffix"`
	SpecialMapping []string `mapstructure:"mapping"`
}
