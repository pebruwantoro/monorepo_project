package driver

import (
	"github.com/pebruwantoro/monorepo_project/backend/config"
	postgres "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDatabase(cfg config.DBConfig) (db *gorm.DB) {
	dsn := cfg.GetDSN()
	dialector := postgres.Open(dsn)

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("db connection failed")
	}
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConnections)

	if cfg.DebugMode {
		db = db.Debug()
	}

	return
}
