package db

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	promcollectors "github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

var once sync.Once

var dbAdapter DBAdapter

// DBAdapter interface represent adapter connect to DB
type DBAdapter interface {
	Open(config MysqlConfig) error
	DB() *gorm.DB
	Connection() *gorm.DB
}

type adapter struct {
	connection *gorm.DB
	session    *gorm.DB
}

func (db *adapter) Open(config MysqlConfig) error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Pass, config.Host, config.Port, config.Name,
	)

	newLogger := gormLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		gormLogger.Config{
			SlowThreshold: time.Second,     // Slow SQL threshold
			LogLevel:      gormLogger.Info, // Log level
			Colorful:      false,           // Disable color
		},
	)

	DB, err := gorm.Open(
		mysql.Open(
			dsn,
		),
		&gorm.Config{
			Logger: newLogger,
		},
	)

	if err != nil {
		return err
	}

	sqlDB, err := DB.DB()
	if err != nil {
		logrus.Errorf("error %v", err)
		return err
	}

	// Config connection pool
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetConnMaxIdleTime(config.ConnMaxIdleTimeSec * time.Second)
	sqlDB.SetConnMaxLifetime(config.ConnMaxLifetimeSec * time.Second)

	db.connection = DB
	db.session = db.connection.Session(&gorm.Session{})
	collector := promcollectors.NewDBStatsCollector(sqlDB, fmt.Sprintf("%s_gorm", config.Name))
	prometheus.MustRegister(collector)

	return nil
}

func (db *adapter) DB() *gorm.DB {
	return db.connection
}

func (db *adapter) Connection() *gorm.DB {
	return db.connection
}

// NewDB returns a new instance of DB.
func newDB() DBAdapter {
	return &adapter{}
}

func GetDBInstance() DBAdapter {
	if dbAdapter == nil {
		once.Do(func() {
			dbAdapter = newDB()
		})
	}
	return dbAdapter
}
