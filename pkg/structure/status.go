package structure

import (
	"sync"
	"time"
)

type (
	// Status 维护了已存在的 MainChain 和 SideChain 的 WorldStatus，支持一系列管理操作，只需传值
	Status struct {
		sync.RWMutex

		MainChainWorldStatus    *WorldStatus            `json:"main_chain_world_status"`
		SideChainWorldStatusMap map[string]*WorldStatus `json:"side_chain_world_status_map"`
	}
	// StatusTransfer 是对与 WorldStatus 的一次操作，称为`状态迁移`，传指针，
	// 由于 Status 本身会受到并发影响，单独 Snapshot 内的 Revision 可能只对 Status 产生部分影响，
	// Revision 只能记录它本身造成的 StatusTransfer，从局部的 Revision 不一定推出完整的 StatusTransfer，
	// 所以在处理数据时，需要将当前 Revision 时间范围内的所有 OP （可能来自多个 Revision）按照 Result 的时间戳串行化排序，来从 initStatus 推导出 finalStatus
	StatusTransfer struct {
		ChainID string `json:"chain_id"` // 唯一标识当前状态迁移针对的哪一个链

		// 状态相关变更
		FieldName     string      `json:"field_name"`      // 操作的目标 Field，也可以通过 reflect 获取 field
		FieldNewValue interface{} `json:"field_new_value"` // 新的值
		FieldOldValue interface{} `json:"field_old_value"` // 在完成 field 修改前，重新从网络中获取到的旧的值，可用于双向的核对

		// 链相关变更
		ChainNewWorldStatus *WorldStatus `json:"chain_new_world_status"` // 新的 worldstatus
		RemoveOldMainChain  bool         `json:"remove_old_main_chain"`  // 是否移除旧 mainchain

		// 描述性信息
		OP       TransferOperate `json:"op"`        // 操作类型
		OPDetail string          `json:"op_detail"` // 操作的具体内容，根据 OP 还原出操作

		// 关联性信息
		RelevantSnapshotID  string `json:"relevant_snapshot_id"`
		RelevantRevisionTag string `json:"relevant_revision_tag"`
		RelevantEvent       Event  `json:"relevant_event"`  // 状态迁移所关联的事件，事件发生会导致结果
		RelevantResult      Result `json:"relevant_result"` // 事件对应的结果，这个结果才最终会导致状态的迁移
	}
	// WorldStatus 是各种自定义标准化属性的集合，传指针
	WorldStatus struct {
		ChainID        string    `json:"chain_id"`
		ChainHeight    int32     `json:"chain_height"`     // 当前链的高度
		Bits           int64     `json:"bits"`             // 当前链的网络难度
		TotalTxn       int64     `json:"total_txn"`        // 当前链的交易数目
		NextMedianTime time.Time `json:"next_median_time"` // 下一次出块时间估计
	}
	TransferOperate int
)

const (
	ResetField      TransferOperate = iota // 重置单个 field
	SwapField                              // 替换所有 field，适合改变多个 field 的情况（也可使用多次 reset 的方法代替）
	AddSideChain                           // 添加一个 sidechain
	RemoveSideChain                        // 移除一个 sidechain
	ResetMainChain                         // 重置 mainchain，新的 mainchain 可能是新的 chain 也可能是 sidechain，旧 mainchain 直接删除
	SwapMainChain                          // 重置 mainchain，新的 mainchain 可能是新的 chain 也可能是 sidechain，旧 mainchain 变为 sidechain
	NoneOperate     = -1                   // no operate
)

// NewStatus 初始化当前的状态
func NewStatus(mainChainWorldStatus *WorldStatus, sideChainWorldStatusMap map[string]*WorldStatus) *Status {
	if sideChainWorldStatusMap == nil {
		// 如果 sideChainWorldStatusMap 不存在则初始化一个空的
		sideChainWorldStatusMap = map[string]*WorldStatus{}
	}
	return &Status{
		MainChainWorldStatus:    mainChainWorldStatus,
		SideChainWorldStatusMap: sideChainWorldStatusMap,
	}
}

