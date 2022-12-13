package protocol

// meta 基于简单 KV 进行元信息管理, 需要提供的接口包括:
// - 注册 exporter 并与 resolver 建立映射关系: exporter tag 对应 resolver tag
// - exporter tag 查询在不在及与 resolver tag 的双向查询
// - 各种组件的地址信息的注册和查询 (相当于 KV): receiver, mq, collector, resolver-mgr

type (
	// 通用的接口

	// MetaGetKeyArgs 查询 key 是否存在
	MetaGetKeyArgs struct {
		Key string `json:"key"`
	}
	// MetaGetKeyReply 返回是否，不包括 Value
	MetaGetKeyReply struct {
		OK bool `json:"ok"`
	}
	// MetaDelKeyArgs 查询 key 是否存在
	MetaDelKeyArgs struct {
		Key string `json:"key"`
	}
	// MetaDelKeyReply 返回是否，不包括 Value
	MetaDelKeyReply struct {
		OK bool `json:"ok"`
	}
	// MetaGetValueArgs 返回 Value
	MetaGetValueArgs struct {
		Key string `json:"key"`
	}
	// MetaGetValueReply 返回 Value 和是否存在
	MetaGetValueReply struct {
		Value string `json:"value"`
		OK    bool   `json:"ok"`
	}
	// MetaSetValueArgs 设置 Value
	MetaSetValueArgs struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	// MetaSetValueReply 返回 Value 设置情况，true 代表添加成功，false 更新成功，即 Key 原本已存在
	MetaSetValueReply struct {
		OK bool `json:"ok"`
	}
	// MetaClearArgs 清空 meta 信息
	MetaClearArgs struct {
	}
	// MetaClearReply 返回清空结果
	MetaClearReply struct {
		OK bool `json:"ok"`
	}
)
