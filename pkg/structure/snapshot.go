package structure

import "time"

type (
	// Snapshot 代表一次完整的区块同步流程，主要由一系列的 Revision 组成，
	// 输出两次，初始化一次，结束一次，之间所有 ID 一致的 Revision 都属于这个 Snapshot
	Snapshot struct {
		ID                string     `json:"id"` // chain id + chain height + init timestamp
		TargetChainID     string     `json:"target_chain_id"`
		TargetChainHeight int64      `json:"target_chain_height"`
		InitTimestamp     Timestamp  `json:"init_timestamp"`
		FinalTimestamp    Timestamp  `json:"final_timestamp"`
		InitStatus        Status     `json:"init_status"`
		FinalStatus       Status     `json:"final_status"`  // 这个其实不是必要的，可以用于核对 status transfer list 的操作结果是否正确
		RevisionList      []Revision `json:"revision_list"` // 输出的时候不会输出这个，也不会维护它（在结构外维护一个 List），这个是逻辑上的关系（由 resolver 处理得到）
	}
)

func InitSnapshot(targetChainID string, targetChainHeight int64, initTime time.Time, initStatus Status) Snapshot {
	initTimestamp := FromTime(initTime)
	id := GenSnapshotID(targetChainID, targetChainHeight, initTimestamp)
	s := Snapshot{
		ID:                id,
		TargetChainID:     targetChainID,
		TargetChainHeight: targetChainHeight,
		InitTimestamp:     initTimestamp,
		FinalTimestamp:    0,
		InitStatus:        initStatus,
		FinalStatus:       Status{},
		RevisionList:      nil,
	}
	return s
}

func FinalSnapshot(s Snapshot, finalTime time.Time, finalStatus Status) Snapshot {
	finalTimestamp := FromTime(finalTime)
	s.FinalTimestamp = finalTimestamp
	s.FinalStatus = finalStatus
	return s
}