// IsMainChain 判断 chain id 是否是主链
func (s *Status) IsMainChain(chainID string) bool {
	s.RLock()
	defer s.RUnlock()

	return chainID == s.MainChainWorldStatus.ChainID
}

// IsSideChain 判断 chain id 是否是侧链
func (s *Status) IsSideChain(chainID string) bool {
	s.RLock()
	defer s.RUnlock()

	_, ok := s.SideChainWorldStatusMap[chainID]
	return ok
}

// AddSideChain 添加一个侧链，如果待添加的链已经是主链返回 false，不进行任何操作，否则直接替换或者添加
func (s *Status) addSideChain(sideChainWorldStatus *WorldStatus) bool {
	s.Lock()
	defer s.Unlock()

	chainID := sideChainWorldStatus.ChainID
	if s.IsMainChain(chainID) {
		// 待添加的 side chain 在 status 处记得是 main chain，则什么都不做，返回 false
		return false
	}
	// 直接添加或者替换
	s.SideChainWorldStatusMap[chainID] = sideChainWorldStatus
	return true
}

// RemoveSideChain 移除一个侧链，如果待移除的链已经是主链返回 false，不进行任何操作，否则直接移除
func (s *Status) removeSideChain(chainID string) bool {
	s.Lock()
	defer s.Unlock()

	if s.IsMainChain(chainID) {
		// 待删除的 side chain 在 status 处记得是 main chain，则什么都不做，返回 false
		return false
	}
	// 直接删除
	delete(s.SideChainWorldStatusMap, chainID)
	return true
}

// ResetMainChain 重新设置主链，如果新的主链原来是侧链，则先删除，然后直接设置（不考虑将旧主链设为侧链）
func (s *Status) resetMainChain(newMainChainWorldStatus *WorldStatus) {
	s.Lock()
	defer s.Unlock()

	chainID := newMainChainWorldStatus.ChainID
	if s.IsSideChain(chainID) {
		// 如果待交换的 chain 是 side chain 则先删除
		s.removeSideChain(chainID)
	}
	s.MainChainWorldStatus = newMainChainWorldStatus
}

// SwapMainChain 重新设置主链，并将旧主链设为侧链，如果新的主链原来是侧链，则先删除。如果新主链已经是主链，返回 false，什么都不做
func (s *Status) swapMainChain(newMainChainWorldStatus *WorldStatus, removeOldMainChain bool) bool {
	s.Lock()
	defer s.Unlock()

	chainID := newMainChainWorldStatus.ChainID
	if s.IsMainChain(chainID) {
		// 如果待交换的 chain 已经是 main chain 则什么都不做，返回 false
		return false
	}
	if s.IsSideChain(chainID) {
		// 如果待交换的 chain 是 side chain 则先删除
		s.removeSideChain(chainID)
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

// Transfer 更新一次世界状态
// TODO chain 相关操作支持
func (s *Status) Transfer(trans *StatusTransfer) bool {
	s.RLock()
	defer s.RUnlock()

	chainID := trans.ChainID
	if chainID == s.MainChainWorldStatus.ChainID {
		s.MainChainWorldStatus.Transfer(trans)
		return true
	}
	if sideChainWorldStatus, ok := s.SideChainWorldStatusMap[chainID]; ok {
		sideChainWorldStatus.Transfer(trans)
	}
	return false
}

func NewWorldStatus(forkHeight int32, bits int64, totalTxn int64, nextMedianTime time.Time) WorldStatus {
	chainID := GenChainID(forkHeight)
	s := WorldStatus{
		ChainID:        chainID,
		ChainHeight:    forkHeight,
		Bits:           bits,
		TotalTxn:       totalTxn,
		NextMedianTime: nextMedianTime,
	}
	return s
}

func (ws *WorldStatus) Transfer(trans *StatusTransfer) {
	if trans.OP == NoneOperate || (trans.OP != ResetField && trans.OP != SwapField) || trans.FieldName == "" {
		return
	}
	// TODO 根据 st 中 field 相关的信息改状态

}
