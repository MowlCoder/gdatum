// Copyright 2025 Stepan Rabotkin.
// SPDX-License-Identifier: Apache-2.0.

package magestic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	serverListURL = "https://api.majestic-files.com/meta/servers"
)

// Client ...
type Client struct {
	client *http.Client
}

// NewOpts ...
type NewOpts struct {
	HTTPClient *http.Client
}

func (o *NewOpts) setDefaults() {
	if o.HTTPClient == nil {
		o.HTTPClient = http.DefaultClient
	}
}

// New returns new Client.
func New(opts NewOpts) *Client {
	opts.setDefaults()

	return &Client{
		client: opts.HTTPClient,
	}
}

// Servers returns magestic servers.
func (c *Client) Servers(ctx context.Context) (Servers, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, serverListURL, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client.Do: %w", err)
	}

	defer resp.Body.Close() //nolint:errcheck

	var responseBody getServersResponse

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return nil, fmt.Errorf("json.Decode: %w", err)
	}

	return responseBody.Result.Servers, nil
}
