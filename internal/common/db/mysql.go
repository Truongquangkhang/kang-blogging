package db

import (
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	promcollectors "github.com/prometheus/client_golang/prometheus/collectors"
)

type MysqlConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string

	MaxOpenConns       int
	MaxIdleConns       int
	ConnMaxIdleTimeSec time.Duration
	ConnMaxLifetimeSec time.Duration
}

func NewMySQLConnection(config MysqlConfig) (*sqlx.DB, error) {
	mysqlConfig := mysql.NewConfig()

	mysqlConfig.Net = "tcp"
	mysqlConfig.Addr = fmt.Sprintf("%s:%s", config.Host, config.Port)
	mysqlConfig.User = config.User
	mysqlConfig.Passwd = config.Pass
	mysqlConfig.DBName = config.Name
	mysqlConfig.ParseTime = true

	db, err := sqlx.Connect("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		return nil, errors.Wrap(err, "Cannot connect to MySQL")
	}

	// Config connection pool
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxIdleTime(config.ConnMaxIdleTimeSec * time.Second)
	db.SetConnMaxLifetime(config.ConnMaxLifetimeSec * time.Second)

	collector := promcollectors.NewDBStatsCollector(db.DB, config.Name)
	prometheus.MustRegister(collector)

	return db, nil
}
