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
// File: internal/application/usecases/get_gateway_status_usecase.go
// Author: Gabriel Moraes
// Date: October 2026

package usecases

import (
	"time"

	"noxfort/edge-gateway/internal/application/dtos"
	"noxfort/edge-gateway/internal/domain/models"
)

// GetGatewayStatusUseCase builds real-time status DTO of the edge gateway.
type GetGatewayStatusUseCase struct {
	startTime time.Time
	syncUC    *SyncAllDDNSUseCase
	healthUC  *MonitorHealthUseCase
	routes    []models.RouteRule
}

// NewGetGatewayStatusUseCase constructs the use case instance.
func NewGetGatewayStatusUseCase(
	startTime time.Time,
	syncUC *SyncAllDDNSUseCase,
	healthUC *MonitorHealthUseCase,
	routes []models.RouteRule,
) *GetGatewayStatusUseCase {
	return &GetGatewayStatusUseCase{
		startTime: startTime,
		syncUC:    syncUC,
		healthUC:  healthUC,
		routes:    routes,
	}
}

// Execute compiles and returns the current GatewayStatusDTO.
func (uc *GetGatewayStatusUseCase) Execute() dtos.GatewayStatusDTO {
	now := time.Now()
	uptime := int64(now.Sub(uc.startTime).Seconds())

	upstreams := uc.healthUC.GetUpstreams()
	upstreamDTOs := make([]dtos.UpstreamStatusDTO, 0, len(upstreams))
	healthyCount := 0

	for _, u := range upstreams {
		if u.Status == models.StatusOnline {
			healthyCount++
		}
		upstreamDTOs = append(upstreamDTOs, dtos.UpstreamStatusDTO{
			ID:          u.ID,
			Name:        u.Name,
			Target:      u.TargetAddress(),
			Status:      u.Status,
			LatencyMs:   u.ResponseTime,
			LastChecked: u.LastChecked,
		})
	}

	targets := uc.syncUC.GetTargets()
	targetDTOs := make([]dtos.DDNSTargetStatusDTO, 0, len(targets))
	detectedV4 := ""
	detectedV6 := ""

	for _, t := range targets {
		if t.LastIPv4 != "" {
			detectedV4 = t.LastIPv4
		}
		if t.LastIPv6 != "" {
			detectedV6 = t.LastIPv6
		}
		targetDTOs = append(targetDTOs, dtos.DDNSTargetStatusDTO{
			Domain:      t.FullDomain(),
			Provider:    t.Provider,
			LastIPv4:    t.LastIPv4,
			LastIPv6:    t.LastIPv6,
			SyncSuccess: t.SyncSuccess,
			LastSync:    t.LastSync,
			LastError:   t.LastError,
		})
	}

	return dtos.GatewayStatusDTO{
		Timestamp:       now,
		UptimeSeconds:   uptime,
		DetectedIPv4:    detectedV4,
		DetectedIPv6:    detectedV6,
		Upstreams:       upstreamDTOs,
		DDNSTargets:     targetDTOs,
		TotalRoutes:     len(uc.routes),
		HealthyServices: healthyCount,
	}
}
