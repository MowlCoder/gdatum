// Copyright 2025 Stepan Rabotkin.
// SPDX-License-Identifier: Apache-2.0.

package backoff

import "github.com/cenkalti/backoff/v5"

// EmptyReturnOperation is helper to call operation without return statement.
func EmptyReturnOperation(operation func() error) backoff.Operation[int] {
	return func() (int, error) {
		return 0, operation()
	}
}
