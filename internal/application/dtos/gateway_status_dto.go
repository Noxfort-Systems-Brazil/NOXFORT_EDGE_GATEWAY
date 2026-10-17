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
// File: internal/application/dtos/gateway_status_dto.go
// Author: Gabriel Moraes
// Date: October 2026

// Package dtos defines Data Transfer Objects used across Application and Presentation boundaries.
package dtos

import (
	"time"

	"noxfort/edge-gateway/internal/domain/models"
)

// GatewayStatusDTO summarizes real-time telemetry of the Edge Gateway.
type GatewayStatusDTO struct {
	Timestamp       time.Time                `json:"timestamp"`
	UptimeSeconds   int64                    `json:"uptime_seconds"`
	DetectedIPv4    string                   `json:"detected_ipv4"`
	DetectedIPv6    string                   `json:"detected_ipv6"`
	Upstreams       []UpstreamStatusDTO      `json:"upstreams"`
	DDNSTargets     []DDNSTargetStatusDTO    `json:"ddns_targets"`
	TotalRoutes     int                      `json:"total_routes"`
	HealthyServices int                      `json:"healthy_services"`
}

type UpstreamStatusDTO struct {
	ID           string              `json:"id"`
	Name         string              `json:"name"`
	Target       string              `json:"target"`
	Status       models.HealthStatus `json:"status"`
	LatencyMs    int64               `json:"latency_ms"`
	LastChecked  time.Time           `json:"last_checked"`
}

type DDNSTargetStatusDTO struct {
	Domain      string    `json:"domain"`
	Provider    string    `json:"provider"`
	LastIPv4    string    `json:"last_ipv4"`
	LastIPv6    string    `json:"last_ipv6"`
	SyncSuccess bool      `json:"sync_success"`
	LastSync    time.Time `json:"last_sync"`
	LastError   string    `json:"last_error,omitempty"`
}
