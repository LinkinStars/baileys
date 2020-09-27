package creator

// StructInfoCreator 结构体数据创建者，用于实现各种创建方法
type StructInfoCreator interface {
	// 生成字段的数据类型
	CreateTypeString() string
	// 生成字段的orm框架标签
	CreateORMTag() string
	// 生成字段的val框架标签
	CreateValTag() string
}
