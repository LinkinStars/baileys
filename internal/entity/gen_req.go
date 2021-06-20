package entity

// GenReq 生成时请求，记录选中模板和表格
type GenReq struct {
	GenTplNameList   []string `json:"gen_tpl_name_list"`
	GenTableNameList []string `json:"gen_table_name_list"`
}

// ConvertGoStruct2PbMessageReq 将 golang 结构体转换为 Protocol Buffers 请求
type ConvertGoStruct2PbMessageReq struct {
	GoStruct string `json:"go_struct"`
}
