package structure

type (
	// Structure 适用于 Event，Result 和 Revision 的接口定义
	Structure interface {
		Tag() Tag             // 标识一种类型
		Context() string      // 关联的各种自定义上下文信息，根据需求设置，根据 Tag 确定解析方式
		Timestamp() Timestamp // 实时的时间戳
	}
)
