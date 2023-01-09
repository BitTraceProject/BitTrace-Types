package structure

import (
	"time"
)

type (
	// Revision 代表一个区块同步过程中的某一个阶段结束，
	// 同一个 Snapshot 期间，每一个过程结束时输出一次
	Revision struct {
		SnapshotID string       `json:"snapshot_id"`
		Type       RevisionType `json:"type"` // 代表区块同步过程的一个阶段

		InitTimestamp  Timestamp    `json:"timestamp"`        // Revision 开始时间戳
		InitData       RevisionData `json:"init_data"`        // 根据 Tag 不同获取不同的数据，在 Revision 开始和时输出
		FinalTimestamp Timestamp    `json:"commit_timestamp"` // Revision 生效时间，结合 Timestamp 和 CommitTimestamp 确定 Revision 持续时间
		FinalData      RevisionData `json:"final_data"`       // 根据 Tag 不同获取不同的数据，在 Revision 结束时输出
	}
	RevisionType int
	RevisionData interface{}

	RevisionDataBlockReceive struct {
	}
	RevisionDataBlockVerify struct {
	}
	RevisionDataChainVerify struct {
	}
	RevisionDataOrphanProcess struct {
	}
	RevisionDataOrphanExtend struct {
	}
	RevisionDataMainChainExtend struct {
	}
	RevisionDataSideChainExtend struct {
	}
	RevisionDataChainSwap struct {
	}
	// ......
)

// 每一种 Type 都对应一种 RevisionData
const (
	RevisionTypeBlockReceive RevisionType = iota
	RevisionTypeBlockVerify
	RevisionTypeChainVerify
	RevisionTypeOrphanProcess
	RevisionTypeOrphanExtend
	RevisionTypeMainChainExtend
	RevisionTypeSideChainExtend
	RevisionTypeChainSwap

	RevisionTypeUnknown
	// ......
)

func NewRevision(t RevisionType, snapshotID string, data RevisionData) *Revision {
	r := &Revision{
		Type:          t,
		InitTimestamp: FromNow(),
		InitData:      data,
		SnapshotID:    snapshotID,
	}
	return r
}

// Commit 当前 Revision 生效，需要切换到下一个 Revision
func (r *Revision) Commit(finalTime time.Time, data RevisionData) {
	r.FinalTimestamp = FromTime(finalTime)
	r.FinalData = data
}
