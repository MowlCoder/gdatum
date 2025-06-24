// Copyright 2025 Stepan Rabotkin.
// SPDX-License-Identifier: Apache-2.0.

package gdatum

import (
	"embed"
)

// MigrationsFS ...
//
//go:embed migrations
var MigrationsFS embed.FS
