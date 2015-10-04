package model

import (
	"database/sql"
	"github.com/folded-ear/datawell/config"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	sqlDb  *sql.DB
	gormDb *gorm.DB
)

func DB() (*sql.DB, error) {
	if sqlDb == nil {
		cfg := config.LoadConfig()
		db, err := sql.Open(cfg.DriverName, cfg.DataSourceName)
		if err != nil {
			return nil, err
		}
		err = db.Ping()
		if err != nil {
			return nil, err
		}
		sqlDb = db
	}
	return sqlDb, nil
}

func GORM() (*gorm.DB, error) {
	if gormDb == nil {
		cfg := config.LoadConfig()
		db, err := DB()
		if err != nil {
			return nil, err
		}
		gdb, err := gorm.Open(cfg.DriverName, db)
		if err != nil {
			return nil, err
		}
		gormDb = &gdb
	}
	return gormDb, nil
}
