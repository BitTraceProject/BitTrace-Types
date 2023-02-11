package protocol

type (
	// ResolverStartArgs 根据提供的 ExporterTag 为其分配一个 Resolver
	ResolverStartArgs struct {
		ExporterTag string `json:"exporter_tag"`
	}
	// ResolverStartReply 返回分配的 Resolver 的 Tag
	ResolverStartReply struct {
		ResolverTag string `json:"resolver_tag"`
		OK          bool   `json:"ok"`
	}
	// ResolverShutdownArgs 根据提供的 ExporterTag 中止并删除 Resolver，要记得检查是否 exporter 已经正确关闭，否则可能造成 mq 中的 exporter 消息永远无法被消费
	ResolverShutdownArgs struct {
		ExporterTag  string `json:"exporter_tag"`
		LazyShutdown bool   `json:"lazy_shutdown"` // 等到消费完消息后再关闭，或者直接关闭
	}
	// ResolverShutdownReply 返回 ok
	ResolverShutdownReply struct {
		OK bool `json:"ok"`
	}
)
