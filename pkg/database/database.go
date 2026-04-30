package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	 _"github.com/lib/pq" 
	 "github.com/yousefggg/common-lib/pkg/config"
)
func NewClient(ctx context.Context , cfg config.DatabaseConfig)(*sql.DB, error){
	database, err := sql.Open("postgres", cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}
	database.SetMaxOpenConns(cfg.MaxOpenConns)
	database.SetMaxIdleConns(cfg.MaxIdleConns)
	database.SetConnMaxLifetime(time.Hour) 
	database.SetConnMaxIdleTime(time.Minute * 30)

	ctx, cancel := context.WithTimeout(ctx, cfg.ConnTimeout)
	defer cancel()
	if err := database.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return database, nil
}