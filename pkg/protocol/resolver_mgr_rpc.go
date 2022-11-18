package protocol

type (
	NewResolverArgs struct {
		ExporterTag string `json:"exporter_tag"`
	}
	NewResolverReply struct {
		ResolverTag string `json:"resolver_tag"`
	}
)
