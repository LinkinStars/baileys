package deal

import (
	"strings"

	"github.com/go-xorm/xorm"
	"xorm.io/core"

	"github.com/LinkinStars/baileys/internal/conf"
	"github.com/LinkinStars/baileys/internal/deal/creator"
	"github.com/LinkinStars/baileys/internal/entity"
	"github.com/LinkinStars/baileys/internal/util"
)

var (
	// TableNameSuffix 表名后缀
	TableNameSuffix = ""
	// TableNamePrefix 表名前缀
	TableNamePrefix = ""
	// TableCommentSuffix 表注释后缀
	TableCommentSuffix = ""
)

// GetRawTablesData 获取源数据
func GetRawTablesData(connection string) (tables []*core.Table, err error) {
	// 修改为配置文件读取
	engine, err := xorm.NewEngine("mysql", connection)
	if err != nil {
		return nil, err
	}
	// 测试数据库连接
	if err = engine.Ping(); err != nil {
		return nil, err
	}
	engine.SetColumnMapper(core.GonicMapper{})
	// 获取数据库中所有表的各种信息
	tables, err = engine.DBMetas()
	if err != nil {
		return nil, err
	}
	return tables, err
}

// ConvertRawData2Model 将原数据转换成模型
func ConvertRawData2Model(tables []*core.Table) (tableData []*entity.TableData) {
	tableData = make([]*entity.TableData, 0, len(tables))
	for _, table := range tables {
		td := createTableData(table.Name, table.Comment)
		td.Fields = make([]entity.FieldData, 0)

		columns := table.Columns()
		for _, column := range columns {
			fd := createFieldData(column, table)
			td.Fields = append(td.Fields, fd)
		}
		tableData = append(tableData, td)
	}
	return tableData
}

// createTableData 创建表实体数据
func createTableData(tableName, tableComment string) *entity.TableData {
	td := &entity.TableData{}
	// 去掉表名后缀
	newTableName := strings.TrimSuffix(tableName, TableNameSuffix)
	// 去掉表名前缀
	newTableName = strings.TrimPrefix(newTableName, TableNamePrefix)
	// 去掉表注释的后缀
	td.Comment = strings.TrimSuffix(tableComment, TableCommentSuffix)

	td.UpperCamelName = util.SQLStr2GoStr(newTableName)
	td.LowerCamelName = util.UpperToLowerCamel(td.UpperCamelName)
	td.UnderlineName = newTableName
	return td
}

func createFieldData(column *core.Column, table *core.Table) entity.FieldData {
	fd := entity.FieldData{}

	// 如果用户的命名是小驼峰就用小驼峰转大驼峰
	if conf.All.IsLowerCamelName {
		fd.UpperCamelName = util.LowerToUpperCamel(column.Name)
		fd.LowerCamelName = column.Name
	} else {
		fd.UpperCamelName = util.SQLStr2GoStr(column.Name)
		fd.LowerCamelName = util.UpperToLowerCamel(fd.UpperCamelName)
	}
	fd.UnderlineName = column.Name
	fd.Comment = column.Comment

	// 根据配置选择不同的orm框架实现
	var infoCreator creator.StructInfoCreator
	if conf.All.ORMName == conf.GORMName {
		infoCreator = &creator.MyStructInfoCreatorForGORM{
			Column: column,
			Table:  table,
		}
	} else {
		infoCreator = &creator.DefaultStructInfoCreator{
			Column: column,
			Table:  table,
		}
	}

	fd.Type = infoCreator.CreateTypeString()
	fd.ORMTag = infoCreator.CreateORMTag()
	fd.ValTag = infoCreator.CreateValTag()

	return fd
}
