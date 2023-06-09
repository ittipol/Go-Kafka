package db

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type sqlLogger struct {
	logger.Interface
}

func (sqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()

	fmt.Printf("%v \n===============================\n", sql)
}

var Conn *gorm.DB

func Init(dsn string) {
	Conn = GetConnection(dsn, false)
}

func GetConnection(dsn string, dryRun bool) *gorm.DB {

	// dial := postgres.Open(dsn)
	dial := mysql.Open(dsn)

	fmt.Println("Create Connection")

	conn, err := gorm.Open(dial, &gorm.Config{
		Logger: &sqlLogger{},
		DryRun: dryRun,
	})

	if err != nil {
		// panic("failed to connect database")
		panic(err)
	}

	fmt.Println("Database Connected")

	return conn
}
