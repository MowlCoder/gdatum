// Copyright 2025 Stepan Rabotkin.
// SPDX-License-Identifier: Apache-2.0.

package migrations

import (
	"context"
	"fmt"
	"io/fs"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/pressly/goose/v3"
)

const (
	gooseDialect       = "clickhouse"
	gooseMigrationsDir = "migrations"
)

// Run migrations on provided dsn.
func Run(ctx context.Context, dsn string, fSys fs.FS) error {
	opts, err := clickhouse.ParseDSN(dsn)
	if err != nil {
		return fmt.Errorf("clickhouse.ParseDSN: %w", err)
	}

	db := clickhouse.OpenDB(opts)

	if err = db.PingContext(ctx); err != nil {
		return fmt.Errorf("db.Ping: %w", err)
	}

	goose.SetBaseFS(fSys)

	if err = goose.SetDialect(gooseDialect); err != nil {
		return fmt.Errorf("goose.SetDialect: %w", err)
	}

	if err = goose.UpContext(ctx, db, gooseMigrationsDir); err != nil {
		return fmt.Errorf("goose.Up: %w", err)
	}

	return nil
}
