// Copyright 2025 Stepan Rabotkin.
// SPDX-License-Identifier: Apache-2.0.

package domain

import "time"

// Multiplayer ...
type Multiplayer string

const (
	// MultiplayerRagemp ...
	MultiplayerRagemp = "ragemp"
)

// Server ...
type Server struct {
	Multiplayer Multiplayer
	Identifier  string
	Name        string
	URL         string
	Gamemode    string
	Lang        string
	Players     uint
	CollectedAt time.Time
}
