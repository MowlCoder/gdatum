// Copyright 2025 Stepan Rabotkin.
// SPDX-License-Identifier: Apache-2.0.

package altv

// Servers is a altv servers.
type Servers []Server

// Server is a altv server.
type Server struct {
	Name         string `json:"name"`
	Gamemode     string `json:"gameMode"`
	Website      string `json:"website"`
	Language     string `json:"language"`
	PlayersCount int32  `json:"playersCount"`
	Address      string `json:"address"`
}
