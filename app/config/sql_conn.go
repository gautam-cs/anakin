package config

import (
	"fmt"
	"time"

	"github.com/phuslu/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//ReadMySQLClient mysql handle
var readSQLConnPool *gorm.DB

//WriteMySQLClient mysql handle
var writeSQLConnPool *gorm.DB

func ReadDB() *gorm.DB {
	if readSQLConnPool != nil {
		return readSQLConnPool
	}

	readSQLConnPool = makeSqlConnPool(true)
	return readSQLConnPool
}

func WriteDB() *gorm.DB {
	if writeSQLConnPool != nil {
		return writeSQLConnPool
	}

	writeSQLConnPool = makeSqlConnPool(false)
	return writeSQLConnPool
}

func makeSqlConnPool(readOnly bool) *gorm.DB {

	dbConfig := sqlConfig(readOnly)

	SQLDSN := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Address,
		dbConfig.Database,
		"charset=utf8mb4&parseTime=True&loc=UTC&allowAllFiles=true&readTimeout=5s&timeout=30s&writeTimeout=5s",
	)

	customGormLogger := &GormLogger{
		Log:           log.DefaultLogger,
		SlowThreshold: 10 * time.Second,
		Silent:        false, //!AppConfig.Common.IsDebugEnv(),
	}

	config := mysql.Config{
		DSN: SQLDSN, // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name,
	}

	sql := mysql.New(config)

	gormconfig := &gorm.Config{
		Logger: customGormLogger,
	}

	SQLClient, initError := gorm.Open(sql, gormconfig)

	if initError != nil {
		fmt.Println("SQL connection error:", initError)
		panic(initError)
	}

	db, initError := SQLClient.DB()

	if initError != nil {
		fmt.Println("SQL error:", initError)
		panic(initError)
	}

	MaxIdleConnections := dbConfig.MaxConnections / 10

	db.SetMaxIdleConns(MaxIdleConnections)
	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetConnMaxIdleTime(time.Minute * 5)
	db.SetMaxOpenConns(dbConfig.MaxConnections)

	return SQLClient
}

func sqlConfig(readOnly bool) *mySqlConfig {
	if IsDebugEnv() {
		if readOnly {
			return appConfig.ReadMySQL
		} else {
			return appConfig.WriteMySQL
		}
	}

	if readOnly && appConfig.ReadMySQL != nil {
		return appConfig.ReadMySQL
	} else if !readOnly && appConfig.WriteMySQL != nil {
		return appConfig.WriteMySQL
	}

	key := "readsqldb"
	if !readOnly {
		key = "writesqldb"
	}

	config := new(mySqlConfig)
	readDecodedSecret(key, config)

	if readOnly {
		appConfig.ReadMySQL = config
	} else {
		appConfig.WriteMySQL = config
	}

	return config
}
