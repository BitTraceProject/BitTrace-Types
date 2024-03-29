package database

import (
	"github.com/BitTraceProject/BitTrace-Types/pkg/common"
	"github.com/BitTraceProject/BitTrace-Types/pkg/config"
	"github.com/BitTraceProject/BitTrace-Types/pkg/constants"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBInstance(dbConf config.DatabaseConfig) (*gorm.DB, error) {
	var (
		dbInst *gorm.DB
		err    error
	)
	err = common.ExecuteFunctionByRetry(func() error {
		dsn := dbConf.DSN(constants.DATABASE_NAME_BITTRACE)
		dbInst, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		return err
	})
	if err != nil {
		return nil, err
	}
	return dbInst, nil
}

func NewDBInstanceCopy(dbInst *gorm.DB) (*gorm.DB, error) {
	var (
		err error
	)
	err = common.ExecuteFunctionByRetry(func() error {
		dbInst, err = gorm.Open(mysql.New(mysql.Config{Conn: dbInst.ConnPool}), &gorm.Config{})
		return err
	})
	if err != nil {
		return nil, err
	}
	return dbInst, nil
}

// write 函数是通用的，这里提供；read 函数，不通用，在 openapi 中自定义，这里不提供

func TryExecSql(dbInst *gorm.DB, sql string, dbConf config.DatabaseConfig) (*gorm.DB, error) {
	var err error
	if dbInst == nil {
		dbInst, err = NewDBInstance(dbConf)
		if err != nil {
			return nil, err
		}
	} else {
		dbInst, err = NewDBInstanceCopy(dbInst)
		if err != nil {
			return nil, err
		}
	}
	dbInst = dbInst.Exec(sql)
	err = dbInst.Error
	if err != nil && !strings.Contains(err.Error(), "Error 1065 (42000): Query was empty") {
		dbInst = nil
		return nil, err
	}
	return dbInst, nil // 查询用 Raw，写入用 Exec
}

func TryExecPipelineSql(dbInst *gorm.DB, sqlList []string, dbConf config.DatabaseConfig) (*gorm.DB, error) {
	var err error
	for _, sql := range sqlList {
		dbInst, err = TryExecSql(dbInst, sql, dbConf)
		if err != nil {
			return dbInst, err
		}
	}
	return dbInst, nil
}
