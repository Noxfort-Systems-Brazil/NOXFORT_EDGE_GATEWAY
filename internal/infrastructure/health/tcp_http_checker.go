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
// File: internal/infrastructure/health/tcp_http_checker.go
// Author: Gabriel Moraes
// Date: 2026-10-12

package health

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"noxfort/edge-gateway/internal/domain/contracts"
	"noxfort/edge-gateway/internal/domain/models"
)

// TCPAndHTTPChecker checks service health via HTTP endpoint, falling back to raw TCP connect.
type TCPAndHTTPChecker struct {
	httpClient *http.Client
	dialer     *net.Dialer
}

// NewTCPAndHTTPChecker creates a robust health checker.
func NewTCPAndHTTPChecker() *TCPAndHTTPChecker {
	return &TCPAndHTTPChecker{
		httpClient: &http.Client{Timeout: 2 * time.Second},
		dialer:     &net.Dialer{Timeout: 1500 * time.Millisecond},
	}
}

// CheckHealth implements IHealthMonitor contract.
func (c *TCPAndHTTPChecker) CheckHealth(ctx context.Context, service models.UpstreamService) contracts.CheckResult {
	start := time.Now()

	// 1. Try HTTP health endpoint first if available
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, service.HealthURL(), nil)
	if err == nil {
		resp, httpErr := c.httpClient.Do(req)
		if httpErr == nil {
			defer resp.Body.Close()
			latency := time.Since(start).Milliseconds()
			if resp.StatusCode < 500 {
				return contracts.CheckResult{
					ServiceID: service.ID,
					IsAlive:   true,
					LatencyMs: latency,
				}
			}
		}
	}

	// 2. Fallback to raw TCP ping on target port
	conn, dialErr := c.dialer.DialContext(ctx, "tcp", service.TargetAddress())
	latency := time.Since(start).Milliseconds()
	if dialErr != nil {
		return contracts.CheckResult{
			ServiceID:    service.ID,
			IsAlive:      false,
			LatencyMs:    0,
			ErrorMessage: fmt.Sprintf("TCP dial failed: %v", dialErr),
		}
	}
	_ = conn.Close()

	return contracts.CheckResult{
		ServiceID: service.ID,
		IsAlive:   true,
		LatencyMs: latency,
	}
}
