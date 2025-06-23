// Copyright 2025 Stepan Rabotkin.
// SPDX-License-Identifier: Apache-2.0.

package api

import (
	"context"
	"github.com/EpicStep/gdatum/pkg/api"
)

// Handlers ...
type Handlers struct {
	api.UnimplementedHandler
}

// New returns new Handlers.
func New() *Handlers {
	return &Handlers{}
}

func (h *Handlers) GetMultiplayersSummary(ctx context.Context, params api.GetMultiplayersSummaryParams) (api.GetMultiplayersSummaryOKItem, error) {
	params.Order.Get()
}
