package config

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/seki-shinnosuke/study-golang/util/logger"
)

func NewDB(
	timeZone string,
	rdb *RDB,
) (*sql.DB, error) {
	jst, _ := time.LoadLocation(timeZone)
	c := mysql.Config{
		Addr:      rdb.Host,
		DBName:    rdb.DBName,
		User:      rdb.User,
		Passwd:    rdb.Passwd,
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	logger.Info("Connect RDB Host:%v", rdb.Host)
	db, err := sql.Open("mysql", c.FormatDSN())
	return db, err
}
