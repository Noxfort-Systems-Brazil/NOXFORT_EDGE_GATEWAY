// NOXFORT EDGE GATEWAY - Unified Multi-Tenant Edge Gateway & Reverse Proxy
// SPDX-FileCopyrightText: 2026 Noxfort Systems
// SPDX-License-Identifier: Apache-2.0
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
// File: internal/domain/models/upstream_service.go
// Author: Gabriel Moraes
// Date: October 2026

// Package models defines pure domain entities and value objects (Domain Layer).
package models

import (
	"fmt"
	"time"
)

// HealthStatus represents the availability state of an upstream service.
type HealthStatus string

const (
	StatusOnline      HealthStatus = "ONLINE"
	StatusDegraded    HealthStatus = "DEGRADED"
	StatusOffline     HealthStatus = "OFFLINE"
	StatusUnmonitored HealthStatus = "UNMONITORED"
)

// UpstreamService is a pure domain entity representing an internal backend service.
type UpstreamService struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Host         string       `json:"host"`
	Port         int          `json:"port"`
	HealthPath   string       `json:"health_path"`
	Status       HealthStatus `json:"status"`
	LastChecked  time.Time    `json:"last_checked"`
	ResponseTime int64        `json:"response_time_ms"`
}

// NewUpstreamService creates and validates a new UpstreamService entity.
func NewUpstreamService(id, name, host string, port int, healthPath string) (*UpstreamService, error) {
	if id == "" {
		return nil, fmt.Errorf("upstream service id cannot be empty")
	}
	if host == "" {
		host = "127.0.0.1"
	}
	if port <= 0 || port > 65535 {
		return nil, fmt.Errorf("invalid upstream port %d", port)
	}
	if healthPath == "" {
		healthPath = "/health"
	}

	return &UpstreamService{
		ID:          id,
		Name:        name,
		Host:        host,
		Port:        port,
		HealthPath:  healthPath,
		Status:      StatusUnmonitored,
		LastChecked: time.Time{},
	}, nil
}

// TargetAddress returns the network address string (host:port).
func (u *UpstreamService) TargetAddress() string {
	return fmt.Sprintf("%s:%d", u.Host, u.Port)
}

// HealthURL returns the full local URL for health checking.
func (u *UpstreamService) HealthURL() string {
	return fmt.Sprintf("http://%s:%d%s", u.Host, u.Port, u.HealthPath)
}

// MarkOnline updates status to online with measured latency.
func (u *UpstreamService) MarkOnline(latencyMs int64) {
	u.Status = StatusOnline
	u.ResponseTime = latencyMs
	u.LastChecked = time.Now()
}

// MarkOffline updates status to offline.
func (u *UpstreamService) MarkOffline() {
	u.Status = StatusOffline
	u.ResponseTime = 0
	u.LastChecked = time.Now()
}
