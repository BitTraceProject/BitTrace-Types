package protocol

// meta 基于简单 KV 进行元信息管理, 需要提供的接口包括:
// - 注册 exporter 并与 resolver 建立映射关系: exporter tag 对应 resolver tag
// - exporter tag 查询在不在及与 resolver tag 的双向查询
// - 各种组件的地址信息的注册和查询 (相当于 KV): receiver, mq, collector, resolver-mgr

type (
	// 通用的接口

	GetKeyArgs struct {
		Key string `json:"key"`
	}
	GetKeyReply struct {
		OK bool `json:"ok"`
	}

	GetValueArgs struct {
		Key string `json:"key"`
	}
	GetValueReply struct {
		Value string `json:"value"`
		OK    bool   `json:"ok"`
	}

	SetValueArgs struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	SetValueReply struct {
		OK bool `json:"ok"`
	}

	// TODO 便于使用的接口

)
