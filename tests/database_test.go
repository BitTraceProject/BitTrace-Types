package tests

import (
	"github.com/BitTraceProject/BitTrace-Types/pkg/common"
	"github.com/BitTraceProject/BitTrace-Types/pkg/structure"
	"testing"

	"github.com/BitTraceProject/BitTrace-Types/pkg/config"
	"github.com/BitTraceProject/BitTrace-Types/pkg/database"
	"github.com/BitTraceProject/BitTrace-Types/pkg/protocol"
)

func TestNewDBInstance(t *testing.T) {
	dbInst, err := database.NewDBInstance(config.DatabaseConfig{
		Address:  "localhost:3306",
		Username: "root",
		Password: "admin",
	})
	if err != nil {
		t.Fatal(err)
	}
	var result string
	dbInst.Raw("show databases;").Last(&result)
	t.Log(result)

	dbInst, err = database.NewDBInstanceCopy(dbInst)
	if err != nil {
		t.Fatal(err)
	}
	dbInst.Raw("show tables;").First(&result)
	t.Log(result)
}

func TestCreateTable(t *testing.T) {
	dbConf := config.DatabaseConfig{
		Address:  "localhost:3306",
		Username: "root",
		Password: "admin",
	}
	dbInst, err := database.NewDBInstance(dbConf)
	if err != nil {
		t.Fatal(err)
	}

	table := protocol.ExporterTable{
		TableNameSnapshotData: "table_snapshot_data_exporter.test",
		TableNameSnapshotSync: "table_snapshot_sync_exporter.test",
		TableNameState:        "table_state_exporter.test",
		TableNameRevision:     "table_revision_exporter.test",
		TableNameEventOrphan:  "table_event_orphan_exporter.test",
	}

	sqlList := []string{
		database.SqlCreateTableSnapshotData(table.TableNameSnapshotData),
		database.SqlCreateTableSnapshotSync(table.TableNameSnapshotSync),
		database.SqlCreateTableState(table.TableNameState),
		database.SqlCreateTableRevision(table.TableNameRevision),
		database.SqlCreateTableEventOrphan(table.TableNameEventOrphan),
	}
	t.Log(sqlList)
	dbInst, err = database.TryExecPipelineSql(dbInst, sqlList, dbConf)
	if err != nil {
		t.Fatal(err)
	}
}

func TestInsert(t *testing.T) {
	dbConf := config.DatabaseConfig{
		Address:  "localhost:3306",
		Username: "root",
		Password: "admin",
	}
	dbInst, err := database.NewDBInstance(dbConf)
	if err != nil {
		t.Fatal(err)
	}

	snapshotSyncList := []database.TableSnapshotSync{
		database.TableSnapshotSync{
			SnapshotID:        "1676038806838405298-0-541457",
			TargetChainID:     "0",
			TargetChainHeight: 541457,
			SyncTimestamp:     "1676038806838405298",
		},
		database.TableSnapshotSync{
			SnapshotID:        "1676038828881628593-0-541458",
			TargetChainID:     "0",
			TargetChainHeight: 541458,
			SyncTimestamp:     "1676038828881628593",
		},
	}
	stateList := []database.TableState{
		{
			SnapshotID:      "1676038828881628593-0-541458",
			SnapshotType:    2,
			BestBlockHash:   "00000000000000000000f997e373cf36f6cdc1e460508c4c27cd64691ba25bce",
			Height:          541458,
			Bits:            0,
			BlockSize:       0,
			BlockWeight:     0,
			NumTxns:         0,
			TotalTxns:       0,
			MedianTimestamp: "",
		},
	}
	revisionList := []database.TableRevision{
		{
			SnapshotID:     "1676038806838405298-0-541457",
			RevisionType:   0,
			InitTimestamp:  "1676038806838405298",
			FinalTimestamp: "",
			InitData:       `{"a":"aaa","b":12,"c":true}`,
			FinalData:      `{"a":"aaa","b":12,"c":true}`,
		},
		{
			SnapshotID:     "1676038828881628593-0-541458",
			RevisionType:   0,
			InitTimestamp:  "1676038828881628593",
			FinalTimestamp: "",
			InitData:       common.StructToJsonStr(nil),
			FinalData:      common.StructToJsonStr(stateList[0]),
		},
	}
	eventOrphanList := []database.TableEventOrphan{
		{
			SnapshotID:           "1676038806838405298-0-541457",
			EventTypeOrphan:      int(structure.EventTypeOrphanConnect),
			OrphanBlockHash:      "0xsss",
			EventOrphanTimestamp: "",
		},
		{
			SnapshotID:           "1676038828881628593-0-541458",
			EventTypeOrphan:      int(structure.EventTypeOrphanDiscard),
			OrphanBlockHash:      "0xzzz",
			EventOrphanTimestamp: "",
		},
	}
	snapshotDataList := []database.TableSnapshotData{
		database.TableSnapshotData{
			SnapshotID:        "11111",
			TargetChainID:     "1",
			TargetChainHeight: 0,
			BlockHash:         "",
			IsOrphan:          1,
			InitTimestamp:     "1",
			FinalTimestamp:    "1",
		},
	}
	sqlList := []string{
		database.SqlInsertRevision("table_revision_exporter.test", revisionList...),
		database.SqlInsertSnapshotSync("table_snapshot_sync_exporter.test", snapshotSyncList...),
		database.SqlInsertState("table_state_exporter.test", stateList...),
		database.SqlInsertOrphanEvent("table_event_orphan_exporter.test", eventOrphanList...),
		database.SqlInsertSnapshotData("table_snapshot_data_exporter.test", snapshotDataList...),
	}
	t.Log(sqlList)
	dbInst, err = database.TryExecPipelineSql(dbInst, sqlList, dbConf)
	if err != nil {
		t.Fatal(err)
	}
}
