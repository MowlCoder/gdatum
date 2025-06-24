// Copyright 2025 Stepan Rabotkin.
// SPDX-License-Identifier: Apache-2.0.

package config

import (
	"fmt"
	"os"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.uber.org/zap/zapcore"
)

// Config of the application.
type Config struct {
	DatabaseDSN string `json:"-"`

	PublicListenAddress string
	AdminListenAddress  string
}

func (c *Config) validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.DatabaseDSN, validation.Required),
		validation.Field(&c.PublicListenAddress, validation.Required),
		validation.Field(&c.AdminListenAddress, validation.Required),
	)
}

// MarshalLogObject help function for zap, that add ability to log config.
func (c *Config) MarshalLogObject(e zapcore.ObjectEncoder) error {
	return e.AddReflected("config", c)
}

// Load Config from env.
func Load() (*Config, error) {
	cfg := &Config{
		DatabaseDSN:         loadValue("DATABASE_DSN", ""),
		PublicListenAddress: loadValue("PUBLIC_LISTEN_ADDRESS", "127.0.0.1:8080"),
		AdminListenAddress:  loadValue("ADMIN_LISTEN_ADDRESS", "127.0.0.1:8081"),
	}

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	return cfg, nil
}

func loadValue(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
