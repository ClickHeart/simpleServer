package db

import (
	"fmt"
	"location_program/config"
	"log"
	"log/slog"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var once sync.Once

var dao *gorm.DB

func Init(config *config.Pgsql) *gorm.DB {
	once.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.Host, config.User, config.Password, config.DB, config.Port)

		gormconfig := &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		}

		var err error
		dao, err = gorm.Open(postgres.Open(dsn), gormconfig)
		if err != nil {
			log.Panic("Database Connect Error:", err.Error())
		}

		dbCon, err := dao.DB()
		if err != nil {
			log.Panic("Get dbCon Error:", err.Error())
		}
		dbCon.SetMaxIdleConns(1)
		dbCon.SetMaxOpenConns(5)
		dbCon.SetConnMaxLifetime(time.Hour)
		slog.Info("Pgsql is Connected")
	})

	return dao
}
