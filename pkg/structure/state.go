package structure

import (
	"github.com/BitTraceProject/BitTrace-Types/pkg/common"
)

type (
	// BestState from btcd chain.BestState
	BestState struct {
		Hash            string           `json:"hash"`             // The hash of the block.
		Height          int32            `json:"height"`           // The height of the block.
		Bits            uint32           `json:"bits"`             // The difficulty bits of the block.
		BlockSize       uint64           `json:"block_size"`       // The size of the block.
		BlockWeight     uint64           `json:"block_weight"`     // The weight of the block.
		NumTxns         uint64           `json:"num_txns"`         // The number of txns in the block.
		TotalTxns       uint64           `json:"total_txns"`       // The total number of txns in the chain.
		MedianTimestamp common.Timestamp `json:"median_timestamp"` // Median time as per CalcPastMedianTime.
	}
)
