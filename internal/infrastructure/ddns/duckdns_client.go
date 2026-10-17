// NOXFORT EDGE GATEWAY - Unified Multi-Tenant Edge Gateway & Reverse Proxy
// Copyright (C) 2026 Noxfort Systems
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// File: internal/infrastructure/ddns/duckdns_client.go
// Author: Gabriel Moraes
// Date: October 2026

package ddns

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"noxfort/edge-gateway/internal/domain/contracts"
	"noxfort/edge-gateway/internal/domain/models"
)

// DuckDNSClient implements IDDNSSynchronizer for duckdns.org.
type DuckDNSClient struct {
	httpClient *http.Client
	baseURL    string
}

// NewDuckDNSClient builds a client with appropriate timeout.
func NewDuckDNSClient() *DuckDNSClient {
	return &DuckDNSClient{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		baseURL:    "https://www.duckdns.org/update",
	}
}

func (c *DuckDNSClient) ProviderName() string {
	return "duckdns"
}

// Sync contacts DuckDNS HTTP API and updates A/AAAA records.
func (c *DuckDNSClient) Sync(ctx context.Context, target models.DDNSTarget, ipv4, ipv6 string) contracts.SyncResult {
	// If IPv4 is disabled (empty) and IPv6 is present, clear stale A records so DuckDNS doesn't auto-detect CGNAT IP
	if ipv4 == "" && ipv6 != "" {
		clearURL := fmt.Sprintf("%s?domains=%s&token=%s&clear=true", c.baseURL, url.QueryEscape(target.Subdomain), url.QueryEscape(target.Token))
		if clearReq, err := http.NewRequestWithContext(ctx, http.MethodGet, clearURL, nil); err == nil {
			clearReq.Header.Set("User-Agent", "Noxfort-Edge-Gateway/1.0")
			if clearResp, err := c.httpClient.Do(clearReq); err == nil {
				_ = clearResp.Body.Close()
			}
		}
	}

	params := url.Values{}
	params.Set("domains", target.Subdomain)
	params.Set("token", target.Token)
	if ipv4 != "" {
		params.Set("ip", ipv4)
	}
	if ipv6 != "" {
		params.Set("ipv6", ipv6)
	}

	reqURL := fmt.Sprintf("%s?%s", c.baseURL, params.Encode())
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return contracts.SyncResult{
			Subdomain: target.Subdomain,
			Success:   false,
			Error:     fmt.Sprintf("failed constructing request: %v", err),
		}
	}
	req.Header.Set("User-Agent", "Noxfort-Edge-Gateway/1.0")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return contracts.SyncResult{
			Subdomain: target.Subdomain,
			Success:   false,
			Error:     fmt.Sprintf("http network error: %v", err),
		}
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(io.LimitReader(resp.Body, 512))
	if err != nil {
		return contracts.SyncResult{
			Subdomain: target.Subdomain,
			Success:   false,
			Error:     fmt.Sprintf("failed reading response: %v", err),
		}
	}

	body := strings.TrimSpace(string(bodyBytes))
	if strings.EqualFold(body, "OK") {
		return contracts.SyncResult{
			Subdomain: target.Subdomain,
			IPv4:      ipv4,
			IPv6:      ipv6,
			Success:   true,
		}
	}

	return contracts.SyncResult{
		Subdomain: target.Subdomain,
		Success:   false,
		Error:     fmt.Sprintf("DuckDNS rejected update (response: '%s'). Verify Token or Subdomain", body),
	}
}
