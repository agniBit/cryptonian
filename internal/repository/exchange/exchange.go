package exchange

import (
	"context"

	"github.com/agniBit/cryptonian/internal/logger"
	"github.com/agniBit/cryptonian/model/cfg"
	"gorm.io/gorm"
)

type ExchangeRepository struct {
	cfg *cfg.Config
	gdb *gorm.DB
}

func NewExchangeRepository(cfg *cfg.Config, gfb *gorm.DB) *ExchangeRepository {
	return &ExchangeRepository{
		cfg: cfg,
		gdb: gfb,
	}
}

func (r *ExchangeRepository) Migrate() error {

	err := r.gdb.AutoMigrate(&Account{})
	if err != nil {
		logger.Fatal(context.Background(), "Error while migrating table", err, map[string]interface{}{"table": "Account"})
		return err
	}

	// create sequence for Exchange table
	err = r.gdb.Raw("CREATE SEQUENCE exchange_id_seq; ALTER TABLE ADD COLUMN id VARCHAR(10) GENERATED ALWAYS AS ('EXC-' || nextval('exchange_id_seq')::text, 6) STORED").Error
	if err != nil {
		logger.Fatal(context.Background(), "Error while adding column to table", err, map[string]interface{}{"table": "ExchangeRepository"})
	}

	// create sequence for Account table
	err = r.gdb.Raw("CREATE SEQUENCE account_id_seq; ALTER TABLE ADD COLUMN id VARCHAR(10) GENERATED ALWAYS AS ('ACC-' || nextval('account_id_seq')::text, 6) STORED").Error
	if err != nil {
		logger.Fatal(context.Background(), "Error while adding column to table", err, map[string]interface{}{"table": "ExchangeRepository"})
	}

	return nil
}
