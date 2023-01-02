package structure

import "time"

type (
	// Snapshot 代表一次完整的区块同步流程，主要由一系列的 Revision 组成，
	// 输出两次，初始化一次，结束一次，之间所有 ID 一致的 Revision 都属于这个 Snapshot
	Snapshot struct {
		ID                string       `json:"id"` // init timestamp + chain id + chain height
		TargetChainID     string       `json:"target_chain_id"`
		TargetChainHeight int32        `json:"target_chain_height"`
		Type              SnapshotType `json:"type"`
		Timestamp         Timestamp    `json:"timestamp"`
		Status            *Status      `json:"status"`
		RevisionList      []Revision   `json:"revision_list"` // 输出的时候不会输出这个，也不会维护它（在结构外维护一个 List），这个是逻辑上的关系（由 resolver 处理得到）
	}
	SnapshotType int
)

const (
	SnapshotInit = iota
	SnapshotFinal
)

func InitSnapshot(targetChainID string, targetChainHeight int32, t time.Time, initStatus *Status) Snapshot {
	timestamp := FromTime(t)
	id := GenSnapshotID(targetChainID, targetChainHeight, timestamp)
	initSnapshot := Snapshot{
		ID:                id,
		TargetChainID:     targetChainID,
		TargetChainHeight: targetChainHeight,
		Type:              SnapshotInit,
		Timestamp:         timestamp,
		Status:            initStatus,
		RevisionList:      []Revision{},
	}
	return initSnapshot
}

func FinalSnapshot(initSnapshot Snapshot, t time.Time, finalStatus *Status) Snapshot {
	timestamp := FromTime(t)
	finalSnapshot := Snapshot{
		ID:                initSnapshot.ID,
		TargetChainID:     initSnapshot.TargetChainID,
		TargetChainHeight: initSnapshot.TargetChainHeight,
		Type:              SnapshotFinal,
		Timestamp:         timestamp,
		Status:            finalStatus,
		RevisionList:      initSnapshot.RevisionList,
	}
	return finalSnapshot
}
