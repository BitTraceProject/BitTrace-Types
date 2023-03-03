package database

import (
	"fmt"
	"strings"
)

// sql 语句编码在 go 文件中，便于复用
// 这里只定义通用的那些，有些场景的 sql，可以自定义在其他位置

const (
	sqlCreateTableSnapshotData = `
CREATE TABLE IF NOT EXISTS %s (
  snapshot_id varchar(35) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  target_chain_id int NOT NULL,
  target_chain_height int NOT NULL,
  block_hash varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  is_orphan bool NOT NULL,
  init_timestamp varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  final_timestamp varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY(snapshot_id),
  INDEX index_snapshot(target_chain_id,target_chain_height,block_hash),
  INDEX index_snapshot_block_hash(block_hash)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='snapshot data table for exporter.';
`
	sqlCreateTableSnapshotSync = `
CREATE TABLE IF NOT EXISTS %s (
  snapshot_id varchar(35) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  target_chain_id int NOT NULL,
  target_chain_height int NOT NULL,
  sync_timestamp varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY(snapshot_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='snapshot sync table for exporter';
`
	sqlCreateTableState = `
CREATE TABLE IF NOT EXISTS %s (
  snapshot_id varchar(35) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  snapshot_type tinyint(1) NOT NULL,
  best_block_hash varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  height int NOT NULL,
  bits int unsigned NOT NULL,
  block_size bigint unsigned NOT NULL,
  block_weight bigint unsigned NOT NULL,
  num_txns bigint unsigned NOT NULL,
  total_txns bigint unsigned NOT NULL,
  median_timestamp varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY(snapshot_id,snapshot_type)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='state table for exporter';
`
	sqlCreateTableRevision = `
CREATE TABLE IF NOT EXISTS %s (
  snapshot_id varchar(35) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  revision_type tinyint unsigned NOT NULL,
  init_timestamp varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  final_timestamp varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  init_data json DEFAULT NULL,
  final_data json DEFAULT NULL,
  PRIMARY KEY(snapshot_id,revision_type,init_timestamp)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='revision table for exporter';
`
	sqlCreateTableEventOrphan = `
CREATE TABLE IF NOT EXISTS %s (
  snapshot_id varchar(35) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  event_type_orphan tinyint(1) NOT NULL,
  orphan_block_hash varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  event_orphan_timestamp varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY(snapshot_id,event_type_orphan),
  INDEX index_orphan_block_hash(orphan_block_hash)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='snapshot data table for exporter.';
`
	sqlInsert = `INSERT INTO %s(%s) VALUES %s;` // INSERT INTO table1(i) VALUES(1),(2),(3),(4),(5);
)

func SqlCreateTableSnapshotData(tableName string) string {
	return fmt.Sprintf(sqlCreateTableSnapshotData, fmt.Sprintf("`%s`", tableName))
}

func SqlCreateTableSnapshotSync(tableName string) string {
	return fmt.Sprintf(sqlCreateTableSnapshotSync, fmt.Sprintf("`%s`", tableName))
}

func SqlCreateTableState(tableName string) string {
	return fmt.Sprintf(sqlCreateTableState, fmt.Sprintf("`%s`", tableName))
}

func SqlCreateTableRevision(tableName string) string {
	return fmt.Sprintf(sqlCreateTableRevision, fmt.Sprintf("`%s`", tableName))
}

func SqlCreateTableEventOrphan(tableName string) string {
	return fmt.Sprintf(sqlCreateTableEventOrphan, fmt.Sprintf("`%s`", tableName))
}

func SqlInsertSnapshotData(tableName string, snapshotData ...TableSnapshotData) string {
	if len(snapshotData) == 0 {
		return ""
	}
	fields := "snapshot_id,target_chain_id,target_chain_height,block_hash,is_orphan,init_timestamp,final_timestamp"
	values := make([]string, len(snapshotData))
	for i, data := range snapshotData {
		value := fmt.Sprintf("('%s','%s',%d,'%s','%v','%s','%s')",
			data.SnapshotID,
			data.TargetChainID,
			data.TargetChainHeight,
			data.BlockHash,
			data.IsOrphan,
			data.InitTimestamp,
			data.FinalTimestamp,
		)
		values[i] = value
	}
	return fmt.Sprintf(sqlInsert, fmt.Sprintf("`%s`", tableName), fields, strings.Join(values, ","))
}

func SqlInsertSnapshotSync(tableName string, snapshotSync ...TableSnapshotSync) string {
	if len(snapshotSync) == 0 {
		return ""
	}
	fields := "snapshot_id,target_chain_id,target_chain_height,sync_timestamp"
	values := make([]string, len(snapshotSync))
	for i, sync := range snapshotSync {
		value := fmt.Sprintf("('%s','%s',%d,'%s')",
			sync.SnapshotID,
			sync.TargetChainID,
			sync.TargetChainHeight,
			sync.SyncTimestamp,
		)
		values[i] = value
	}
	return fmt.Sprintf(sqlInsert, fmt.Sprintf("`%s`", tableName), fields, strings.Join(values, ","))
}

func SqlInsertState(tableName string, state ...TableState) string {
	if len(state) == 0 {
		return ""
	}
	fields := "snapshot_id,snapshot_type,best_block_hash,height,bits,block_size,block_weight,num_txns,total_txns,median_timestamp"
	values := make([]string, len(state))
	for i, s := range state {
		value := fmt.Sprintf("('%s',%d,'%s',%d,%d,%d,%d,%d,%d,'%s')",
			s.SnapshotID,
			s.SnapshotType,
			s.BestBlockHash,
			s.Height,
			s.Bits,
			s.BlockSize,
			s.BlockWeight,
			s.NumTxns,
			s.TotalTxns,
			s.MedianTimestamp,
		)
		values[i] = value
	}
	return fmt.Sprintf(sqlInsert, fmt.Sprintf("`%s`", tableName), fields, strings.Join(values, ","))
}

func SqlInsertRevision(tableName string, revision ...TableRevision) string {
	if len(revision) == 0 {
		return ""
	}
	fields := "snapshot_id,revision_type,init_timestamp,final_timestamp,init_data,final_data"
	values := make([]string, len(revision))
	for i, r := range revision {
		value := fmt.Sprintf("('%s',%d,'%s','%s','%s','%s')",
			r.SnapshotID,
			r.RevisionType,
			r.InitTimestamp,
			r.FinalTimestamp,
			r.InitData,
			r.FinalData,
		)
		values[i] = value
	}
	return fmt.Sprintf(sqlInsert, fmt.Sprintf("`%s`", tableName), fields, strings.Join(values, ","))
}

func SqlInsertOrphanEvent(tableName string, orphanEvent ...TableEventOrphan) string {
	if len(orphanEvent) == 0 {
		return ""
	}
	fields := "snapshot_id,event_type_orphan,orphan_block_hash,event_orphan_timestamp"
	values := make([]string, len(orphanEvent))
	for i, e := range orphanEvent {
		value := fmt.Sprintf("('%s',%d,'%s','%s')",
			e.SnapshotID,
			e.EventTypeOrphan,
			e.OrphanBlockHash,
			e.EventOrphanTimestamp,
		)
		values[i] = value
	}
	return fmt.Sprintf(sqlInsert, fmt.Sprintf("`%s`", tableName), fields, strings.Join(values, ","))
}
