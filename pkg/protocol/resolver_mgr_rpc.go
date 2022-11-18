package protocol

type (
	// NewResolverArgs 根据提供的 ExporterTag 为其分配一个 Resolver
	NewResolverArgs struct {
		ExporterTag string `json:"exporter_tag"`
	}
	// NewResolverReply 返回分配的 Resolver 的 Tag
	NewResolverReply struct {
		ResolverTag string `json:"resolver_tag"`
	}
)
