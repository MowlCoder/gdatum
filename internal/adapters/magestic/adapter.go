// Copyright 2025 Stepan Rabotkin.
// SPDX-License-Identifier: Apache-2.0.

package magestic

import (
	"context"
	"time"

	"github.com/samber/lo"

	"github.com/EpicStep/gdatum/internal/domain"
	"github.com/EpicStep/gdatum/internal/infrastructure/clients/magestic"
)

type client interface {
	Servers(ctx context.Context) (magestic.Servers, error)
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

	return lo.Map(servers, func(server magestic.Server, _ int) domain.Server {
		return domain.Server{
			Multiplayer:  domain.MultiplayerMagestic,
			Host:         server.Ip,
			Name:         server.Name,
			URL:          "https://majestic-rp.ru",
			Gamemode:     "Roleplay",
			Language:     server.Country,
			PlayersCount: server.Players,
			CollectedAt:  collectedAt,
		}
	}), nil
}
