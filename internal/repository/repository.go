// Copyright 2025 Stepan Rabotkin.
// SPDX-License-Identifier: Apache-2.0.

package repository

import (
	"context"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"

	"github.com/EpicStep/gdatum/internal/domain"
)

// Repository ...
type Repository struct {
	db driver.Conn
}

// New ...
func New(db driver.Conn) *Repository {
	return &Repository{
		db: db,
	}
}

// InsertServers ...
func (r *Repository) InsertServers(ctx context.Context, servers []*domain.Server) error {
	if len(servers) == 0 {
		return nil
	}

	batch, err := r.db.PrepareBatch(ctx, "INSERT INTO servers_metrics_raw (multiplayer, identifier, name, lang, gamemode, url, players, timestamp)")
	if err != nil {
		return fmt.Errorf("r.db.PrepareBatch: %w", err)
	}

	for _, server := range servers {
		err = batch.Append(
			string(server.Multiplayer),
			server.Identifier,
			server.Name,
			server.Lang,
			server.Gamemode,
			server.URL,
			server.Players,
			server.CollectedAt,
		)
		if err != nil {
			return fmt.Errorf("batch.Append: %w", err)
		}
	}

	if err = batch.Send(); err != nil {
		return fmt.Errorf("batch.Send: %w", err)
	}

	return nil
}
