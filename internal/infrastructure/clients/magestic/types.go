// Copyright 2025 Stepan Rabotkin.
// SPDX-License-Identifier: Apache-2.0.

package magestic

type getServersResponse struct {
	Code   int              `json:"code"`
	Status bool             `json:"status"`
	Result getServersResult `json:"result"`
}

type getServersResult struct {
	Ok      bool    `json:"ok"`
	Servers Servers `json:"servers"`
}

// Servers is a magestic servers.
type Servers []Server

// Server is a magestic server.
type Server struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Players int32  `json:"players"`
	IP      string `json:"ip"`
}
