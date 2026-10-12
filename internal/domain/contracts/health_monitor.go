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
// File: internal/domain/contracts/health_monitor.go
// Author: Gabriel Moraes
// Date: 2026-10-12

package contracts

import (
	"context"

	"noxfort/edge-gateway/internal/domain/models"
)

// CheckResult contains health evaluation outcomes for an upstream service.
type CheckResult struct {
	ServiceID    string
	IsAlive      bool
	LatencyMs    int64
	ErrorMessage string
}

// IHealthMonitor defines the contract for testing upstream services health (DIP).
type IHealthMonitor interface {
	CheckHealth(ctx context.Context, service models.UpstreamService) CheckResult
}
