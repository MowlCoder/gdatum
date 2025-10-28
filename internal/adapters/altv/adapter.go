// Copyright 2025 Stepan Rabotkin.
// SPDX-License-Identifier: Apache-2.0.

package altv

import (
	"context"
	"time"

	"github.com/samber/lo"

	"github.com/EpicStep/gdatum/internal/domain"
	"github.com/EpicStep/gdatum/internal/infrastructure/clients/altv"
)

type client interface {
	Servers(ctx context.Context) (altv.Servers, error)
}

// Adapter ...
type Adapter struct {
	client client
}

// New ...
func New(client client) *Adapter {
	return &Adapter{
		client: client,
	}
}

// Servers ...
func (a *Adapter) Servers(ctx context.Context, collectedAt time.Time) ([]domain.Server, error) {
	servers, err := a.client.Servers(ctx)
	if err != nil {
		return nil, err
	}

	return lo.Map(servers, func(server altv.Server, _ int) domain.Server {
		return domain.Server{
			Multiplayer:  domain.MultiplayerAltv,
			Host:         server.Address,
			Name:         server.Name,
			Gamemode:     server.Gamemode,
			Language:     server.Language,
			PlayersCount: server.PlayersCount,
			CollectedAt:  collectedAt,
		}
	}), nil
}
