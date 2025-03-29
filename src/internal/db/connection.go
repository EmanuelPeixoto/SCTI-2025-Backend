package db

import (
  "log"
	"scti/config"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func Connect(cfg config.Config) *gorm.DB {
  var err error
  if cfg.DSN == "" {
    log.Fatalf("DSN was empty")
  }
  DB, err = gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),
  })
  if err != nil {
    log.Fatalf("Failed to connect to database: %v", err)
  }
  log.Println("Connected to postgres instance")
  return DB
}


