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
// File: internal/domain/models/route_rule.go
// Author: Gabriel Moraes
// Date: October 2026

package models

import (
	"fmt"
	"strings"
)

// RouteRule is an immutable Value Object representing a reverse proxy routing rule.
type RouteRule struct {
	Domain          string `json:"domain"`
	UpstreamID      string `json:"upstream_id"`
	EnableTLS       bool   `json:"enable_tls"`
	SupportWebsocket bool   `json:"support_websocket"`
	StripPathPrefix string `json:"strip_path_prefix,omitempty"`
}

// NewRouteRule creates and validates a RouteRule Value Object.
func NewRouteRule(domain, upstreamID string, enableTLS, supportWS bool) (*RouteRule, error) {
	cleanDomain := strings.TrimSpace(domain)
	if cleanDomain == "" {
		return nil, fmt.Errorf("route domain cannot be empty")
	}
	cleanUpstream := strings.TrimSpace(upstreamID)
	if cleanUpstream == "" {
		return nil, fmt.Errorf("upstream id cannot be empty for route %s", cleanDomain)
	}

	return &RouteRule{
		Domain:           cleanDomain,
		UpstreamID:       cleanUpstream,
		EnableTLS:        enableTLS,
		SupportWebsocket: supportWS,
	}, nil
}
