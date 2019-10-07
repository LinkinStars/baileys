package entity

// GenReq 生成时请求，记录选中模板和表格
type GenReq struct {
	GenTplNameList   []string `json:"gen_tpl_name_list"`
	GenTableNameList []string `json:"gen_table_name_list"`
}
