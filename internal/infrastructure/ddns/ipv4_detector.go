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
// File: internal/infrastructure/ddns/ipv4_detector.go
// Author: Gabriel Moraes
// Date: 2026-10-12

package ddns

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"
)

// PublicIPv4Detector determines public IPv4 address via reliable endpoints.
type PublicIPv4Detector struct {
	client    *http.Client
	endpoints []string
}

// NewPublicIPv4Detector constructs an IPv4 detector with fast fallbacks.
func NewPublicIPv4Detector() *PublicIPv4Detector {
	return &PublicIPv4Detector{
		client: &http.Client{Timeout: 5 * time.Second},
		endpoints: []string{
			"https://api.ipify.org",
			"https://icanhazip.com",
			"https://ifconfig.me/ip",
		},
	}
}

// DetectPublicIPv4 checks endpoints sequentially until a valid IPv4 is retrieved.
func (d *PublicIPv4Detector) DetectPublicIPv4(ctx context.Context) string {
	for _, ep := range d.endpoints {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
		if err != nil {
			continue
		}
		resp, err := d.client.Do(req)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		bodyBytes, err := io.ReadAll(io.LimitReader(resp.Body, 64))
		if err != nil {
			continue
		}
		ip := strings.TrimSpace(string(bodyBytes))
		if ip != "" && !strings.Contains(ip, ":") { // Valid IPv4 check
			return ip
		}
	}
	return ""
}
