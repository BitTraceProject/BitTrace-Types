package structure

import (
	"time"

	"github.com/BitTraceProject/BitTrace-Types/pkg/common"
)

type (
	// Revision 代表一个区块同步过程中的某一个阶段结束，
	// 同一个 Snapshot 期间，每一个过程结束时输出一次
	Revision struct {
		SnapshotID string       `json:"snapshot_id"`
		Type       RevisionType `json:"revision_type"` // 代表区块同步过程的一个阶段

		InitTimestamp common.Timestamp `json:"init_timestamp"` // Revision 开始时间戳
		InitData      RevisionData     `json:"init_data"`      // 根据 Tag 不同获取不同的数据，在 Revision 开始和时输出

		FinalTimestamp common.Timestamp `json:"final_timestamp"` // Revision 生效时间，结合 Timestamp 和 CommitTimestamp 确定 Revision 持续时间
		FinalData      RevisionData     `json:"final_data"`      // 根据 Tag 不同获取不同的数据，在 Revision 结束时输出
	}
	RevisionType int

	RevisionData interface{}

	RevisionDataBlockReceiveInit struct {
		PeerIPAddr      string `json:"peer_ip_addr"` // ipv4 or ipv6
		MinerWalletAddr string `json:"miner_wallet_addr"`
	}
	RevisionDataBlockReceiveFinal struct {
		OK bool `json:"ok"`
	}
	RevisionDataBlockVerifyInit struct {
		Hash       string `json:"hash"`
		ParentHash string `json:"parent_hash"`
		Height     int32  `json:"height"`
		NumTxns    uint64 `json:"num_txns"`

		Version        int32  `json:"version"`
		Bits           uint32 `json:"bits"`
		MerkleRootHash string `json:"merkle_root_hash"`
		Nonce          uint32 `json:"nonce"`
	}
	RevisionDataBlockVerifyFinal struct {
		OK bool `json:"ok"`
	}
	RevisionDataOrphanProcessInit struct {
	}
	RevisionDataOrphanProcessFinal struct {
	}
	RevisionDataOrphanExtendInit struct {
	}
	RevisionDataOrphanExtendFinal struct {
	}
	RevisionDataMainChainExtendInit struct {
	}
	RevisionDataMainChainExtendFinal struct {
	}
	RevisionDataSideChainExtendInit struct {
		ForkParentBlockHash string `json:"fork_parent_block_hash"`
	}
	RevisionDataSideChainExtendFinal struct {
		ForkBlockHash string `json:"fork_block_hash"`

		OK bool `json:"ok"`
	}
	RevisionDataChainSwapInit struct {
		OldBestBlockHash string `json:"old_best_block_hash"`
		NewBestBlockHash string `json:"new_best_block_hash"`
	}
	RevisionDataChainSwapFinal struct {
		ForkParentBlockHash string `json:"fork_parent_block_hash"`
		ForkBlockHash       string `json:"fork_block_hash"`

		OK bool `json:"ok"`
	}
	RevisionDataChainVerifyInit struct {
	}
	RevisionDataChainVerifyFinal struct {
	}
	// ......
)

// 每一种 Type 都对应一种 RevisionData
const (
	RevisionTypeBlockReceive RevisionType = iota
	RevisionTypeBlockVerify
	RevisionTypeOrphanProcess
	RevisionTypeOrphanExtend
	RevisionTypeMainChainExtend
	RevisionTypeSideChainExtend
	RevisionTypeChainSwap
	RevisionTypeChainVerify

	RevisionTypeUnknown
	// ......
)

func NewRevision(t RevisionType, snapshotID string, data RevisionData) *Revision {
	r := &Revision{
		Type:          t,
		InitTimestamp: common.FromNow(),
		InitData:      data,
		SnapshotID:    snapshotID,
	}
	return r
}

// Commit 当前 Revision 生效，需要切换到下一个 Revision
func (r *Revision) Commit(finalTime time.Time, data RevisionData) {
	r.FinalTimestamp = common.FromTime(finalTime)
	r.FinalData = data
}
