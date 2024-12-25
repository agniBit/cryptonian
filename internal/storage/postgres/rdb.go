package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/agniBit/cryptonian/internal/constants"
	"github.com/agniBit/cryptonian/internal/logger"
	"github.com/agniBit/cryptonian/model/cfg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func GetRepositoryFromContext(ctx context.Context, db *gorm.DB) *gorm.DB {
	gdb, ok := ctx.Value(string(constants.ContextKeyRepository)).(*gorm.DB)
	if !ok || gdb == nil {
		return db
	}

	timeoutCtx, _ := context.WithTimeout(ctx, time.Duration(3)*time.Second)
	return gdb.WithContext(timeoutCtx)
}

func NewDB(ctx context.Context, cfg *cfg.Config) *gorm.DB {
	_gdb, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=asia/kolkata connect_timeout=%d",
			cfg.Rdb.Host, cfg.Rdb.Username, cfg.Rdb.Password, cfg.Rdb.DbName, cfg.Rdb.Port, cfg.Rdb.ConnectTimeoutInSeconds),
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		FullSaveAssociations:   true,
		PrepareStmt:            true,
		Logger:                 gormLogger.Default.LogMode(gormLogger.Info),
	})

	if err != nil {
		logger.Fatal(ctx, "failed to connect database", err, nil)
	}

	db, err := _gdb.DB()
	if err != nil {
		logger.Fatal(ctx, "failed to connect database", err, nil)
	}

	db.SetMaxIdleConns(cfg.Rdb.MaxIdleConns)
	db.SetMaxOpenConns(cfg.Rdb.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(cfg.Rdb.ConnMaxLifetimeInMinutes) * time.Minute)

	return _gdb
}
