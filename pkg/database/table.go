package database

type (
	TableSnapshotData struct {
		SnapshotID        string `json:"snapshot_id"`
		TargetChainID     string `json:"target_chain_id"`
		TargetChainHeight int32  `json:"target_chain_height"`
		InitTimestamp     string `json:"init_timestamp"`
		FinalTimestamp    string `json:"final_timestamp"`
	}
	TableSnapshotSync struct {
		SnapshotID        string `json:"snapshot_id"`
		TargetChainID     string `json:"target_chain_id"`
		TargetChainHeight int32  `json:"target_chain_height"`
		SyncTimestamp     string `json:"sync_timestamp"`
	}
	TableState struct {
		SnapshotID      string `json:"snapshot_id"`
		SnapshotType    int    `json:"snapshot_type"`
		HashStr         string `json:"hash_str"`
		Height          int32  `json:"height"`
		Bits            uint32 `json:"bits"`
		BlockSize       uint64 `json:"block_size"`
		BlockWeight     uint64 `json:"block_weight"`
		NumTxns         uint64 `json:"num_txns"`
		TotalTxns       uint64 `json:"total_txns"`
		MedianTimestamp string `json:"median_timestamp"`
	}
	TableRevision struct {
		SnapshotID     string `json:"snapshot_id"`
		RevisionType   int    `json:"revision_type"`
		InitTimestamp  string `json:"init_timestamp"`
		FinalTimestamp string `json:"final_timestamp"`
		InitData       string `json:"init_data"`  // json str
		FinalData      string `json:"final_data"` // json str
	}
)
