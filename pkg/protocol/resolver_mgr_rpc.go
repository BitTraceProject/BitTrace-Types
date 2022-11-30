package protocol

type (
	// StartResolverArgs 根据提供的 ExporterTag 为其分配一个 Resolver
	StartResolverArgs struct {
		ExporterTag string `json:"exporter_tag"`
	}
	// StartResolverReply 返回分配的 Resolver 的 Tag
	StartResolverReply struct {
		ResolverTag string `json:"resolver_tag"`
	}
	// ShutdownResolverArgs 根据提供的 ExporterTag 中止并删除 Resolver，要记得检查是否 exporter 已经正确关闭，否则可能造成 mq 中的 exporter 消息永远无法被消费
	ShutdownResolverArgs struct {
		ExporterTag  string `json:"exporter_tag"`
		LazyShutdown bool   `json:"lazy_shutdown"` // 等到消费完消息后再关闭，或者直接关闭
	}
	// ShutdownResolverReply 返回分配的 Resolver 的 Tag
	ShutdownResolverReply struct {
		OK bool `json:"ok"`
	}
)
