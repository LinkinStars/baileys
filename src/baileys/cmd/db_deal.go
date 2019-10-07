package main

import (
	"strings"

	"github.com/go-xorm/xorm"
	"xorm.io/core"

	"baileys/entity"
	"baileys/util"
)

var (
	TableNameSuffix    = "_tab"
	TableCommentSuffix = "表"
)

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

// 将原数据转换成模型
func ConvertRawData2Model(tables []*core.Table) (tableData []*entity.TableData) {
	tableData = make([]*entity.TableData, 0, len(tables))

	for _, table := range tables {
		td := &entity.TableData{}
		// 去掉表名后缀
		newTableName := strings.TrimSuffix(table.Name, TableNameSuffix)
		td.UpperCamelName = util.SqlStr2GoStr(newTableName)
		td.LowerCamelName = util.UpperToLowerCamel(td.UpperCamelName)
		td.UnderlineName = newTableName
		// 去掉表注释的后缀
		td.Comment = strings.TrimSuffix(table.Comment, TableCommentSuffix)
		td.Fields = make([]entity.FieldData, 0)

		columns := table.Columns()
		for _, column := range columns {
			fd := entity.FieldData{}

			fd.UpperCamelName = util.SqlStr2GoStr(column.Name)
			fd.LowerCamelName = util.UpperToLowerCamel(fd.UpperCamelName)
			fd.UnderlineName = column.Name

			fd.Comment = column.Comment
			fd.Type = util.CreateTypeString(column)
			fd.ORMTag = util.CreateXORMTag(table, column)
			fd.ValTag = util.CreateValTag(column, fd.Type)

			td.Fields = append(td.Fields, fd)
		}

		tableData = append(tableData, td)
	}

	return tableData
}