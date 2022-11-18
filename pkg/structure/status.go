package structure

import "time"

type (
	// Status 只需传值
	Status struct {
		MainChainWorldStatus    *WorldStatus            `json:"main_chain_world_status"`
		SideChainWorldStatusMap map[string]*WorldStatus `json:"side_chain_world_status_map"`
	}
	// StatusTransfer 传指针
	// TODO 解决 event 与 result 对应问题
	StatusTransfer struct {
		ChainID        string          `json:"chain_id"`
		FieldName      string          `json:"field_name"` // 也可以通过 reflect 获取 field
		OP             TransferOperate `json:"op"`
		OPDetail       string          `json:"op_detail"`
		RelevantEvent  Event           `json:"relevant_event"`
		RelevantResult Result          `json:"relevant_result"`
	}
	// WorldStatus 传指针
	WorldStatus struct {
		ChainID        string    `json:"chain_id"`
		ChainHeight    int64     `json:"chain_height"`
		Bits           int64     `json:"bits"`
		TotalTxn       int64     `json:"total_txn"`
		NextMedianTime time.Time `json:"next_median_time"`
	}
	TransferOperate int
)

const (
	Reset TransferOperate = iota // reset one field
	Swap                         // swap all field
	None  = -1                   // no operate
)

func NewStatus(mainChainWorldStatus *WorldStatus) Status {
	return Status{MainChainWorldStatus: mainChainWorldStatus, SideChainWorldStatusMap: map[string]*WorldStatus{}}
}

func (s Status) IsMainChain(chainID string) bool {
	return chainID == s.MainChainWorldStatus.ChainID
}

func (s Status) IsSideChain(chainID string) bool {
	_, ok := s.SideChainWorldStatusMap[chainID]
	return ok
}

func (s Status) AddSideChain(sideChainWorldStatus *WorldStatus) bool {
	chainID := sideChainWorldStatus.ChainID
	if s.IsMainChain(chainID) {
		// 待添加的 side chain 在 status 处记得是 main chain，则什么都不做，返回 false
		return false
	}
	// 直接添加或者替换
	s.SideChainWorldStatusMap[chainID] = sideChainWorldStatus
	return true
}

func (s Status) RemoveSideChain(chainID string) bool {
	if s.IsMainChain(chainID) {
		// 待删除的 side chain 在 status 处记得是 main chain，则什么都不做，返回 false
		return false
	}
	// 直接删除
	delete(s.SideChainWorldStatusMap, chainID)
	return true
}

func (s Status) ResetMainChain(newMainChainWorldStatus *WorldStatus) {
	chainID := newMainChainWorldStatus.ChainID
	if s.IsSideChain(chainID) {
		// 如果待交换的 chain 是 side chain 则先删除
		s.RemoveSideChain(chainID)
	}
	s.MainChainWorldStatus = newMainChainWorldStatus
}

func (s Status) SwapMainChain(newMainChainWorldStatus *WorldStatus, removeOldMainChain bool) bool {
	chainID := newMainChainWorldStatus.ChainID
	if s.IsMainChain(chainID) {
		// 如果待交换的 chain 已经是 main chain 则什么都不做，返回 false
		return false
	}
	if s.IsSideChain(chainID) {
		// 如果待交换的 chain 是 side chain 则先删除
		s.RemoveSideChain(chainID)
	}
	oldMainChainWorldStatus := s.MainChainWorldStatus
	// 替换
	s.MainChainWorldStatus = newMainChainWorldStatus
	if !removeOldMainChain {
		// 如果旧的 main chain 不删除，则放入 side chain
		s.SideChainWorldStatusMap[oldMainChainWorldStatus.ChainID] = oldMainChainWorldStatus
	}
	return true
}

func (s Status) Transfer(trans *StatusTransfer) bool {
	chainID := trans.ChainID
	if chainID == s.MainChainWorldStatus.ChainID {
		s.MainChainWorldStatus.Transfer(trans)
		return true
	}
	if sideChainWS, ok := s.SideChainWorldStatusMap[chainID]; ok {
		sideChainWS.Transfer(trans)
	}
	return false
}

func NewStatusTransfer(event Event, result Result, chainID string, fieldName string, op TransferOperate, opDetail string) *StatusTransfer {
	return &StatusTransfer{
		ChainID:        chainID,
		FieldName:      fieldName,
		OP:             op,
		OPDetail:       opDetail,
		RelevantEvent:  event,
		RelevantResult: result,
	}
}

func NewWorldStatus(forkHeight int64, bits int64, totalTxn int64, nextMedianTime time.Time) WorldStatus {
	chainID := GenChainID(forkHeight)
	s := WorldStatus{
		ChainID:        chainID,
		ChainHeight:    forkHeight,
		Bits:           0,
		TotalTxn:       0,
		NextMedianTime: nextMedianTime,
	}
	return s
}

func (ws *WorldStatus) Transfer(trans *StatusTransfer) {
	// TODO 根据 fieldName, op 和 opDetail 修改状态

}
