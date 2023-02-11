package protocol

import (
	"github.com/BitTraceProject/BitTrace-Types/pkg/common"
)

// meta 基于简单 KV 进行元信息管理, 需要提供的接口包括:
// - 注册 exporter 并与 resolver 建立映射关系: exporter tag 对应 resolver tag
// - exporter tag 查询在不在及与 resolver tag 的双向查询
// - 各种组件的地址信息的注册和查询 (相当于 KV): receiver, mq, collector, resolver-mgr

type (
	// ExporterInfo 保存 exporter 对应的相关元信息：resolver、progress、db table
	ExporterInfo struct {
		ResolverTag   string           `json:"resolver_tag"`
		Table         ExporterTable    `json:"table"`
		JoinTimestamp common.Timestamp `json:"join_timestamp"`

		QuitTimestamp common.Timestamp `json:"quit_timestamp"` // update

		StatusCode ExporterStatusCode `json:"status_code"` // update

		CurrentProgress ExporterProgress `json:"current_progress"` // update

		// ...
	}
	ExporterProgress struct {
		CurrentN      int64  `json:"current_n"`       // 当前日志文件的读取进度
		CurrentFileID int64  `json:"current_file_id"` // 当前日志文件的 id
		CurrentDay    string `json:"current_day"`     // 当前日志文件的 day
	}
	ExporterTable struct {
		TableNameSnapshotData string `json:"table_name_snapshot_data"`
		TableNameSnapshotSync string `json:"table_name_snapshot_sync"`
		TableNameState        string `json:"table_name_state"`
		TableNameRevision     string `json:"table_name_revision"`
	}
	ExporterStatusCode int

	MetaGetExporterInfoArgs struct {
		Key string `json:"key"`
	}
	MetaGetExporterInfoReply struct {
		HasJoin bool `json:"has_join"`
		OK      bool `json:"ok"`

		Info ExporterInfo `json:"info"`
	}
	MetaGetExporterResolverArgs struct {
		Key string `json:"key"`
	}
	MetaGetExporterResolverReply struct {
		HasJoin bool `json:"has_join"`
		OK      bool `json:"ok"`

		ResolverTag string `json:"resolver_tag"`
	}
	MetaExporterHasJoinArgs struct {
		Key string `json:"key"`
	}
	MetaExporterHasJoinReply struct {
		HasJoin bool `json:"has_join"`
		OK      bool `json:"ok"`
	}
	MetaGetExporterTableArgs struct {
		Key string `json:"key"`
	}
	MetaGetExporterTableReply struct {
		HasJoin bool `json:"has_join"`
		OK      bool `json:"ok"`

		Table ExporterTable `json:"table"`
	}
	MetaGetExporterStatusArgs struct {
		Key string `json:"key"`
	}
	MetaGetExporterStatusReply struct {
		HasJoin bool `json:"has_join"`
		OK      bool `json:"ok"`

		JoinTimestamp common.Timestamp   `json:"join_timestamp"`
		QuitTimestamp common.Timestamp   `json:"quit_timestamp"` // update
		StatusCode    ExporterStatusCode `json:"status_code"`    // update
	}
	MetaGetExporterProgressArgs struct {
		Key string `json:"key"`
	}
	MetaGetExporterProgressReply struct {
		HasJoin bool `json:"has_join"`
		OK      bool `json:"ok"`

		CurrentProgress ExporterProgress `json:"current_progress"`
	}
	MetaGetAllExporterInfoArgs struct {
	}
	MetaGetAllExporterInfoReply struct {
		OK bool `json:"ok"`

		ExporterInfo map[string]ExporterInfo `json:"exporter_info"`
	}

	MetaNewExporterInfoArgs struct {
		Key  string       `json:"key"`
		Info ExporterInfo `json:"info"`
	}
	MetaNewExporterInfoReply struct {
		HasJoin bool `json:"has_join"`
		OK      bool `json:"ok"`
	}
	MetaUpdateExporterStatusArgs struct {
		Key        string             `json:"key"`
		StatusCode ExporterStatusCode `json:"status_code"`

		// for StatusLazyQuit
		QuitTimestamp common.Timestamp `json:"quit_timestamp"`
	}
	MetaUpdateExporterStatusReply struct {
		HasJoin bool `json:"has_join"`
		OK      bool `json:"ok"`
	}
	MetaUpdateExporterProgressArgs struct {
		Key             string           `json:"key"`
		CurrentProgress ExporterProgress `json:"current_progress"`
	}
	MetaUpdateExporterProgressReply struct {
		HasJoin bool `json:"has_join"`
		OK      bool `json:"ok"`
	}
	MetaDeleteExporterInfoArgs struct {
		Key string `json:"key"`
	}
	MetaDeleteExporterInfoReply struct {
		HasJoin bool `json:"has_join"`
		OK      bool `json:"ok"`

		Info ExporterInfo `json:"info"`
	}
	MetaClearAllExporterInfoArgs struct {
	}
	MetaClearAllExporterInfoReply struct {
		OK     bool `json:"ok"`
		Number int  `json:"number"`
	}
)

const (
	StatusActive ExporterStatusCode = iota
	StatusLazyQuit
	StatusDead
	StatusUnknown
)
